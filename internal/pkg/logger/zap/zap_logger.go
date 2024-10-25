package zap

import (
	"os"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs/environment"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger"
	config2 "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/configs"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	level       string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
	logOptions  *config2.LogOptions
}

type ZapLogger interface {
	logger.Logger
	InternalLogger() *zap.Logger
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Sync() error
}

// For mapping config logger
var loggerLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

// NewZapLogger create new zap logger
func NewZapLogger(
	cfg *config2.LogOptions,
	env environment.Environment,
) ZapLogger {
	zapLogger := &zapLogger{level: cfg.LogLevel, logOptions: cfg}
	zapLogger.initLogger(env)

	return zapLogger
}

func (l *zapLogger) InternalLogger() *zap.Logger {
	return l.logger
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the logger then panics. (See DPanicLevel for details.)
func (l *zapLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the logger then panics. (See DPanicLevel for details.)
func (l *zapLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

// Sync flushes any buffered log entries
func (l *zapLogger) Sync() error {
	go func() {
		err := l.logger.Sync()
		if err != nil {
			l.logger.Error("error while syncing", zap.Error(err))
		}
	}() // nolint: errcheck
	return l.sugarLogger.Sync()
}

func (l *zapLogger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// InitLogger Init logger
func (l *zapLogger) initLogger(env environment.Environment) {
	logLevel := l.getLoggerLevel()

	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	var encoder zapcore.Encoder

	if env.IsProduction() {
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.NameKey = "[SERVICE]"
		encoderCfg.TimeKey = "[TIME]"
		encoderCfg.LevelKey = "[LEVEL]"
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.CallerKey = "[LINE]"
		encoderCfg.MessageKey = "[MESSAGE]"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
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
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))

	var options []zap.Option

	if l.logOptions.CallerEnabled {
		options = append(options, zap.AddCaller())
		options = append(options, zap.AddCallerSkip(1))
	}

	logger := zap.New(core, options...)

	// if l.logOptions.EnableTracing {
	// 	// add logs as events to tracing
	// 	logger = otelzap.New(logger).Logger
	// }

	l.logger = logger
	l.sugarLogger = logger.Sugar()
}

func (l *zapLogger) Configure(cfg func(internalLog interface{})) {
	cfg(l.logger)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *zapLogger) Debugw(msg string, fields logger.Fields) {
	zapFields := mapToZapFields(fields)
	l.logger.Debug(msg, zapFields...)
}

func (l *zapLogger) LogType() models.LogType {
	return models.Zap
}

// Error uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

// Errorw logs a message with some additional context.
func (l *zapLogger) Errorw(msg string, fields logger.Fields) {
	zapFields := mapToZapFields(fields)
	l.logger.Error(msg, zapFields...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

// Err uses error to log a message.
func (l *zapLogger) Err(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

// Info uses fmt.Sprint to construct and log a message
func (l *zapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Infow logs a message with some additional context.
func (l *zapLogger) Infow(msg string, fields logger.Fields) {
	zapFields := mapToZapFields(fields)
	l.logger.Info(msg, zapFields...)
}

func mapToZapFields(data map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(data))

	for key, value := range data {
		field := zap.Field{
			Key:       key,
			Type:      getFieldType(value),
			Interface: value,
		}
		fields = append(fields, field)
	}

	return fields
}

func getFieldType(value interface{}) zapcore.FieldType {
	switch value.(type) {
	case string:
		return zapcore.StringType
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return zapcore.Int64Type
	case bool:
		return zapcore.BoolType
	case float32, float64:
		return zapcore.Float64Type
	case error:
		return zapcore.ErrorType
	default:
		// uses reflection to serialize arbitrary objects, so it can be slow and allocation-heavy.
		return zapcore.ReflectType
	}
}
