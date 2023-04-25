package file

import (
	"github.com/gin-gonic/gin"
	file_srvc "my-arch/internal/service/file"
)

type Handler struct {
	authMW AuthMW
	file   *file_srvc.Service
}

func New(authMW AuthMW, file *file_srvc.Service) *Handler {
	return &Handler{authMW, file}
}

func (h *Handler) GetFile(c *gin.Context) {
	userId := int(0)

	if userIdPtr, ok := h.authMW.GetUserId(c); ok {
		userId = *userIdPtr
	}

	var fileId int
	err := fileId.UnhashStr(c.Param("id"))
	if err != nil {
		FailErr(c, err)
		return
	}

	path, err := h.file.GetPathById(c, fileId, userId)
	if err != nil {
		FailErr(c, err)
		return
	}

	c.File(path)
}
