package user

import (
	"errors"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/config"
	"mindstore/pkg/ctx"
	"mindstore/pkg/encoder"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/util/timeutil"
	"path"
)

type Service struct {
	user       User
	auth       Auth
	systemFile SysFile
	file       File
}

func New(user User, auth Auth, systemFile SysFile, file File) *Service {
	n := new(Service)

	n.user = user
	n.auth = auth
	n.systemFile = systemFile
	n.file = file

	return n
}

func (s *Service) UserById(c ctx.Ctx, id hash.Int) (*model.User, error) {
	return s.user.GetById(c, id)
}

func (s *Service) DetailById(c ctx.Ctx, id *hash.Int) (*user.UserDetail, error) {
	obj, err := s.user.DetailById(c, id)
	if err != nil {
		return nil, err
	}

	timeutil.Format(obj.BirthDate, &obj.BirthDateStr)
	if obj.AvatarId != nil {
		url := path.Join(config.GetFilesBaseUrl(), "avatar", obj.Id.HashToStr())
		obj.AvatarUrl = &url
	}

	return obj, nil
}

func (s *Service) UserUpdate(c ctx.Ctx, input *user.UserUpdate) error {
	var errStr string
	switch {
	case input.Email != nil && (!s.auth.IsValidEmail(*input.Email)):
		errStr = "email is not valid"
	case input.Username != nil && s.auth.IsValidUsername(*input.Username) != nil:
		errStr = s.auth.IsValidUsername(*input.Username).Error()

	case input.Password != nil && (len(*input.Password) < 1 || len(*input.Password) > 30):
		errStr = "password length should be between 1 and 30"
	}
	if errStr != "" {
		return errors.New(errStr)
	}

	if input.Password != nil {
		password, err := encoder.HashPassword(*input.Password)
		if err != nil {
			return err
		}

		input.Password = &password
	}
	if input.Avatar != nil {
		fileData, err := s.systemFile.UploadFile(input.Avatar, "avatar")
		if err != nil {
			return err
		}
		fileData.CreatedBy = input.Id
		err = s.file.Create(c, fileData)
		if err != nil {
			return err
		}
		input.AvatarId = &fileData.Id
	}

	timeutil.Parse(input.BirthDateStr, &input.BirthDate)

	return s.user.Update(c, input)
}

func (s *Service) Delete(c ctx.Ctx, userId, deletedBy hash.Int) error {
	return s.user.Delete(c, userId, deletedBy)
}

func (s *Service) UserSearch(c ctx.Ctx, input *user.UserSearch) ([]*user.UserList, int, error) {
	return s.user.UserSearch(c, input)
}
