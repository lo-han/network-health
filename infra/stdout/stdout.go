package stdout

import (
	"context"
	"fmt"
	"network-health/core/entity/logs"
)

type STDOutLogger struct {
	infoPrefix  string
	fatalPrefix string
	errorPrefix string
}

func NewSTDOutLogger() *STDOutLogger {
	return &STDOutLogger{
		infoPrefix:  "INFO",
		errorPrefix: "ERROR",
		fatalPrefix: "FATAL",
	}
}

func (logger *STDOutLogger) Context(ctx context.Context) logs.Logger {
	return &STDOutLogger{}
}

func (logger *STDOutLogger) Error(message string) {
	fmt.Printf("[%s] %s", logger.errorPrefix, message)
}

func (logger *STDOutLogger) Fatal(message string) {
	fmt.Printf("[%s] %s", logger.errorPrefix, message)
}

func (logger *STDOutLogger) Info(message string) {
	fmt.Printf("[%s] %s", logger.errorPrefix, message)
}
