package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
)

func File(r *gin.RouterGroup) {
	h := handler.File
	mw := handler.MW

	r.GET("/:id", mw.UserIdFromToken(false), h.GetFile)
	r.GET("/avatar/:user_id", h.GetAvatar)
}
