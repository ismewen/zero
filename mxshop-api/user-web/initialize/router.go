package initialize

import (
	"github.com/gin-gonic/gin"
	userRouter "zero/mxshop-api/user-web/router"
)

func Routers() *gin.Engine {
	router := gin.Default()

	router.GET("health", func(ctx *gin.Context) {
		ctx.String(200, "healthy ok")
	})

	// 初始化user router
	apiGroup := router.Group("v1")
	userRouter.InitUserRouter(apiGroup)

	return router
}
