package bootstrap

import (
	"github.com/heqiaomu/webframe/model/config"
	"github.com/heqiaomu/webframe/utils/gos"
	"github.com/heqiaomu/webframe/utils/gviper"
	"log"
	"path/filepath"
)

func InitServer() {
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
}
