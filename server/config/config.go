package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port int
}

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("server-settings")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault("Port", 3000)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &AppConfig{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
