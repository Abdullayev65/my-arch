package user

import (
	"errors"
	"my-arch/internal/dto/user"
	"my-arch/internal/model"
	"my-arch/internal/tools/password_tool"
	"my-arch/internal/tools/time_tool"
	"my-arch/internal/tools/valid_tool"
	"my-arch/pkg/ctx"
)

type Service struct {
	user       User
	systemFile SysFile
	file       File
}

func New(user User, systemFile SysFile, file File) *Service {
	n := new(Service)

	n.user = user
	n.systemFile = systemFile
	n.file = file

	return n
}

func (s *Service) UserById(c ctx.Ctx, id int) (*model.User, error) {
	//return s.user.GetById(c, id)
	return nil, nil
}

func (s *Service) DetailById(c ctx.Ctx, id *int) (*user.UserDetail, error) {
	//obj, err := s.user.DetailById(c, id)
	//if err != nil {
	//	return nil, err
	//}
	//
	//time_tool.Format(obj.BirthDate, &obj.BirthDateStr)
	//if obj.AvatarId != nil {
	//	url := path.Join(config.GetFilesBaseUrl(), "avatar", strconv.Itoa(obj.Id))
	//	obj.AvatarUrl = &url
	//}

	return nil, nil
}

func (s *Service) UserUpdate(c ctx.Ctx, input *user.UserUpdate) error {
	var errStr string
	switch {
	case input.Email != nil && (!valid_tool.IsValidEmail(*input.Email)):
		errStr = "email is not valid"
	case input.Username != nil && valid_tool.IsValidUsername(*input.Username) != nil:
		errStr = valid_tool.IsValidUsername(*input.Username).Error()

	case input.Password != nil && (len(*input.Password) < 1 || len(*input.Password) > 30):
		errStr = "password length should be between 1 and 30"
	}
	if errStr != "" {
		return errors.New(errStr)
	}

	if input.Password != nil {
		password, err := password_tool.HashPassword(*input.Password)
		if err != nil {
			return err
		}

		input.Password = &password
	}
	//if input.Avatar != nil {
	//	fileData, err := s.systemFile.UploadFile(input.Avatar, "avatar")
	//	if err != nil {
	//		return err
	//	}
	//	//fileData.CreatedBy = input.Id
	//	err = s.file.Create(c, fileData)
	//	if err != nil {
	//		return err
	//	}
	//	input.AvatarId = &fileData.Id
	//}

	time_tool.Parse(input.BirthDateStr, &input.BirthDate)

	return s.user.Update(c, input)
}

func (s *Service) Delete(c ctx.Ctx, userId, deletedBy int) error {
	return s.user.Delete(c, nil)
}
