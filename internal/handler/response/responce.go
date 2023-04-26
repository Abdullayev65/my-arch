package response

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, res any) {
	json200(c, map[string]interface{}{
		"res":    res,
		"status": true,
	})
}

func SuccessList(c *gin.Context, res any, count int) {
	json200(c, map[string]interface{}{
		"res":       res,
		"status":    true,
		"count":     count,
		"last_page": (count + 9) / 10,
	})
}

func Response[T any](c *gin.Context, t T, err error) {
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, t)
}

func Fail(c *gin.Context, msg string) {
	json200(c, map[string]interface{}{
		"status": false,
		"msg":    msg,
	})
}

func FailErr(c *gin.Context, err error) {
	Fail(c, err.Error())
}

func json200(c *gin.Context, output any) {
	c.JSON(200, output)
}
