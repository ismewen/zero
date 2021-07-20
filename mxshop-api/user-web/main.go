package main

import (
	"go.uber.org/zap"
	"zero/mxshop-api/user-web/initialize"
)

func main() {
	// init router
	router := initialize.Routers()
	initialize.InitLogger()

	// init config

	initialize.InitConfig()

	addr := ":8001"

	sugar := zap.S() // 一个安全的logger
	sugar.Infof("启动服务器, 地址:%s", addr)

	// run server
	err := router.Run(addr)
	if err != nil {
		sugar.Panic("启动失败", err.Error())
	}
}
