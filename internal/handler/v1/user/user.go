package user

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/object/dto/user"
	user_srvc "mindstore/internal/service/user"
	. "mindstore/pkg/response"
	"strconv"
)

type Handler struct {
	user   *user_srvc.Service
	authMW AuthMW
}

func New(user *user_srvc.Service, authMW AuthMW) *Handler {
	return &Handler{user, authMW}
}

func (h *Handler) UserGetMe(c *gin.Context) {
	userId := h.authMW.MustGetUserId(c)

	detail, err := h.user.DetailById(c, userId)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, detail)
}

func (h *Handler) UserUpdate(c *gin.Context, input *user.UserUpdate) {
	input.Id = *h.authMW.MustGetUserId(c)

	err := h.user.UserUpdate(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "DONE")
}

func (h *Handler) UserDelete(c *gin.Context) {
	userId := *h.authMW.MustGetUserId(c)

	err := h.user.Delete(c, userId, userId)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "DELETED")
}

func (h *Handler) UserSearch(c *gin.Context) {
	input := new(user.UserSearch)
	input.Username = c.Param("username")
	if query, ok := c.GetQuery("page"); ok {
		i, err := strconv.Atoi(query)
		if err == nil {
			input.Limit = 10
			input.Offset = 10 * (i - 1)
		}
	}
	if query, ok := c.GetQuery("limit"); ok {
		i, err := strconv.Atoi(query)
		if err == nil {
			input.Limit = i
		}
	}
	if query, ok := c.GetQuery("offset"); ok {
		i, err := strconv.Atoi(query)
		if err == nil {
			input.Offset = i
		}
	}
	if input.Limit == 0 {
		input.Limit = 10
	}

	list, count, err := h.user.UserSearch(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	SuccessList(c, list, count)
}
