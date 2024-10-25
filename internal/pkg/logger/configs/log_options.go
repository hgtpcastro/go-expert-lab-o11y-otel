package configs

import (
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs/environment"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/models"
	typemapper "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/reflection/type_mapper"
	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typemapper.GetGenericTypeNameByT[LogOptions]())

type LogOptions struct {
	LogLevel      string         `mapstructure:"level"`
	LogType       models.LogType `mapstructure:"logType"`
	CallerEnabled bool           `mapstructure:"callerEnabled"`
	EnableTracing bool           `mapstructure:"enableTracing" default:"true"`
}

func ProvideLogConfigs(env environment.Environment) (*LogOptions, error) {
	return configs.BindConfigKey[*LogOptions](optionName, env)
}
