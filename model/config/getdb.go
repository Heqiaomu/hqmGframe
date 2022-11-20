package config

import (
	"fmt"
	"time"
)

func GetDB() *Database {
	return &cfg.Database
}

func (mysql Mysql) GetWriteDSN() string {
	write := mysql.Write
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", write.User,
		write.Pass, write.Host, write.Port, write.DataBase, write.Charset)
}

func (mysql Mysql) GetReadDSN() string {
	read := mysql.Read
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", read.User,
		read.Pass, read.Host, read.Port, read.DataBase, read.Charset)
}

func (sqlServer SQLServer) GetWriteDSN() string {
	write := sqlServer.Write
	return fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable",
		write.Host,
		write.Port,
		write.DataBase,
		write.User,
		write.Pass,
	)
}
func (sqlServer SQLServer) GetReadDSN() string {
	read := sqlServer.Read
	return fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable",
		read.Host,
		read.Port,
		read.DataBase,
		read.User,
		read.Pass,
	)
}

func (postgreSQL PostgreSQL) GetWriteDSN() string {
	write := postgreSQL.Write
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		write.Host,
		write.Port,
		write.DataBase,
		write.User,
		write.Pass,
	)
}
func (postgreSQL PostgreSQL) GetReadDSN() string {
	read := postgreSQL.Read
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		read.Host,
		read.Port,
		read.DataBase,
		read.User,
		read.Pass,
	)
}

type DSN interface {
	GetWriteDSN() string
	GetReadDSN() string
}

type Duration interface {
	GetWriterConnMaxLifetime() time.Duration
	GetReadConnMaxLifetime() time.Duration
	GetReadMaxIdleConn() int
	GetWritePingFailRetryTime() int
	GetWriteMaxIdleConn() int
	GetWriteMaxOpenConn() int
	GetReadMaxOpenConn() int
	GetWriteReConnectInterval() int
}

func (da Database) GetWriterConnMaxLifetime() time.Duration {
	switch da.UseDbType {
	case "mysql":
		return time.Duration(da.Mysql.Write.SetConnMaxLifetime) * time.Second
	case "sqlserver":
		return time.Duration(da.SQLServer.Write.SetConnMaxLifetime) * time.Second
	case "postgresql":
		return time.Duration(da.PostgreSQL.Write.SetConnMaxLifetime) * time.Second
	default:
		return 30 * time.Second
	}
}

func (da Database) GetReadConnMaxLifetime() time.Duration {
	switch da.UseDbType {
	case "mysql":
		return time.Duration(da.Mysql.Read.SetConnMaxLifetime) * time.Second
	case "sqlserver":
		return time.Duration(da.SQLServer.Read.SetConnMaxLifetime) * time.Second
	case "postgresql":
		return time.Duration(da.PostgreSQL.Read.SetConnMaxLifetime) * time.Second
	default:
		return 30 * time.Second
	}
}

func (da Database) GetReadMaxIdleConn() int {
	switch da.UseDbType {
	case "mysql":
		return da.Mysql.Read.SetMaxIdleConns
	case "sqlserver":
		return da.SQLServer.Read.SetMaxIdleConns
	case "postgresql":
		return da.PostgreSQL.Read.SetMaxIdleConns
	default:
		return 100
	}
}

func (da Database) GetWriteMaxIdleConn() int {
	switch da.UseDbType {
	case "mysql":
		return da.Mysql.Write.SetMaxIdleConns
	case "sqlserver":
		return da.SQLServer.Write.SetMaxIdleConns
	case "postgresql":
		return da.PostgreSQL.Write.SetMaxIdleConns
	default:
		return 100
	}
}

func (da Database) GetWriteMaxOpenConn() int {
	switch da.UseDbType {
	case "mysql":
		return da.Mysql.Write.SetMaxOpenConns
	case "sqlserver":
		return da.SQLServer.Write.SetMaxOpenConns
	case "postgresql":
		return da.PostgreSQL.Write.SetMaxOpenConns
	default:
		return 100
	}
}

func (da Database) GetReadMaxOpenConn() int {
	switch da.UseDbType {
	case "mysql":
		return da.Mysql.Read.SetMaxOpenConns
	case "sqlserver":
		return da.SQLServer.Read.SetMaxOpenConns
	case "postgresql":
		return da.PostgreSQL.Read.SetMaxOpenConns
	default:
		return 100
	}
}

func (da Database) GetWriteReConnectInterval() int {
	switch da.UseDbType {
	case "mysql":
		return da.Mysql.Write.ReConnectInterval
	case "sqlserver":
		return da.SQLServer.Write.ReConnectInterval
	case "postgresql":
		return da.PostgreSQL.Write.ReConnectInterval
	default:
		return 100
	}
}

func (da Database) GetWritePingFailRetryTime() int {
	switch da.UseDbType {
	case "mysql":
		return da.Mysql.Write.PingFailRetryTimes
	case "sqlserver":
		return da.SQLServer.Write.PingFailRetryTimes
	case "postgresql":
		return da.PostgreSQL.Write.PingFailRetryTimes
	default:
		return 100
	}
}
