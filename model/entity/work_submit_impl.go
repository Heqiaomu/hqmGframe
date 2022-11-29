package entity

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//SubmitWork

func (dao *DAO) Submit(sub SubmitWork) error {
	return dao.W.Table("tb_submit_work").Create(&sub).Error
}

func (dao *DAO) GetSubmit(sub SubmitWork)(*SubmitWork,error)  {
	work := SubmitWork{}
	tx := dao.R.Table("tb_submit_work")
	if sub.WorkID != 0 {
		tx.Where("work_id = ?", sub.WorkID)
	}

	if sub.Submitter != 0 {
		tx.Where("submitter = ?", sub.Submitter)
	}

	if sub.ClassID != 0 {
		tx.Where("class_id" , sub.ClassID)
	}
	err := tx.Find(&work).Error
	return &work, err
}

func (dao *DAO) GetSubmitList(page Page) ([]*SubmitWork, int64, error)  {
	var class []*SubmitWork
	var total int64
	err := dao.R.Table("tb_submit_work").
		Offset(page.Offset).
		Limit(page.Size).
		Find(&class).Error
	err = dao.R.Table("tb_submit_work").Count(&total).Error


	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return class, total, nil
}