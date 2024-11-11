package tracing

type TracingOptionsOld struct {
	ServiceName           string                    `mapstructure:"serviceName"`
	Version               string                    `mapstructure:"version"`
	InstrumentationName   string                    `mapstructure:"instrumentationName"`
	Id                    int64                     `mapstructure:"id"`
	AlwaysOnSampler       bool                      `mapstructure:"alwaysOnSampler"`
	ZipkinExporterOptions *ZipkinExporterOptionsOld `mapstructure:"zipkinExporterOptions"`
}

type ZipkinExporterOptionsOld struct {
	Url string `mapstructure:"url"`
}
