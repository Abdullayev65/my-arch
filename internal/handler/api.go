package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"my-arch/pkg/bind"
	"time"
)

func InitApi() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}), MW.ErrorHandler)

	v1 := r.Group("/api/v1")

	{
		auth := v1.Group("/auth")

		auth.POST("sign-up", bind.Binder(Auth.SignUp))
		auth.POST("log-in", bind.Binder(Auth.LogIn))
	}

	{
		user := v1.Group("/user")

		user.GET("me", MW.UserIdFromToken(true), User.UserGetMe)
		user.PUT("me", MW.UserIdFromToken(true), bind.Binder(User.UserUpdate))
		user.DELETE("me", MW.UserIdFromToken(true), User.UserDelete)
	}

	{
		file := v1.Group("/file")

		file.GET("/:id", MW.UserIdFromToken(false), File.GetFile)
	}

	return r
}
