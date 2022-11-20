package clientip

import "github.com/gin-gonic/gin"

func ClientIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIP := ctx.ClientIP()
		ctx.Set("ClientIP", clientIP)
		ctx.Next()
	}
}
