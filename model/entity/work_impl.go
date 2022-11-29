package entity

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (dao *DAO) CreateWork(class Work) error {
	wDB := dao.W.Table("tb_work").Create(&class)

	return wDB.Error
}


func (dao *DAO) GetWorkList(page Page) ([]*Work, int64, error)  {
	var class []*Work
	var total int64
	err := dao.R.Table("tb_work").Where("deleted_at is NOT NULL").
		Offset(page.Offset).
		Limit(page.Size).
		Find(&class).Error
	err = dao.R.Table("tb_work").Where("deleted_at is NOT NULL").Count(&total).Error


	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return class, total, nil
}

func (dao *DAO) GetWork(w Work) (*Work, error)  {
	work := Work{}
	tx := dao.R.Table("tb_class")
	if w.CourseID != 0 {
		tx.Where("courer_id = ?", w.CourseID)
	}

	if w.WorkName != "" {
		tx.Where("work_name = ?", w.WorkName)
	}

	if w.ClassID != 0 {
		tx.Where("class_id" , w.ClassID)
	}


	err := tx.Find(&work).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &work, nil
}