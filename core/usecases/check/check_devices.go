package check

import (
	"network-health/core/entity"
	"time"
)

type ConnectivityHandler interface {
	PingDevice(device *entity.Device) entity.Status
}

type Connectivity struct {
	handler ConnectivityHandler
}

func NewConnectivity(handler ConnectivityHandler) *Connectivity {
	return &Connectivity{
		handler: handler,
	}
}

func (conn *Connectivity) Check(store *entity.DeviceStore, handler ConnectivityHandler) (response *DeviceStatus) {
	var status entity.Status
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

func mapStatusToString(status entity.Status) string {
	var statString string

	switch status {
	case entity.Loaded:
		statString = "LOADED"
	case entity.Online:
		statString = "ONLINE"
	case entity.Offline:
		statString = "OFFLINE"
	}

	return statString
}
