package device_list

import "errors"

var (
	HealthErrorFullDeviceList error = errors.New("List is full")
	HealthErrorDeviceNotFound error = errors.New("Device not found")
	// HealthErrorEmptyDeviceList error = errors.New("List is empty")
)
