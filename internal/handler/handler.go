package handler

import (
	"my-arch/internal/handler/mw"
	"my-arch/internal/handler/v1/auth"
	"my-arch/internal/handler/v1/file"
	"my-arch/internal/handler/v1/user"
	"my-arch/internal/service"
)

var (
	MW   = new(mw.MiddleWere)
	Auth = new(auth.Handler)
	User = new(user.Handler)
	File = new(file.Handler)
)

func init() {
	*MW = *mw.New(service.Auth)
	*Auth = *auth.New(service.Auth)
	*User = *user.New(service.User, MW)
	*File = *file.New(MW, service.File)
}
