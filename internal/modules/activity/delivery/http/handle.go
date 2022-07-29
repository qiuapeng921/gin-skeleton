package http

import (
	"github.com/gin-gonic/gin"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ctx"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ginplus"
	"micro-scrm/internal/modules/activity"
)

func syncActivity(c *gin.Context) {
	var request activity.SyncRequest
	err := c.BindJSON(&request)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	err = activityService.SyncActivityData(ctx.Wrap(c), request)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nil)
}
