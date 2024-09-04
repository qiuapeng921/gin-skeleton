package user

import (
	"github.com/gin-gonic/gin"
	"micro-base/internal/app/handle"
	"micro-base/internal/app/models"
	"micro-base/internal/app/repo"
)

type User struct {
	handle.BaseHandle
}

func NewUserHandle() User {
	return User{}
}

func (h User) List(c *gin.Context) {
	many, err := repo.NewMember().GetMany(h.WarpContext(c), nil)
	if err != nil {
		h.ResError(c, err)
		return
	}
	h.ResSuccess(c, many)
	return
}

func (h User) Create(c *gin.Context) {
	data := models.MemberModel{Userid: "123"}
	err := repo.NewMember().CreateData(h.WarpContext(c), data)
	h.ResSuccess(c, err)
	return
}
