package user

import (
	"github.com/jmoiron/sqlx"
	"my-arch/internal/dto/user"
	"my-arch/internal/model"
	"my-arch/pkg/ctx"
)

type Repo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(c ctx.Ctx, input *user.UserCreate) error {

	return nil
}

func (r *Repo) GetAll(c ctx.Ctx, filter *user.Filter) ([]user.UserList, error) {

	return nil, nil
}

func (r *Repo) GetById(c ctx.Ctx, id int) (*user.UserDetail, error) {

	return nil, nil
}

func (r *Repo) Update(c ctx.Ctx, input *user.UserUpdate) error {
	return nil
}

func (r *Repo) Delete(c ctx.Ctx, input *user.Delete) error {
	return nil
}

// specific functions

func (r *Repo) GetByEmail(c ctx.Ctx, s string) (*model.User, error) {
	return r.getBy(c, "email", s)
}

func (r *Repo) GetByUsername(c ctx.Ctx, s string) (*model.User, error) {
	return r.getBy(c, "username", s)
}

func (r *Repo) getBy(c ctx.Ctx, column string, arg any) (*model.User, error) {
	m := new(model.User)

	err := r.DB.GetContext(c, m, "SELECT * FROM users WHERE deleted_at IS NULL AND "+
		column+"= $1 ", arg)
	if err != nil {
		return nil, err
	}

	return m, nil
}
