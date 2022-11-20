package gdb

import (
	"github.com/Heqiaomu/hqmGframe/model/config"
	"github.com/Heqiaomu/hqmGframe/utils/gerror"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"time"
)

func getSqlDriver(sqlT string) (*gorm.DB, *gorm.DB, error) {
	wDialector, rDialector, err := getDialector(sqlT)
	if err != nil {
		return nil, nil, err
	}
	//cu := sqlLogCu(sqlT)
	w, err := gorm.Open(wDialector,
		GetConfig(),
	)
	if err != nil {
		return nil, nil, err
	}

	r, err := gorm.Open(rDialector,
		GetConfig(),
	)
	if err != nil {
		return nil, nil, err
	}

	// 针对读数据库
	resolverConf := dbresolver.Config{
		Replicas: []gorm.Dialector{rDialector}, //  读 操作库，查询类
		Policy:   dbresolver.RandomPolicy{},    // sources/replicas 负载均衡策略适用于
	}
	err = r.Use(dbresolver.Register(resolverConf).SetConnMaxIdleTime(time.Second * 30).
		SetConnMaxLifetime(config.GetDB().GetReadConnMaxLifetime() * time.Second).
		SetMaxIdleConns(config.GetDB().GetReadMaxIdleConn()).
		SetMaxOpenConns(config.GetDB().GetReadMaxOpenConn()))
	if err != nil {
		return nil, nil, err
	}

	// 查询没有数据，屏蔽 gorm v2 包中会爆出的错误
	// https://github.com/go-gorm/gorm/issues/3789  此 issue 所反映的问题就是我们本次解决掉的
	_ = r.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", MaskNotDataError)

	// https://github.com/go-gorm/gorm/issues/4838
	_ = w.Callback().Create().Before("gorm:before_create").Register("CreateBeforeHook", CreateBeforeHook)
	// 为了完美支持gorm的一系列回调函数
	_ = w.Callback().Update().Before("gorm:before_update").Register("UpdateBeforeHook", UpdateBeforeHook)

	wDB, err := w.DB()
	if err != nil {
		return nil, nil, err
	}
	wDB.SetConnMaxIdleTime(time.Second * 30)
	wDB.SetConnMaxLifetime(config.GetDB().GetWriterConnMaxLifetime() * time.Second)
	wDB.SetMaxIdleConns(config.GetDB().GetWriteMaxIdleConn())
	wDB.SetMaxOpenConns(config.GetDB().GetWriteMaxOpenConn())
	return w, r, nil
}

func getDialector(sqlT string) (gorm.Dialector, gorm.Dialector, error) {
	switch sqlType(sqlT) {
	case MySql:
		return mysql.Open(config.GetDB().Mysql.GetWriteDSN()), mysql.Open(config.GetDB().Mysql.GetReadDSN()), nil
	case SqlServer:
		return sqlserver.Open(config.GetDB().SQLServer.GetWriteDSN()), sqlserver.Open(config.GetDB().SQLServer.GetReadDSN()), nil
	case PostgreSQL:
		return postgres.Open(config.GetDB().PostgreSQL.GetWriteDSN()), postgres.Open(config.GetDB().PostgreSQL.GetReadDSN()), nil
	default:
		return nil, nil, errors.New(gerror.ErrorsDbDriverNotExists + sqlT)
	}
}

type sqlType string

var (
	MySql      sqlType = "mysql"
	SqlServer  sqlType = "sqlserver"
	PostgreSQL sqlType = "postgresql"
)

func sqlLogCu(sqlT string) gormLog.Interface {
	return createCustomGormLog(sqlT,
		SetInfoStrFormat("[info] %s\n"),
		SetWarnStrFormat("[warn] %s\n"),
		SetErrStrFormat("[error] %s\n"),
		SetTraceStrFormat("[traceStr] %s [%.3fms] [rows:%v] %s\n"),
		SetTracWarnStrFormat("[traceWarn] %s %s [%.3fms] [rows:%v] %s\n"),
		SetTracErrStrFormat("[traceErr] %s %s [%.3fms] [rows:%v] %s\n"))

}
