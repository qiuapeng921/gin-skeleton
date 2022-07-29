package mysql

import (
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ctx"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/repos"
)

// ActivityRepo 活动仓库实现
type ActivityRepo struct {
	*repos.SQLRepo
}

// GetActivity 获取活动信息
func (a *ActivityRepo) GetActivity(c ctx.Context, enterpriseID int64, activityID int64) (interface{}, error) {
	var data interface{}
	return data, nil
}
