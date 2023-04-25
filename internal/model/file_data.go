package model

import (
	"mindstore/internal/object/dto/file"
	"mindstore/internal/object/model/submodel"
	"mindstore/pkg/config"
	"path"
)

type FileData struct {
	submodel.BasicModel

	Path     string
	Name     string
	HashedId *int
	Access   int
	Size     int64
}

func (f *FileData) MapToList() *file.List {
	return &file.List{
		Id:       f.Id,
		Url:      path.Join(config.GetFilesBaseUrl(), f.Id.HashToStr()),
		Name:     f.Name,
		HashedId: f.HashedId,
		Access:   f.Access,
		Size:     f.Size,
	}
}
