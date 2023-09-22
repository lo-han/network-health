package device

import "fmt"

var (
	HealthErrorInvalidAddress = func(addressType string) error {
		return fmt.Errorf("Invalid %s address", addressType)
	}
)
