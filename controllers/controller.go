package controllers

import (
	"fmt"
	connect "network-health/core/entity/connectivity"
	"network-health/core/entity/device_store"
	store "network-health/core/entity/device_store"
	"network-health/core/entity/logs"
	check_usecase "network-health/core/usecases/check"
	rename_usecase "network-health/core/usecases/rename"
	time_usecase "network-health/core/usecases/time"
)

type Controller struct {
	store *store.DeviceStore
	time  time_usecase.Time
}

func NewController(store *store.DeviceStore, time time_usecase.Time) *Controller {
	return &Controller{
		store: store,
		time:  time,
	}
}

func (controller *Controller) Check(handler connect.ConnectionHandler) (response *ControllerResponse, err error) {
	var status *check_usecase.DeviceStatus
	connection := check_usecase.NewConnectivity(handler, controller.time)

	logs.Gateway().Info("Checking devices...")
	status = connection.Check(controller.store)

	response = NewControllerResponse(NetStatOK, status)

	return
}

func (controller *Controller) Rename(oldName, newName string) (response *ControllerResponse, err error) {
	logs.Gateway().Info(fmt.Sprintf("Rename '%s' device to '%s'", oldName, newName))

	errStack := rename_usecase.Rename(controller.store, oldName, newName)

	if errStack.HasError() {
		switch errStack.EntityError() {

		case device_store.HealthErrorDeviceNotFound:
			err = errStack.UsecaseError()
			response = NewControllerError(NetStatNotFound, err.Error())
			break

		case device_store.HealthErrorDuplicatedName:
			err = errStack.UsecaseError()
			response = NewControllerError(NetStatBadRequest, err.Error())
		case device_store.HealthErrorInvalidName:
			err = errStack.UsecaseError()
			response = NewControllerError(NetStatBadRequest, err.Error())
			break
		}

		logs.Gateway().Error(fmt.Sprintf("%s", err.Error()))
		return
	}

	response = NewControllerEmptyResponse(NetStatNoContent)

	return
}
