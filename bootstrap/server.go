package bootstrap

import (
	"github.com/Heqiaomu/hqmGframe/app/router"
	"github.com/Heqiaomu/hqmGframe/model/config"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func Start() {
	routers := router.Routers()
	httpServer := config.GetHttpServer()
	port := httpServer.API.Port
	s := initServer(port, routers)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	tls := httpServer.TLS
	if tls.Enabled {
		s.ListenAndServeTLS(tls.Cert, tls.Key)
	} else {
		s.ListenAndServe()
	}
}

type server interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
