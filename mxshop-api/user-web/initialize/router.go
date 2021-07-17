package initialize

import (
	"github.com/gin-gonic/gin"
	userRouter "zero/mxshop-api/user-web/router"
)

func Routers() *gin.Engine{
	router := gin.Default()

	// 初始化user router
	apiGroup := router.Group("v1")
	userRouter.InitUserRouter(apiGroup)

	return router
}