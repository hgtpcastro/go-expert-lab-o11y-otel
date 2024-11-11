package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger() *zap.Logger {
	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	var encoder zapcore.Encoder

	encoderCfg = zap.NewDevelopmentEncoderConfig()
	encoderCfg.NameKey = "[SERVICE]"
	encoderCfg.TimeKey = "[TIME]"
	encoderCfg.LevelKey = "[LEVEL]"
	encoderCfg.FunctionKey = "[CALLER]"
	encoderCfg.CallerKey = "[LINE]"
	encoderCfg.MessageKey = "[MESSAGE]"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeName = zapcore.FullNameEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
	encoderCfg.ConsoleSeparator = " | "
	encoder = zapcore.NewConsoleEncoder(encoderCfg)

	logLevel := zapcore.DebugLevel

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))

	var options []zap.Option

	// options = append(options, zap.AddCallerSkip(1))

	logger := zap.New(core, options...)

	//if l.logOptions.EnableTracing {
	// add logs as events to tracing
	// logger = otelzap.New(logger).Logger
	//}

	return logger
}
