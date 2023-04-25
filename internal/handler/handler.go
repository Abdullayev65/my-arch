package handler

import (
	"mindstore/internal/handler/mw"
	"mindstore/internal/handler/v1/auth"
	"mindstore/internal/handler/v1/file"
	"mindstore/internal/handler/v1/mind"
	"mindstore/internal/handler/v1/user"
	"mindstore/internal/service"
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
	*Mind = *mind.New(MW, service.Mind, service.File)
	*File = *file.New(MW, service.File)
}
