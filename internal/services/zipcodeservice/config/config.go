package config

import (
	"strings"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs/environment"
)

type Config struct {
	AppOptions AppOptions `mapstructure:"appOptions"`
}

func NewConfig(environment environment.Environment) (*Config, error) {
	cfg, err := configs.BindConfig[*Config](environment)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type AppOptions struct {
	DeliveryType string `mapstructure:"deliveryType"`
	ServiceName  string `mapstructure:"serviceName"`
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}
