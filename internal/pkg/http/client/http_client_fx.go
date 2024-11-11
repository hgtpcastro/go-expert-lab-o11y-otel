package client

import "go.uber.org/fx"

var Module = fx.Module(
	"httpclientfx",
	fx.Provide(NewHttpClient),
)
