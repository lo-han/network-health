package rename

import (
	"fmt"
)

var (
	HealthErrorCannotRenameDevice = func(deviceName string, reason string) error {
		return fmt.Errorf("Failed to rename '%s' device: %s", deviceName, reason)
	}
)
