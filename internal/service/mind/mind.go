package mind

import (
	"errors"
	"mindstore/internal/object/dto/mind"
	"mindstore/pkg/config"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/stream"
	"path"
)

type Service struct {
	mind Mind
	file File
}

func New(mind Mind, file File) *Service {
	return &Service{mind, file}
}

func (s *Service) CreateMind(c ctx.Ctx, input *mind.Create) (*hash.Int, error) {
	var errStr string
	switch {
	case input.CreatedBy == nil:
		errStr = "owner not found"
	case input.Topic == nil:
		errStr = "topic can't be null"
	case input.ParentId == nil:
		errStr = "parent of mind not given"
	}
	if errStr != "" {
		return nil, errors.New(errStr)
	}

	if input.Access != 33 && input.Access != 99 {
		input.Access = 99
	}

	return s.mind.Create(c, input)
}

func (s *Service) UpdateMind(c ctx.Ctx, input *mind.Update) error {
	var errStr string
	switch {
	case input.CreatedBy == nil:
		errStr = "owner not found"
	case input.Access != nil && (*input.Access != 33 && *input.Access != 99):
		errStr = "access can be 33 or 99"
	}
	if errStr != "" {
		return errors.New(errStr)
	}

	return s.mind.Update(c, input)
}

func (s *Service) ChildrenById(c ctx.Ctx, filter *mind.ChildrenFilter) ([]mind.List, error) {
	mindList, err := s.mind.ChildrenById(c, filter, false)
	if err != nil {
		return nil, err
	}

	return mindList, s.setFilesToMinds(c, mindList)
}

func (s *Service) WithChildrenById(c ctx.Ctx, filter *mind.ChildrenFilter) (*mind.List, error) {
	mindList, err := s.mind.ChildrenById(c, filter, true)
	if err != nil {
		return nil, err
	}

	err = s.setFilesToMinds(c, mindList)
	if err != nil {
		return nil, err
	}

	var root *mind.List
	children := make([]mind.List, 0, len(mindList)-1)

	stream.ForEach(mindList, func(list mind.List) {
		if list.Id == filter.MindId {
			root = &list
		} else {
			children = append(children, list)
		}
	})
	root.Children = children

	return root, nil
}

func (s *Service) setFilesToMinds(c ctx.Ctx, mindList []mind.List) error {
	mindIds := stream.Mapper(mindList, func(m mind.List) hash.Int {
		return m.Id
	})

	fileMap, err := s.file.GetByMindIds(c, mindIds)
	if err != nil {
		return err
	}

	for _, files := range fileMap {
		for i, _ := range files {
			files[i].Url = path.Join(config.GetFilesBaseUrl(), files[i].Id.HashToStr())
		}
	}

	for i, _ := range mindList {
		mindList[i].Files = fileMap[mindList[i].Id]
	}

	return nil
}
