package device_store

import "errors"

var (
	HealthErrorFullDeviceList error = errors.New("List is full")
	HealthErrorDeviceNotFound error = errors.New("Device not found")
	HealthErrorDuplicatedName error = errors.New("Duplicated device name")
	HealthErrorInvalidName    error = errors.New("Invalid device name")
	// HealthErrorEmptyDeviceList error = errors.New("List is empty")
)
