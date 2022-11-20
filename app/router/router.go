package router

import (
	"github.com/Heqiaomu/hqmGframe/app/middleware/clientip"
	"github.com/Heqiaomu/hqmGframe/app/middleware/cors"
	"github.com/Heqiaomu/hqmGframe/model/config"
	"github.com/Heqiaomu/hqmGframe/utils/ggin"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var engine *gin.Engine
	if config.GetAppDebug() {
		engine = gin.Default()
		pprof.Register(engine)
	} else {
		engine = ggin.New()
	}

	engine.Use(cors.Next(), clientip.ClientIP()) // 直接放行全部跨域请求

	// 私有化的路由请求
	private := engine.Group("")
	{
		fileUpload.InitFileApi(private)
		userRouter.InitUserRouter(private)
	}

	return engine
}
