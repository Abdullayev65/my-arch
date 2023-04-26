package file

import (
	"github.com/jmoiron/sqlx"
	"my-arch/internal/dto/file"
	"my-arch/pkg/ctx"
)

type Repo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(c ctx.Ctx, input *file.Create) error {

	return nil
}

func (r *Repo) GetAll(c ctx.Ctx, filter *file.Filter) ([]file.List, error) {

	return nil, nil
}

func (r *Repo) GetById(c ctx.Ctx, id int) (*file.Detail, error) {

	return nil, nil
}

func (r *Repo) Update(c ctx.Ctx, input *file.Update) error {
	return nil
}

func (r *Repo) Delete(c ctx.Ctx, input *file.Delete) error {
	return nil
}
