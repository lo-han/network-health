package check

import (
	device "network-health/core/entity/device"
	store "network-health/core/entity/device_store"
	"time"
)

type ConnectivityHandler interface {
	PingDevice(device *device.Device) device.Status
}

type Connectivity struct {
	handler ConnectivityHandler
}

func NewConnectivity(handler ConnectivityHandler) *Connectivity {
	return &Connectivity{
		handler: handler,
	}
}

func (conn *Connectivity) Check(store *store.DeviceStore, handler ConnectivityHandler) (response *DeviceStatus) {
	var status device.Status
	response = new(DeviceStatus)

	devices := store.IterateDevices()

	for _, device := range devices.List() {
		status = conn.handler.PingDevice(device)

		device.SetStatus(status)

		response.Devices = append(response.Devices, Device{
			Name:    device.Name(),
			Address: device.Address(),
			Status:  mapStatusToString(device.Status()),
		})

		response.Datetime = time.Now()
	}

	return
}

func mapStatusToString(status device.Status) string {
	var statString string

	switch status {
	case device.Loaded:
		statString = "LOADED"
	case device.Online:
		statString = "ONLINE"
	case device.Offline:
		statString = "OFFLINE"
	}

	return statString
}
