package config

import (
	"github.com/Heqiaomu/hqmGframe/utils/gviper"
)

var cfg App

func New() error {
	if err := gviper.GetV().Unmarshal(&cfg); err != nil {
		return err
	}
	return nil
}

func GetConfig() *App {
	return &cfg
}

func GetHttpServer() *HTTPServer {
	return &cfg.HTTPServer
}
