package gdb

import (
	"context"
	"github.com/Heqiaomu/hqmGframe/model/config"
	"gorm.io/gorm"
)

var w *gorm.DB
var r *gorm.DB

type DB interface {
	GetWDB(ctx context.Context) *gorm.DB
	GetRDB(ctx context.Context) *gorm.DB
}

var DBInter DB

var GormDB = new(DataBase)

type DataBase struct{}

func InitDB() error {
	var err error
	w, r, err = getSqlDriver(config.GetDB().UseDbType)
	if err != nil {
		return err
	}
	return nil
}

func (m *DataBase) GetWDB(ctx context.Context) *gorm.DB {
	return w.WithContext(ctx)
}
func (m *DataBase) GetRDB(ctx context.Context) *gorm.DB {
	return r.WithContext(ctx)
}
