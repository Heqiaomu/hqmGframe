package router

import "github.com/gin-gonic/gin"

type classApiRouter struct {
}

func (far *classApiRouter) InitClassApi(root *gin.RouterGroup) {
	classGroup := root.Group("/api/v1/class")
	classGroup.POST("create", API.NewClass)

}

