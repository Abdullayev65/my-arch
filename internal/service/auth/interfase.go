package auth

import (
	"my-arch/internal/dto/user"
	"my-arch/internal/model"
	"my-arch/pkg/ctx"
)

type User interface {
	GetByEmail(ctx.Ctx, string) (*model.User, error)
	GetByUsername(ctx.Ctx, string) (*model.User, error)
	Create(ctx.Ctx, *user.UserCreate) error
}
