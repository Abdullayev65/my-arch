package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
	"mindstore/pkg/bind"
)

func Auth(r *gin.RouterGroup) {
	h := handler.Auth

	r.POST("sign-up", bind.Binder(h.SignUp))
	r.POST("log-in", bind.Binder(h.LogIn))
	r.POST("available", bind.Binder(h.Available))
}
