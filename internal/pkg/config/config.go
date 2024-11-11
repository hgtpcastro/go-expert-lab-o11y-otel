package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func getViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config.yaml")
	v.SetConfigType("yaml")
	//v.AddConfigPath("./internal/cmd/app")
	v.AddConfigPath("./cmd/app")
	return v
}

func NewConfig(log *zap.Logger) (*Config, error) {
	log.Info("Loading configuration")
	v := getViper()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = v.Unmarshal(&cfg)
	return &cfg, err
}
