package file

import (
	"mime/multipart"
	"mindstore/internal/object/dto/file"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type iFile interface {
	CreateWithMind(c ctx.Ctx, files []*model.FileData, mindId hash.Int) error
	GetByMindIds(c ctx.Ctx, mindIds []hash.Int) ([]file.List, error)
	Delete(c ctx.Ctx, input *file.DeleteMind) (err error)
	GetPathById(c ctx.Ctx, fileId, userId hash.Int) (path string, err error)
	GetAvatarPathByUserId(c ctx.Ctx, userId hash.Int) (path string, err error)
}

type SysFile interface {
	MultipleUploadFile(files []*multipart.FileHeader, folder string) (fds []*model.FileData, err error)
	UploadFile(file *multipart.FileHeader, folder string) (*model.FileData, error)
}
