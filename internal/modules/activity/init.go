package activity

// New Service初始化
func New(repository Repository) *Service {
	return &Service{
		repo: repository,
	}
}
