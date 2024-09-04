package repo

import (
	"micro-base/internal/app"
	"micro-base/internal/app/models"
)

type Member struct {
	Repo[models.MemberModel]
}

func NewMember() *Member {
	return &Member{
		Repo: Repo[models.MemberModel]{DB: app.DB()},
	}
}
