package repo

import (
	"context"
	"gorm.io/gorm"
)

type Dao struct {
	*gorm.DB
}

// QueryClone 克隆一个带原条件的新实例,WithContext必须放在Session后，否则不会带上之前的条件
func (d *Dao) QueryClone(c context.Context, query *gorm.DB) *gorm.DB {
	return query.Session(&gorm.Session{NewDB: true}).WithContext(c)
}
