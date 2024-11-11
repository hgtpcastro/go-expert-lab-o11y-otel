package log

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"zapfx",
	fx.Provide(NewZapLogger),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
