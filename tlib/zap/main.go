package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func Sugar() {
	// 简单模式，易用性很高
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	url := "http://www.baidu.com"
	sugar.Infow("Sugar: Failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
}

func GtMode() {
	// 极致性能模式，输出日志的时候需要指定类型, 避免反射
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "http://www.baidu.com"
	logger.Warn("GtModel: Failed to fetch URL", zap.String("URL", url), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))

}

func FileModel() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./zap_file_mode.log",
		"stdout", // console 流
	}
	logger, _ := cfg.Build()
	url := "http://www.baidu.com"
	logger.Warn("FileModel: Failed to fetch URL", zap.String("URL", url), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))

}

func LoggerWithStacktrace(){
	// 打印堆栈信息
	logger, _ := zap.NewProduction(zap.AddStacktrace(zapcore.WarnLevel))
	url := "http://www.baidu.com"
	logger.Warn("WithStackTrace: Failed to fetch URL", zap.String("URL", url), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
}
func main() {
	Sugar()
	GtMode()
	FileModel()
	LoggerWithStacktrace()

	customLogger, _:= zap.NewProduction()
	zap.ReplaceGlobals(customLogger)

	zap.S() // 获取全局的Sugar
	zap.L() // 获取全局的logger, customLogger
}
