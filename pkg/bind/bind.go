package bind

import (
	"github.com/gin-gonic/gin"
	"my-arch/internal/handler/response"
)

func Binder[In any](handler func(*gin.Context, *In)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(In)

		err := c.ShouldBind(in)
		if err != nil {
			response.FailErr(c, err)
			c.Abort()
			return
		}

		handler(c, in)
	}
}
