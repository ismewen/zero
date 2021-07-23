package config


type UserSrvConfig struct {
	Name string `mastructure:"name"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv"`
	ConsulConfig `mapstructure:"consul"`
}

