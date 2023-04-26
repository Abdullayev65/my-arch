package user

import (
	"mime/multipart"
	"my-arch/internal/dto/user"
	"my-arch/internal/model"
	"my-arch/pkg/ctx"
)

type User interface {
	Create(ctx.Ctx, *user.UserCreate) error
	GetById(ctx.Ctx, int) (*user.UserDetail, error)
	Update(c ctx.Ctx, input *user.UserUpdate) error
	Delete(c ctx.Ctx, input *user.Delete) error
}

type SysFile interface {
	UploadFile(file *multipart.FileHeader, folder string) (*model.FileData, error)
}

type File interface {
	//Create(c ctx.Ctx, input *model.FileData) (err error)
}
