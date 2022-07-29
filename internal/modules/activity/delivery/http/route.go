package http

import (
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/docs/swagger"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/docs/swagger/endpoint"
	"micro-base/internal/modules/activity"
	activityMysql "micro-base/internal/modules/activity/respository/mysql"
)

var activityService = activity.New(activityMysql.New())

// Endpoints 路由配置
func Endpoints() []*swagger.Endpoint {
	return []*swagger.Endpoint{
		endpoint.Post("/activity/sync", "活动数据同步", endpoint.Handler(syncActivity)),
	}
}
