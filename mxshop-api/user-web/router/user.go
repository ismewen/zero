package router

import (
	"github.com/gin-gonic/gin"
	"zero/mxshop-api/user-web/api"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userRouter.GET("list", api.GetUserList)
	userRouter.POST("login", api.Login)
}