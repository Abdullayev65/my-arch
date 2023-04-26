package file

import (
	"github.com/gin-gonic/gin"
)

type AuthMW interface {
	GetUserId(c *gin.Context) (id int, ok bool)
	MustGetUserId(c *gin.Context) int
}
