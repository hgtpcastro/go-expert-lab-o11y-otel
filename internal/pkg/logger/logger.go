package logger

import "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/models"

type Fields map[string]interface{}

type Logger interface {
	Configure(cfg func(internalLog interface{}))
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, fields Fields)
	Error(args ...interface{})
	Errorw(msg string, fields Fields)
	Errorf(template string, args ...interface{})
	Err(msg string, err error)
	Fatal(args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, fields Fields)
	LogType() models.LogType
}
