package activity

import (
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/repos"
)

// Repository 存储定义
type Repository interface {
	repos.IRepository
}
