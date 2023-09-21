package rename

import (
	"fmt"
)

var (
	HealthErrorCannotRenameDevice = func(deviceName string) error {
		return fmt.Errorf("Failed to rename %s device", deviceName)
	}
)
