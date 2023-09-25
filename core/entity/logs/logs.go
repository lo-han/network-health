package logs

import (
	"context"
	"time"
)

type Logger interface {
	Context(ctx context.Context) Logger
	Error(message string) time.Time
	Fatal(message string) time.Time
	Info(message string) time.Time
}

var loggerGateway Logger = nil

func SetLogger(logger Logger) (err error) {
	if loggerGateway == nil {
		loggerGateway = logger
	}

	return HealthErrorLoggerAlreadyInitialized
}

func Gateway() Logger {
	return loggerGateway
}
