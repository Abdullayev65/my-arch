package mind

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"mindstore/internal/object/dto/mind"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/repoutill"
	"strings"
)

type Repo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(c ctx.Ctx, input *mind.Create) (*hash.Int, error) {
	res, err := r.DB.NamedQueryContext(c, `INSERT INTO mind(topic, caption, parent_id, access, 
hashed_id, created_by) VALUES (:topic, :caption, :parent_id, :access, :hashed_id, :created_by) 
RETURNING id`, input)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	if !res.Next() {
		return nil, errors.New("500")
	}
	id := new(hash.Int)
	err = res.Scan(id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *Repo) Update(c ctx.Ctx, input *mind.Update) error {
	argNum, args, setValues := 1, []any{}, []string{}
	repoutill.UpdateSetColumn(input.Topic, "topic", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.Caption, "caption", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.ParentId, "parent_id", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.Access, "access", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.HashedId, "hashed_id", &setValues, &args, &argNum)

	if argNum == 1 {
		return errors.New("field not found for updating")
	}
	setStr := strings.Join(setValues, " ,")
	query := fmt.Sprintf(`UPDATE mind SET %s WHERE id= %d AND deleted_at IS NULL AND created_by=%d`,
		setStr, input.Id, *input.CreatedBy)
	_, err := r.DB.ExecContext(c, query, args...)
	return err
}

func (r *Repo) ChildrenById(c ctx.Ctx, filter *mind.ChildrenFilter, getOwnSelf bool) ([]mind.List, error) {
	list := make([]mind.List, 0, 8)
	whereOwn, createdBy := "", 0
	if getOwnSelf {
		whereOwn = "OR id=$1"
	}
	if filter.UserId != nil {
		createdBy = int(*filter.UserId)
	}

	err := r.DB.SelectContext(c, &list,
		fmt.Sprintf(`SELECT id, topic, caption, parent_id, access, hashed_id FROM mind 
WHERE parent_id=$1 %s AND (created_by=$2 OR access = 99)`, whereOwn), filter.MindId, createdBy)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Repo) Delete(c ctx.Ctx, userId hash.Int, deletedBy hash.Int) error {
	_, err := r.DB.ExecContext(c, `UPDATE users SET deleted_at = now(), deleted_by = $1
	 WHERE id = $2`, deletedBy, userId)
	if err != nil {
		return err
	}

	return err
}

//func (r *Repo) filter(f *user.Filter) (*[]model.user, *bun.SelectQuery) {
//	ms, q := r.Filter(f.Filter)
//
//	if f.Username != nil {
//		WhereGroupAnd(q, "username = ?", *f.Username)
//	}
//	if f.Email != nil {
//		WhereGroupAnd(q, "email = ?", *f.Email)
//	}
//
//	return ms, q
//}

//func (r *Repo) GetAll(c ctx.Ctx, f *user.Filter) ([]model.User, int, error) {
//	//ms, query := r.filter(f)
//	//
//	//count, err := query.ScanAndCount(c)
//	//
//	//return *ms, count, err
//	return nil, 0, nil
//}
//
//func (r *Repo) GetById(c ctx.Ctx, id hash.Int) (*model.User, error) {
//	m := new(model.Mind)
//
//	err := r.DB.GetContext(c, m, "SELECT * FROM users WHERE id= $1", id)
//	if err != nil {
//		return nil, err
//	}
//
//	return m, nil
//}
