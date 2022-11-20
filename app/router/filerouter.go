package router

import "github.com/gin-gonic/gin"

type fileApiRouter struct {
}

func (far *fileApiRouter) InitFileApi(root *gin.RouterGroup) {
	fileGroup := root.Group("/api/v1/upload")
	fileGroup.POST("file", fileApi.UploadFile)

}
