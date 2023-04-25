package user

import (
	"mime/multipart"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type User interface {
	GetByUsername(c ctx.Ctx, username string) (*model.User, error)
	GetByEmail(c ctx.Ctx, email string) (*model.User, error)
	Create(ctx.Ctx, *user.UserCreate) (*model.User, error)
	GetById(ctx.Ctx, hash.Int) (*model.User, error)
	DetailById(c ctx.Ctx, id *hash.Int) (*user.UserDetail, error)
	Update(c ctx.Ctx, input *user.UserUpdate) error
	Delete(c ctx.Ctx, userId hash.Int, deletedBy hash.Int) error
	UserSearch(c ctx.Ctx, input *user.UserSearch) ([]*user.UserList, int, error)
}

type Auth interface {
	IsValidEmail(email string) bool
	IsValidUsername(username string) error
}

type SysFile interface {
	UploadFile(file *multipart.FileHeader, folder string) (*model.FileData, error)
}

type File interface {
	Create(c ctx.Ctx, input *model.FileData) (err error)
}
