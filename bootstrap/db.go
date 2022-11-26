package bootstrap

import (
	"context"
	"github.com/Heqiaomu/hqmGframe/model/entity"
	"github.com/Heqiaomu/hqmGframe/utils/gdb"
)

// 数据库初始化OK后，开始初始化表

func InitTable() error {
	return gdb.GormDB.GetWDB(context.Background()).
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&entity.User{},
			&entity.Work{},
			&entity.SubmitWork{},
			&entity.Class{},
		)
}
