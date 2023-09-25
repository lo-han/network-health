package stdout

import (
	"context"
	"fmt"
	"network-health/core/entity/logs"
	"time"
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

func (logger *STDOutLogger) Error(message string) time.Time {
	now := time.Now()
	fmt.Printf("[%s] %s [%s]\n", logger.errorPrefix, message, now)
	return now
}

func (logger *STDOutLogger) Fatal(message string) time.Time {
	now := time.Now()
	fmt.Printf("[%s] %s [%s]\n", logger.fatalPrefix, message, now)
	return now
}

func (logger *STDOutLogger) Info(message string) time.Time {
	now := time.Now()
	fmt.Printf("[%s] %s [%s]\n", logger.infoPrefix, message, now)
	return now
}
