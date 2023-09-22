package entity

import "network-health/core/entity/device"

type ConnectionHandler interface {
	PingDevice(device *device.Device) device.Status
}
