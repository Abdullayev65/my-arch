package file

type Service struct {
	iFile
}

func New(file iFile) *Service {
	return &Service{file}
}

//
//func (s *Service) CreateWithMind(c ctx.Ctx, input *file.CreateWithMind) ([]*file.List, error) {
//	var errStr string
//	switch {
//	case input.CreatedBy == nil:
//		errStr = "owner not found"
//	case input.MindId == nil:
//		errStr = "mind_id can't be null"
//	case len(input.Files) == 0:
//		errStr = "file not given"
//	}
//	if errStr != "" {
//		return nil, errors.New(errStr)
//	}
//
//	if input.Access != 33 && input.Access != 99 {
//		input.Access = 99
//	}
//
//	files, err := s.sysFile.MultipleUploadFile(input.Files, "mind-file")
//	if err != nil {
//		return nil, err
//	}
//
//	for _, f := range files {
//		f.CreatedBy = *input.CreatedBy
//		f.Access = input.Access
//	}
//
//	err = s.iFile.CreateWithMind(c, files, *input.MindId)
//	if err != nil {
//		return nil, err
//	}
//
//	list := make([]*file.List, len(files))
//	for i, f := range files {
//		list[i] = f.MapToList()
//	}
//
//	return list, nil
//}
//
//func (s *Service) GetByMindIds(c ctx.Ctx, mindIds []int) (map[int][]file.List, error) {
//	fileList, err := s.iFile.GetByMindIds(c, mindIds)
//	if err != nil {
//		return nil, err
//	}
//
//	for i := range fileList {
//		fileList[i].Url = filepath.Join(config.GetFilesBaseUrl(), strconv.Itoa(fileList[i].Id))
//	}
//
//	fileMap := stream.SliceToMap(fileList, func(f file.List) int {
//		f.Url = filepath.Join(config.GetFilesBaseUrl(), strconv.Itoa(f.Id))
//		return f.MindId
//	})
//
//	return fileMap, nil
//}
