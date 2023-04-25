package service

import (
	"mindstore/internal/repository"
	"mindstore/internal/service/auth"
	"mindstore/internal/service/file"
	"mindstore/internal/service/mind"
	"mindstore/internal/service/user"
)

var (
	Auth       = new(auth.Service)
	User       = new(user.Service)
	Mind       = new(mind.Service)
	File       = new(file.Service)
	SystemFile = new(file.SystemFile)
)

func init() {
	*Auth = *auth.New(repo.User)
	*User = *user.New(repo.User, Auth, SystemFile, repo.File)
	*Mind = *mind.New(repo.Mind, File)
	*File = *file.New(repo.File, SystemFile)
	*SystemFile = *file.NewSystemFile()
}
