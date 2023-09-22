package entity

import "network-health/core/entity/device"

type ConnectivityHandler interface {
	PingDevice(device *device.Device) device.Status
}
