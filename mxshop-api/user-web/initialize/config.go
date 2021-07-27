package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
	"zero/mxshop-api/user-web/global"
)

var ViperInstance *viper.Viper

func InitConfig() {
	// 获取环境变量
	viper.AutomaticEnv()
	v := viper.New()
	env := viper.GetString("MXSHOP_ENV")

	if env == "" {
		env = "local"
	}

	// 读取默认的配置
	v.SetConfigFile("mxshop-api/user-web/config/config-base.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("init config failed" + err.Error())
	}

	defer func() {
		// 复制给全局变量
		ViperInstance = v
		v.WatchConfig()

		// 初始化 server config 信息
		err := v.Unmarshal(global.ServerConfig)

		// 定时打印serverconfig信息, 查看数据是否有变化
		go func() {
			for {
				time.Sleep(time.Second * 2)
				//zap.S().Infof("%+v", global.ServerConfig)
			}
		}()

		if err != nil {
			panic("init server config error" + err.Error())
		}
		v.OnConfigChange(func(in fsnotify.Event) {
			err = v.Unmarshal(global.ServerConfig)
			if err != nil {
				panic("init server config error")
			}
		})
	}()

	if env == "local" {
		// 读取local的配置
		configPath := "mxshop-api/user-web/config/configlocal.yaml"
		v.SetConfigFile(configPath)
		err := v.MergeInConfig()
		if err != nil {
			panic("init config failed" + err.Error())
		}
		return
	}

	if env == "develop" {
		// 读取develop的配置
		configPath := "mxshop-api/user-web/config/config-develop.yaml"
		v.SetConfigFile(configPath)
		err := v.MergeInConfig()
		if err != nil {
			panic("init config failed" + err.Error())
		}
		return
	}

	if env == "production" {
		// 读取正式环境中的配置
		configPath := "mxshop-api/user-web/config/config-production.yaml"
		v.SetConfigFile(configPath)
		err := v.MergeInConfig()
		if err != nil {
			panic("init config failed" + err.Error())
		}
		return
	}
}
