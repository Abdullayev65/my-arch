package mw

import (
	"github.com/gin-gonic/gin"
	"my-arch/internal/handler/response"
)

func (mw *MiddleWere) ErrorHandler(c *gin.Context) {
	defer func() {
		if a := recover(); a != nil {
			err, hasErr := a.(error)
			msg, hasMsg := a.(string)
			switch {
			case hasErr:
				msg = err.Error()
			case !hasMsg:
				msg = "something gone wrong"
			}

			response.Fail(c, msg)
		}
	}()

	c.Next()
}
