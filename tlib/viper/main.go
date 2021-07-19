package main

import (
	"fmt"
	"github.com/spf13/viper"
)

//如何隔离 线上和线下的配置？
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
	v.SetConfigFile("tlib/viper/config/config-base.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("init config failed" + err.Error())
	}

	defer func() {
		fmt.Println("*******")
		fmt.Println(env)
		ViperInstance = v
	}()

	if env == "local" {
		// 读取local的配置
		configPath := "tlib/viper/config/config-local.yaml"
		v.SetConfigFile(configPath)
		err := v.ReadInConfig()
		if err != nil {
			panic("init config failed" + err.Error())
		}
		return
	}

	if env == "develop" {
		// 读取develop的配置
		configPath := "tlib/viper/config/config-develop.yaml"
		v.SetConfigFile(configPath)
		err := v.ReadInConfig()
		if err != nil {
			panic("init config failed" + err.Error())
		}
		return
	}

	if env == "production" {
		// 读取正式环境中的配置
		configPath := "tlib/viper/config/config-production.yaml"
		v.SetConfigFile(configPath)
		err := v.ReadInConfig()
		if err != nil {
			panic("init config failed" + err.Error())
		}
		return
	}
}

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MySqlConfig MySqlConfig `mapstructure:"mysql"` // 嵌套数据
}

type MySqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func main() {
	InitConfig()
	// 从配置文件中初始化 ServerConfig
	serverConfig := &ServerConfig{}
	if err := ViperInstance.Unmarshal(serverConfig); err != nil {
		panic(err)
	}
	fmt.Print(serverConfig)
}
