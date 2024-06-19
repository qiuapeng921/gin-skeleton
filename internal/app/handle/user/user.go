package user

import (
	"github.com/gin-gonic/gin"
	"micro-base/internal/app/handle"
)

type User struct {
	handle.BaseHandle
}

func NewUserHandle() User {
	return User{}
}

func (h User) List(c *gin.Context) {
	h.ResSuccess(c, "user-list")
	return
}
