package logs

import "context"

type Logger interface {
	Context(ctx context.Context) Logger
	Error(message string)
	Fatal(message string)
	Info(message string)
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
