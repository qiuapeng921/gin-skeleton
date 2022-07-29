package activity

import (
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ctx"
)

// Service 活动服务
type Service struct {
	repo Repository
}

// SyncActivityData 同步活动数据
func (s *Service) SyncActivityData(c ctx.Context, request SyncRequest) error {
	return nil
}
