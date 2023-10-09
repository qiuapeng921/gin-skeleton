package v1

import (
	"github.com/gin-gonic/gin"
	"micro-base/internal/app/handle/user"
)

func InitRouter(router *gin.RouterGroup) {

	userGroup := router.Group("user")
	{
		userHandle := user.NewUserHandle()
		userGroup.GET("list", userHandle.List)
	}
}
