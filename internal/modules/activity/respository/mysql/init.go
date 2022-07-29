package mysql

import (
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/repos"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/sqlplus"
	"micro-base/internal/modules/activity"
)

func New() activity.Repository {
	return &ActivityRepo{repos.NewSQLRepo("t_shop_activity", sqlplus.MustOpen("mysql", ""))}
}
