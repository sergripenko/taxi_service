package config

import (
	"github.com/spf13/viper"
)

// Init configuration file
func Init() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
