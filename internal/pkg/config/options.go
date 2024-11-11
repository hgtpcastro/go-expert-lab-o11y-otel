package config

type httpOptions struct {
	Name string
	Host string
	Port string
}

type weatherApiOptions struct {
	Url string
}

type zipCodeExternalApiOptions struct {
	Url string
}

type weatherExternalApiOptions struct {
	Key string
	Url string
}

type Config struct {
	Http               httpOptions
	WeatherApi         weatherApiOptions
	ZipCodeExternalApi zipCodeExternalApiOptions
	WeatherExternalApi weatherExternalApiOptions
	Tracing            TracingOptions
}

type TracingOptions struct {
	ServiceName           string                 `mapstructure:"serviceName"`
	Version               string                 `mapstructure:"version"`
	InstrumentationName   string                 `mapstructure:"instrumentationName"`
	Id                    int64                  `mapstructure:"id"`
	AlwaysOnSampler       bool                   `mapstructure:"alwaysOnSampler"`
	ZipkinExporterOptions *ZipkinExporterOptions `mapstructure:"zipkinExporterOptions"`
}

type ZipkinExporterOptions struct {
	Url string `mapstructure:"url"`
}
