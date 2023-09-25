package mocks

import (
	"context"
	"network-health/core/entity/logs"
	"time"
)

type MockLogger struct {
}

func (logger *MockLogger) Context(ctx context.Context) logs.Logger {
	return &MockLogger{}
}

func (logger *MockLogger) Error(message string) time.Time {
	return time.Now()
}

func (logger *MockLogger) Fatal(message string) time.Time {
	return time.Now()
}

func (logger *MockLogger) Info(message string) time.Time {
	return time.Now()
}
