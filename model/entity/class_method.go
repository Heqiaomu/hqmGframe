package entity

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (dao *DAO) CreateClass(class Class) error {
	classDB := dao.W.Table("tb_class").Create(&class)

	return classDB.Error
}

func (dao *DAO) GetClass(className string) (*Class, error) {
	class := Class{}
	err := dao.R.Table("tb_class").Where("class_name = ?", className).Find(&class).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &class, nil
}

func (dao *DAO) GetClassList(page Page) ([]*Class, int64, error) {
	var class []*Class
	var total int64
	err := dao.R.Table("tb_class").Where("deleted_at is NOT NULL").
		Offset(page.Offset).
		Limit(page.Size).
		Find(&class).Error

	err = dao.R.Table("tb_class").Where("deleted_at is NOT NULL").Count(&total).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return class, total, nil
}
