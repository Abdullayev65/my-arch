package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
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

func (r *Repo) GetAll(c ctx.Ctx, f *user.Filter) ([]model.User, int, error) {
	//ms, query := r.filter(f)
	//
	//count, err := query.ScanAndCount(c)
	//
	//return *ms, count, err
	return nil, 0, nil
}

func (r *Repo) GetById(c ctx.Ctx, id hash.Int) (*model.User, error) {
	m := new(model.User)

	err := r.DB.GetContext(c, m, "SELECT * FROM users WHERE id= $1", id)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Repo) Create(c ctx.Ctx, input *user.UserCreate) (*model.User, error) {
	res, err := r.DB.NamedExecContext(c, `INSERT INTO users(username, email, password, 
first_name, middle_name, last_name, birth_date) VALUES (:username, :email, :password,
:first_name, :middle_name, :last_name, :birth_date) RETURNING id`, input)
	if err != nil {
		return nil, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	println(id)
	return nil, nil
}
func (r *Repo) CreateWithMind(c ctx.Ctx, input *user.UserCreate) (hash.Int, error) {
	txx, err := r.DB.BeginTxx(c, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer txx.Commit()

	res, err := txx.NamedQuery(`INSERT INTO users(username, email, password, 
first_name, middle_name, last_name, birth_date) VALUES (:username, :email, :password,
:first_name, :middle_name, :last_name, :birth_date) RETURNING id`, input)
	if err != nil {
		txx.Rollback()
		return 0, err
	}

	if !res.Next() {
		txx.Rollback()
		return 0, errors.New("500")
	}
	var id int64
	err = res.Scan(&id)
	res.Close()
	if err != nil {
		return 0, err
	}

	var mindId int64
	err = txx.QueryRowx(`INSERT INTO mind(topic, created_by, access) 
VALUES ($1, $2, 99) RETURNING id`, "root", id).Scan(&mindId)
	if err != nil {
		txx.Rollback()
		return 0, err
	}

	_, err = txx.ExecContext(c, `UPDATE users SET mind_id = $1 WHERE id=$2`,
		mindId, id)
	if err != nil {
		txx.Rollback()
		return 0, err
	}

	return hash.Int(id), nil
}

func (r *Repo) Available(c ctx.Ctx, column, value string) (bool, error) {
	type Data struct {
		Count int
	}
	d := new(Data)

	err := r.DB.GetContext(c, d, fmt.Sprintf(
		`SELECT count(id) FROM users WHERE deleted_at IS NULL AND %s=$1`, column), value)
	if err != nil {
		return false, err
	}

	if d.Count == 0 {
		return true, nil
	}

	return false, fmt.Errorf("%s is invalid or already taken", column)
}

func (r *Repo) Update(c ctx.Ctx, input *user.UserUpdate) error {
	argNum, args, setValues := 1, []any{}, []string{}
	repoutill.UpdateSetColumn(input.Username, "username", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.Email, "email", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.Password, "password", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.FirstName, "first_name", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.MiddleName, "middle_name", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.LastName, "last_name", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.BirthDate, "birth_date", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.AvatarId, "avatar_id", &setValues, &args, &argNum)

	if argNum == 1 {
		return errors.New("field not found for updating")
	}
	setStr := strings.Join(setValues, " ,")
	query := fmt.Sprintf(`UPDATE users SET %s WHERE id= %d AND deleted_at IS NULL`,
		setStr, input.Id)
	_, err := r.DB.ExecContext(c, query, args...)
	return err
}

func (r *Repo) GetByUsername(c ctx.Ctx, username string) (*model.User, error) {
	return r.getBy(c, "username", username)
}

func (r *Repo) GetByEmail(c ctx.Ctx, email string) (*model.User, error) {
	return r.getBy(c, "email", email)
}

func (r *Repo) UserSearch(c ctx.Ctx, input *user.UserSearch) ([]*user.UserList, int, error) {
	list := make([]*user.UserList, 0, 10)
	count := new(int)

	err := r.DB.SelectContext(c, &list, `SELECT id, username, mind_id, first_name, middle_name, last_name
	FROM users WHERE deleted_at IS NULL AND username LIKE $1
	LIMIT $2 OFFSET $3`, "%"+input.Username+"%", input.Limit, input.Offset)
	if err != nil {
		return nil, 0, err
	}

	err = r.DB.GetContext(c, count, `SELECT count(id)
FROM users WHERE deleted_at IS NULL AND username LIKE $1`, "%"+input.Username+"%")
	if err != nil {
		return nil, 0, err
	}

	return list, *count, nil
}

func (r *Repo) DetailById(c ctx.Ctx, id *hash.Int) (*user.UserDetail, error) {
	o := new(user.UserDetail)

	err := r.DB.GetContext(c, o,
		`SELECT id, username, email, mind_id, first_name, 
middle_name, last_name, birth_date, avatar_id FROM users WHERE id=$1`, id)

	if err != nil {
		return nil, err
	}

	return o, nil
}

func (r *Repo) Delete(c ctx.Ctx, userId hash.Int, deletedBy hash.Int) error {
	_, err := r.DB.ExecContext(c, `UPDATE users SET deleted_at = now(), deleted_by = $1
	 WHERE id = $2`, deletedBy, userId)
	if err != nil {
		return err
	}

	return err
}

// specific functions

func (r *Repo) getBy(c ctx.Ctx, column string, arg any) (*model.User, error) {
	m := new(model.User)

	err := r.DB.GetContext(c, m, "SELECT * FROM users WHERE deleted_at IS NULL AND "+
		column+"= $1 ", arg)
	if err != nil {
		return nil, err
	}

	return m, nil
}
