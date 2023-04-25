package auth

import (
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type User interface {
	GetById(ctx.Ctx, hash.Int) (*model.User, error)
	GetByEmail(ctx.Ctx, string) (*model.User, error)
	GetByUsername(ctx.Ctx, string) (*model.User, error)
	Create(ctx.Ctx, *user.UserCreate) (*model.User, error)
	CreateWithMind(c ctx.Ctx, input *user.UserCreate) (hash.Int, error)
	Available(c ctx.Ctx, column, value string) (bool, error)
}
