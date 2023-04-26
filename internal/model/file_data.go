package model

import (
	"my-arch/internal/dto/file"
	"my-arch/internal/model/submodel"
	"my-arch/internal/pkg/config"
	"path"
	"strconv"
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
		Url:      path.Join(config.GetFilesBaseUrl(), strconv.Itoa(f.Id)),
		Name:     f.Name,
		HashedId: f.HashedId,
		Access:   f.Access,
		Size:     f.Size,
	}
}
