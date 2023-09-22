package rename

import (
	"fmt"
)

var (
	HealthErrorCannotRenameDevice = func(deviceName string, reason string) error {
		return fmt.Errorf("Failed to rename '%s' device: %s", deviceName, reason)
	}
)

type RenameUsecaseErrorStack struct {
	errors []error
}

func (stack *RenameUsecaseErrorStack) HasError() bool {
	if len(stack.errors) > 0 {
		return true
	}

	return false
}

func (stack *RenameUsecaseErrorStack) Append(err error) {
	stack.errors = append(stack.errors, err)
}

func (stack *RenameUsecaseErrorStack) UsecaseError() (err error) {
	if stack.HasError() {
		return stack.errors[len(stack.errors)-1]
	}
	return
}

func (stack *RenameUsecaseErrorStack) EntityError() (err error) {
	if stack.HasError() {
		return stack.errors[len(stack.errors)-2]
	}
	return
}
