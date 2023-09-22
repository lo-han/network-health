package controllers

import (
	entity "network-health/core/entity/device_list"
	check_usecase "network-health/core/usecases/check"
	rename_usecase "network-health/core/usecases/rename"

	"github.com/fatih/structs"
)

type Controller struct {
	store *entity.DeviceStore
}

func NewController(store *entity.DeviceStore) *Controller {
	return &Controller{
		store: store,
	}
}

func (controller *Controller) Check(handler check_usecase.ConnectivityHandler) (response *ControllerResponse, err error) {
	var status *check_usecase.DeviceStatus
	connection := check_usecase.NewConnectivity(handler)

	status = connection.Check(controller.store, handler)

	response = NewControllerResponse(NetStatOK, structs.Map(status))

	return
}

func (controller *Controller) Rename(oldName, newName string) (response *ControllerResponse, err error) {
	err = rename_usecase.Rename(controller.store, oldName, newName)

	if err != nil {
		response = NewControllerError(NetStatNotFound, err.Error())
		return
	}

	response = NewControllerResponse(NetStatNoContent, map[string]interface{}{})

	return
}
