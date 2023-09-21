package check

import (
	"fmt"
)

var (
	HealthErrorCannotConnectToServer = func(serverType string) error {
		return fmt.Errorf("Failed to connect with %s server", serverType)
	}
	HealthErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("Server error: %s", errorMsg)
	}
)
