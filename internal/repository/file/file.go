package file

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"mindstore/internal/object/dto/file"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/stream"
	"strconv"
	"strings"
)

type Repo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) CreateWithMind(c ctx.Ctx, files []*model.FileData, mindId hash.Int) (err error) {
	fs := stream.Mapper(files, func(in *model.FileData) model.FileData {
		return *in
	})
	query, args, err := r.DB.BindNamed(`INSERT INTO file(path, name, hashed_id, access, size,
created_by) VALUES (:path, :name, :hashed_id, :access, :size, :created_by) RETURNING id;`, fs)
	if err != nil {
		return errors.Join(errors.New("500: "), err)
	}

	err = r.DB.SelectContext(c, &fs, query, args...)
	if err != nil {
		return err
	}

	for i, f := range fs {
		files[i].Id = f.Id
	}

	mindFiles := stream.Mapper(fs, func(f model.FileData) *model.MindFile {
		return &model.MindFile{MindId: mindId, FileId: f.Id}
	})

	_, err = r.DB.NamedExecContext(c, `INSERT INTO mind_file(mind_id, file_id) 
VALUES (:mind_id, :file_id)`, mindFiles)
	if err != nil {
		stream.ForEach(fs, func(f model.FileData) {
			r.DB.Exec(`DELETE FROM file WHERE id=$1`, f.Id)
		})
		return err
	}

	return nil
}

func (r *Repo) Create(c ctx.Ctx, input *model.FileData) (err error) {
	if input.Access == 0 {
		input.Access = 99
	}
	query, args, err := r.DB.BindNamed(`INSERT INTO file(path, name, hashed_id, access, size,
created_by) VALUES (:path, :name, :hashed_id, :access, :size, :created_by) RETURNING id;`, input)
	if err != nil {
		return errors.Join(errors.New("500: "), err)
	}

	err = r.DB.GetContext(c, input, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetByMindIds(c ctx.Ctx, mindIds []hash.Int) ([]file.List, error) {
	if len(mindIds) == 0 {
		return []file.List{}, nil
	}

	mindIdsStr := stream.Mapper(mindIds, func(i hash.Int) string {
		return strconv.Itoa(int(i))
	})
	mindIdsIn := strings.Join(mindIdsStr, ",")

	list := make([]file.List, 0, len(mindIds))

	err := r.DB.SelectContext(c, &list,
		fmt.Sprintf(`SELECT mf.mind_id, f.id,f.name,f.hashed_id,f.access,f.size FROM file f
INNER JOIN mind_file mf on f.id = mf.file_id
WHERE mf.deleted_at IS NULL AND mf.mind_id IN (%s) ORDER BY mind_id, f.id DESC`, mindIdsIn))

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Repo) Delete(c ctx.Ctx, input *file.DeleteMind) (err error) {
	mfs := make([]model.MindFile, 0)
	err = r.DB.SelectContext(c, &mfs, `SELECT mf.mind_id FROM mind_file mf
JOIN file f on f.id = mf.file_id
WHERE f.deleted_at IS NULL AND mf.deleted_at IS NULL AND
f.id=$1 AND f.created_by=$2`, input.FileId, input.UserId)

	if err != nil {
		return err
	}

	hasMindFile := stream.AnyMatch(mfs, func(mf model.MindFile) bool {
		return mf.MindId == input.MindId
	})
	if !hasMindFile {
		return errors.New("mind_file not fount")
	}

	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.ExecContext(c, `UPDATE mind_file SET deleted_at=now()
WHERE mind_id=$1 AND file_id=$2`, input.MindId, input.FileId)
	if err != nil {
		return err
	}

	if len(mfs) == 1 {
		_, err = tx.ExecContext(c, `UPDATE file SET deleted_at=now(), deleted_by=$1
WHERE id=$2`, input.DeletedBy, input.FileId)
		if err != nil {
			return err
		}
	}

	return err
}

func (r *Repo) GetPathById(c ctx.Ctx, fileId, userId hash.Int) (path string, err error) {
	f := new(model.FileData)
	err = r.DB.GetContext(c, f, `SELECT path FROM file
WHERE deleted_at IS NULL AND id=$1 AND (access = 99 OR created_by=$2)`, fileId, userId)

	if err != nil {
		return "", err
	}

	return f.Path, nil
}

func (r *Repo) GetAvatarPathByUserId(c ctx.Ctx, userId hash.Int) (path string, err error) {
	f := new(model.FileData)
	err = r.DB.GetContext(c, f, `SELECT f.path FROM file f
JOIN users u on f.id = u.avatar_id
WHERE f.deleted_at IS NULL AND u.deleted_at IS NULL AND u.id=$1 AND access = 99`, userId)

	if err != nil {
		return "", err
	}

	return f.Path, nil
}
