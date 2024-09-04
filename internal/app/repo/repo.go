package repo

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"micro-base/internal/app"
)

type Repo[T schema.Tabler] struct {
	*gorm.DB
	model T
}

// QueryClone 克隆一个带原条件的新实例,WithContext必须放在Session后，否则不会带上之前的条件
func (d *Repo[T]) QueryClone(c context.Context, query *gorm.DB) *gorm.DB {
	return query.Session(&gorm.Session{NewDB: true}).WithContext(c)
}

// SetDb 切换db连接,用于事物
func (d *Repo[T]) SetDb(db *gorm.DB) *Repo[T] {
	d.DB = db
	return d
}

func (d *Repo[T]) GetOne(c app.Ctx, query any, args ...any) (T, error) {
	var result T
	err := d.WithContext(c).Model(d.model).Where(query, args...).Find(&result).Error
	return result, err
}

func (d *Repo[T]) GetMany(c app.Ctx, query any, args ...any) ([]T, error) {
	var result []T
	err := d.WithContext(c).Model(d.model).Where(query, args...).Find(&result).Error
	return result, err
}

func (d *Repo[T]) CreateData(c app.Ctx, data T) error {
	err := d.WithContext(c).Model(d.model).Create(&data).Error
	return err
}

func (d *Repo[T]) CreateInBatchesData(c app.Ctx, data []T) error {
	err := d.WithContext(c).Model(d.model).CreateInBatches(&data, len(data)).Error
	return err
}

func (d *Repo[T]) SaveData(c app.Ctx, data T) error {
	err := d.WithContext(c).Model(d.model).Save(&data).Error
	return err
}

func (d *Repo[T]) UpdateData(c app.Ctx, query any, data map[string]any) error {
	err := d.WithContext(c).Model(d.model).Where(query).UpdateColumns(data).Error
	return err
}

func (d *Repo[T]) DeleteData(c app.Ctx, query any, args ...any) error {
	var model T
	err := d.WithContext(c).Model(d.model).Where(query, args...).Delete(&model).Error
	return err
}

func (d *Repo[T]) ForceDelete(c app.Ctx, query any, args ...any) error {
	var model T
	err := d.WithContext(c).Model(d.model).Where(query, args...).Unscoped().Delete(&model).Error
	return err
}

func (d *Repo[T]) Increment(c app.Ctx, query any, column string, num int64) error {
	err := d.WithContext(c).
		Model(d.model).
		Where(query).
		UpdateColumn(column, gorm.Expr(fmt.Sprintf("%s + ?", column), num)).Error
	return err
}

func (d *Repo[T]) Decrement(c app.Ctx, query any, column string, num int64) error {
	err := d.WithContext(c).
		Model(d.model).
		Where(query).
		UpdateColumn(column, gorm.Expr(fmt.Sprintf("%s - ?", column), num)).Error
	return err
}
