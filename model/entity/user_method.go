package entity

import (
	"github.com/Heqiaomu/hqmGframe/utils/gerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (dao *DAO) GetUserByPhone(phoneNum int64) (*User, error) {
	var userM = User{}
	userDB := dao.R.Table("tb_user").Where("phone = ?", phoneNum).Find(&userM)

	if userDB.Error != nil && !errors.Is(userDB.Error, gorm.ErrRecordNotFound) {
		return nil, userDB.Error
	}

	if userDB.RowsAffected == 0 {
		return nil, gerror.ErrorNotDFoundInDBByCondition
	}
	return &userM, nil
}

func (dao *DAO) GetUserByID(id int) (*User, error) {
	var userM = User{}
	userDB := dao.R.Table("tb_user").Where("id = ?", id).Find(&userM)

	if userDB.Error != nil && !errors.Is(userDB.Error, gorm.ErrRecordNotFound) {
		return nil, userDB.Error
	}

	if userDB.RowsAffected == 0 {
		return nil, gerror.ErrorNotDFoundInDBByCondition
	}
	return &userM, nil
}
