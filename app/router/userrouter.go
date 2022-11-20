package router

import "github.com/gin-gonic/gin"

type User struct {
}

var userRouter = new(User)

func (u *User) InitUserRouter(root *gin.RouterGroup) {
	userGroup := root.Group("/api/v1/user")
	userGroup.POST("login", userApi.Login)
	userGroup.POST("register", userApi.Register)
}
