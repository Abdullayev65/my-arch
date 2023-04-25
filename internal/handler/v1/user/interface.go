package user

import (
	"github.com/gin-gonic/gin"
	"mindstore/pkg/hash-types"
)

type AuthMW interface {
	GetUserId(c *gin.Context) (id *hash.Int, ok bool)
	MustGetUserId(c *gin.Context) *hash.Int
}
