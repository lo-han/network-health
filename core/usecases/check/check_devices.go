package check

import "network-health/core/entity"

type ConnectivityHandler interface {
	PingDevice(device *entity.Device) (entity.Status, error)
}

type Connectivity struct {
	handler ConnectivityHandler
}

func NewConnectivity(handler ConnectivityHandler) *Connectivity {
	return &Connectivity{
		handler: handler,
	}
}

func (conn *Connectivity) Check(devices entity.Iteration) (err error) {
	var status entity.Status

	for _, device := range devices {
		status, err = conn.handler.PingDevice(device)
		if err != nil {
			break
		}

		device.SetStatus(status)
	}

	return
}
