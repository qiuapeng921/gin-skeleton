package models

import (
	"gorm.io/gorm"
	"micro-base/internal/pkg/core/db"
	"time"
)

// BaseModel 基础模型，大部分表都有的字段
type BaseModel struct {
	ID        int64        `gorm:"primaryKey;column:id;type:bigint(20)" json:"id"`
	CreatedAt int64        `gorm:"column:created_at;type:int(10);not null;default:0" json:"created_at"`
	UpdatedAt int64        `gorm:"column:updated_at;type:int(10);not null;default:0" json:"updated_at"`
	DeletedAt db.DeletedAt `gorm:"softDelete;column:deleted_at;type:int(10) unsigned;not null;default:0" json:"deleted_at"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.CreatedAt == 0 {
		b.CreatedAt = time.Now().Unix()
	}
	if b.UpdatedAt == 0 {
		b.UpdatedAt = time.Now().Unix()
	}
	return
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now().Unix()
	return
}

func CreateBase(Id int64) BaseModel {
	nowUnix := time.Now().Unix()
	return BaseModel{
		ID:        Id,
		CreatedAt: nowUnix,
		UpdatedAt: nowUnix,
	}
}
