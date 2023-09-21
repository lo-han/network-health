package controllers

import (
	"network-health/core/entity"
	"network-health/core/usecases/check"
	"network-health/core/usecases/rename"
	"time"
)

type Controller struct {
	store *entity.DeviceStore
}

func NewController(store *entity.DeviceStore) *Controller {
	return &Controller{
		store: store,
	}
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

func (controller *Controller) Check(handler check.ConnectivityHandler) (response *check.DeviceStatus, err error) {
	connection := check.NewConnectivity(handler)
	response = new(check.DeviceStatus)

	devices := controller.store.IterateDevices()

	err = connection.Check(devices)

	for _, device := range devices {
		response.Devices = append(response.Devices, check.Device{
			Name:    device.GetName(),
			Address: device.GetAddress(),
			Status:  mapStatusToString(device.GetStatus()),
		})

		response.Datetime = time.Now()
	}

	return
}

func (controller *Controller) Rename(oldName, newName string) (err error) {
	err = controller.store.RenameDevice(oldName, newName)

	if err != nil {
		err = rename.HealthErrorCannotRenameDevice(oldName)
		return
	}

	return
}
