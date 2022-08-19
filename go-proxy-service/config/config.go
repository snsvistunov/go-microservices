package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"PORT"`
	AuthServiceName string `mapstructure:"AUTH_SVC_URL"`
	UserServiceName string `mapstructure:"USER_SVC_URL"`
}

var Cfg *Config

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	Cfg = &config
	fmt.Printf("Config loaded: %v/n", Cfg)
	return
}
