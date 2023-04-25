package auth

import (
	"github.com/gin-gonic/gin"
	"my-arch/internal/dto/auth"
	"my-arch/internal/dto/user"
	. "my-arch/internal/handler/response"
	auth_srvc "my-arch/internal/service/auth"
)

type Handler struct {
	auth *auth_srvc.Service
}

func New(auth *auth_srvc.Service) *Handler {
	return &Handler{auth}
}

func (h *Handler) SignUp(c *gin.Context, input *user.UserCreate) {

	err := h.auth.SignUp(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "user created")
}

func (h *Handler) LogIn(c *gin.Context, input *auth.LogIn) {
	outPut, err := h.auth.LogIn(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, outPut)
}
