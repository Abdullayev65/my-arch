package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
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
	}), handler.MW.ErrorHandler)

	v1 := r.Group("/api/v1")

	Auth(v1.Group("/auth"))
	User(v1.Group("/user"))
	Mind(v1.Group("/mind"))
	File(v1.Group("/file"))

	return r
}
