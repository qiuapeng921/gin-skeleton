package repo

import (
	"fmt"
	"gorm.io/gorm"
	"micro-base/internal/app"
	"micro-base/internal/app/models"
)

type Member struct {
	dao   Dao
	model *models.MemberModel
}

func NewMember() *Member {
	return &Member{
		dao:   Dao{DB: app.DB()},
		model: new(models.MemberModel),
	}
}

// SetDb 切换db连接,用于事物
func (d *Member) SetDb(db *gorm.DB) *Member {
	d.dao.DB = db
	return d
}

func (d *Member) GetOne(c app.Ctx, query any, args ...any) (models.MemberModel, error) {
	var result models.MemberModel
	err := d.dao.WithContext(c).Model(d.model).Where(query, args...).Find(&result).Error
	return result, err
}

func (d *Member) GetMany(c app.Ctx, query any, args ...any) ([]models.MemberModel, error) {
	var result []models.MemberModel
	err := d.dao.WithContext(c).Model(d.model).Where(query, args...).Find(&result).Error
	return result, err
}

func (d *Member) CreateData(c app.Ctx, data models.MemberModel) error {
	err := d.dao.WithContext(c).Model(d.model).Create(&data).Error
	return err
}

func (d *Member) CreateInBatchesData(c app.Ctx, data []models.MemberModel) error {
	err := d.dao.WithContext(c).Model(d.model).CreateInBatches(&data, len(data)).Error
	return err
}

func (d *Member) SaveData(c app.Ctx, data models.MemberModel) error {
	err := d.dao.WithContext(c).Model(d.model).Save(&data).Error
	return err
}

func (d *Member) UpdateData(c app.Ctx, query any, data map[string]any) error {
	err := d.dao.WithContext(c).Model(d.model).Where(query).UpdateColumns(data).Error
	return err
}

func (d *Member) DeleteData(c app.Ctx, query any, args ...any) error {
	var model models.MemberModel
	err := d.dao.WithContext(c).Model(d.model).Where(query, args...).Delete(&model).Error
	return err
}

func (d *Member) ForceDelete(c app.Ctx, query any, args ...any) error {
	var model models.MemberModel
	err := d.dao.WithContext(c).Model(d.model).Where(query, args...).Unscoped().Delete(&model).Error
	return err
}

func (d *Member) Increment(c app.Ctx, query any, column string, num int64) error {
	err := d.dao.WithContext(c).
		Model(d.model).
		Where(query).
		UpdateColumn(column, gorm.Expr(fmt.Sprintf("%s + ?", column), num)).Error
	return err
}

func (d *Member) Decrement(c app.Ctx, query any, column string, num int64) error {
	err := d.dao.WithContext(c).
		Model(d.model).
		Where(query).
		UpdateColumn(column, gorm.Expr(fmt.Sprintf("%s - ?", column), num)).Error
	return err
}
