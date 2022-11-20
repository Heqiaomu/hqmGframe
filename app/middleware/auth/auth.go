package auth

import (
	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}

func CheckTokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
