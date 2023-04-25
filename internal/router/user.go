package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
	"mindstore/pkg/bind"
)

func User(r *gin.RouterGroup) {
	h := handler.User
	mw := handler.MW

	r.GET("me", mw.UserIdFromToken(true), h.UserGetMe)
	r.PUT("me", mw.UserIdFromToken(true), bind.Binder(h.UserUpdate))
	r.DELETE("me", mw.UserIdFromToken(true), h.UserDelete)
	r.GET("search/:username", h.UserSearch)
}
