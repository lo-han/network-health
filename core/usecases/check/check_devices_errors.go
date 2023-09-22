package check

import (
	"fmt"
)

var (
	HealthErrorCannotConnectToServer = func(serverType string, serverError string) error {
		return fmt.Errorf("Failed to connect with %s address: %s", serverType, serverError)
	}
	HealthErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("Server error: %s", errorMsg)
	}
)
