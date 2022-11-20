package main

import (
	"github.com/Heqiaomu/hqmGframe/bootstrap"
)

func main() {
	bootstrap.InitConfig()

	// 开始启动服务
	bootstrap.Start()
}
