package entity

import (
	"context"
	"github.com/Heqiaomu/hqmGframe/utils/gdb"
	"gorm.io/gorm"
)

type DAO struct {
	Table string
	R     *gorm.DB
	W     *gorm.DB
}

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

func New(ctx context.Context) *DAO {
	return &DAO{
		W: gdb.GormDB.GetWDB(ctx),
		R: gdb.GormDB.GetRDB(ctx),
	}
}
