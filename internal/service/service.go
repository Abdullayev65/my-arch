package service

import (
	"my-arch/internal/repository"
	"my-arch/internal/service/auth"
	"my-arch/internal/service/file"
	"my-arch/internal/service/user"
)

var (
	Auth       = new(auth.Service)
	User       = new(user.Service)
	File       = new(file.Service)
	SystemFile = new(file.SystemFile)
)

func init() {
	*Auth = *auth.New(repo.User)
	*User = *user.New(repo.User, SystemFile, repo.File)
	*File = *file.New(repo.File)
	*SystemFile = *file.NewSystemFile()
}
