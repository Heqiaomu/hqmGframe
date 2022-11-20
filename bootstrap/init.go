package bootstrap

import (
	"github.com/Heqiaomu/hqmGframe/model/config"
	"github.com/Heqiaomu/hqmGframe/utils/gdb"
	"github.com/Heqiaomu/hqmGframe/utils/gos"
	"github.com/Heqiaomu/hqmGframe/utils/gsnow"
	"github.com/Heqiaomu/hqmGframe/utils/gviper"
	"log"
	"path/filepath"
)

func InitConfig() {
	// 开始初始化服务
	// 先检查当前配置文件是否存在
	homeDir := gos.HomeDir()
	cfgPath := filepath.Join(homeDir, "/cmd/config.yaml")
	_, err := gos.FileIsExist(cfgPath)
	if err != nil {
		log.Fatalln(err)
	}
	// 文件存在，开始执行 viper
	if err = gviper.New(cfgPath); err != nil {
		log.Fatalln(err)
	}
	// 开始初始化 config
	if err := config.New(); err != nil {
		log.Fatalln(err)
	}
	// 开始初始化 雪花算法
	gsnow.New()

	// 数据库初始化
	if err = gdb.InitDB(); err != nil {
		log.Fatalln(err)
	}

	if err = InitTable(); err != nil {
		log.Fatalln(err)
	}
}
