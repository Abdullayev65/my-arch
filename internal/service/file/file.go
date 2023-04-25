package file

import (
	"errors"
	"mindstore/internal/object/dto/file"
	"mindstore/internal/object/model"
	"mindstore/pkg/config"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/stream"
	"path/filepath"
)

type Service struct {
	iFile
	sysFile SysFile
}

func New(file iFile, sysFile SysFile) *Service {
	return &Service{file, sysFile}
}

func (s *Service) CreateWithMind(c ctx.Ctx, input *file.CreateWithMind) ([]*file.List, error) {
	var errStr string
	switch {
	case input.CreatedBy == nil:
		errStr = "owner not found"
	case input.MindId == nil:
		errStr = "mind_id can't be null"
	case len(input.Files) == 0:
		errStr = "file not given"
	}
	if errStr != "" {
		return nil, errors.New(errStr)
	}

	if input.Access != 33 && input.Access != 99 {
		input.Access = 99
	}

	files, err := s.sysFile.MultipleUploadFile(input.Files, "mind-file")
	if err != nil {
		return nil, err
	}

	stream.ForEach(files, func(f *model.FileData) {
		f.CreatedBy = *input.CreatedBy
		f.Access = input.Access
	})

	err = s.iFile.CreateWithMind(c, files, *input.MindId)
	if err != nil {
		return nil, err
	}

	list := make([]*file.List, len(files))
	for i, f := range files {
		list[i] = f.MapToList()
	}

	return list, nil
}

func (s *Service) GetByMindIds(c ctx.Ctx, mindIds []hash.Int) (map[hash.Int][]file.List, error) {
	fileList, err := s.iFile.GetByMindIds(c, mindIds)
	if err != nil {
		return nil, err
	}

	stream.ForEach(fileList, func(f file.List) {
		f.Url = filepath.Join(config.GetFilesBaseUrl(), f.Id.HashToStr())
	})

	fileMap := stream.SliceToMap(fileList, func(f file.List) hash.Int {
		f.Url = filepath.Join(config.GetFilesBaseUrl(), f.Id.HashToStr())
		return f.MindId
	})

	return fileMap, nil
}
