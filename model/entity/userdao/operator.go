package userdao

import (
	"context"
	"github.com/Heqiaomu/hqmGframe/utils/gdb"
	"github.com/Heqiaomu/hqmGframe/utils/gerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Method interface {
	// GetUserByUserName 根据用户名获取用户信息
	GetUserByUserName(phoneNum string) (*User, error)
	// GetUserByPhone 通过手机号查询用户
	GetUserByPhone(phoneNum int64) (*User, error)
	// GetUserByEmail 根据Email查询用户
	GetUserByEmail(email string) (*User, error)
	// UpdateUser 更新用户
	UpdateUser(model *User) (*User, error)
	// DeleteUser 删除用户
	DeleteUser(userID string) (*User, error)
}

type userDb struct {
	Table string
	r     *gorm.DB
	W     *gorm.DB
}

var db *userDb

func NewUserDAO(ctx context.Context) *userDb {
	if db == nil {
		db = &userDb{
			Table: "tb_user",
			W:     gdb.GormDB.GetWDB(ctx),
			r:     gdb.GormDB.GetRDB(ctx),
		}
	}
	return db
}

func (db *userDb) GetUserByPhone(phoneNum int64) (*User, error) {
	var userM = User{}
	userDB := db.r.Table("tb_user").Where("phone = ?", phoneNum).Find(&userM)

	if userDB.Error != nil && !errors.Is(userDB.Error, gorm.ErrRecordNotFound) {
		return nil, userDB.Error
	}

	if userDB.RowsAffected == 0 {
		return nil, gerror.ErrorNotDFoundInDBByCondition
	}
	return &userM, nil
}
