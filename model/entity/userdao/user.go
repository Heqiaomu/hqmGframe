package userdao

import (
	"github.com/Heqiaomu/hqmGframe/model/entity"
	"time"
)

type User struct {
	entity.Model            // 常规
	UserID        string    `json:"user_id" gorm:"user_id,index:idx_user"`   // 用户的 ID
	Username      string    `json:"username" gorm:"username,index:idx_user"` // 登录用户名，一般为手机号
	Password      string    `json:"password" gorm:"password"`                // 登录密码
	UserType      string    `json:"user_type" gorm:"user_type"`              // 老师 or 学生
	Email         string    `json:"email" gorm:"email"`                      // 邮箱
	Phone         int64     `json:"phone" gorm:"phone"`                      // 手机号
	Sex           int64     `json:"sex" gorm:"sex"`                          // 性别
	Married       int64     `json:"married" gorm:"married"`                  // 是否结婚
	Provice       string    `json:"provice" gorm:"provice"`
	IsOver        int64     `json:"is_over" gorm:"is_over"`       // 是否毕业 / 是否离职
	StartTime     time.Time `json:"start_time" gorm:"start_time"` // 入学时间 / 入职
	EndTime       time.Time `json:"end_time" gorm:"end_time"`     // 毕业时间 / 离职时间
	LastLoginTime time.Time `json:"last_login_time" gorm:"last_login_time"`
}

func (u User) TableName() string {
	return "tb_user"
}
