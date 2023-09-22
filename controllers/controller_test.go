package controllers

import (
	device "network-health/core/entity/device"
	"network-health/core/entity/device_store"
	"network-health/core/usecases/check"
	"network-health/core/usecases/rename"
	"reflect"
	"testing"

	"github.com/fatih/structs"
)

type mockAddress struct{}

func (ip *mockAddress) Set(address string) (err error) {
	return
}

func (ip *mockAddress) Get() (address string) {
	return "address"
}

type connHandlerMock struct{}

func (connHandlerMock) PingDevice(dev *device.Device) (deviceStatus device.Status) {
	return device.Online
}

func Test_Controller_Check(t *testing.T) {
	deviceStoreTest, _ := device_store.NewDeviceStore(device.NewDevice(&mockAddress{}, "device"))
	content := structs.Map(check.DeviceStatus{
		Devices: []check.Device{
			{
				Name:    "device",
				Address: "address",
				Status:  "ONLINE",
			},
		},
	})

	testCases := []struct {
		name       string
		controller *Controller
		response   *ControllerResponse
	}{
		{
			name:       "Succesfull check",
			controller: NewController(deviceStoreTest),
			response: NewControllerResponse(
				NetStatOK,
				content,
			),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			response, _ := testCase.controller.Check(connHandlerMock{})

			if response.Code != testCase.response.Code {
				t.Errorf("Test_Controller_Check().Code = %s, expected %s", response.Code, testCase.response.Code)
			}

			if !reflect.DeepEqual(response.Content, testCase.response.Content) {
				t.Error("Test_Controller_Check().Content different from expected!")
			}
		})
	}
}
func Test_Controller_Rename(t *testing.T) {
	failError := rename.HealthErrorCannotRenameDevice("doesnt_exists", device_store.HealthErrorDeviceNotFound.Error())

	testCases := []struct {
		name     string
		oldName  string
		newName  string
		response *ControllerResponse
		err      error
	}{
		{
			name:    "Succesfull rename",
			oldName: "device",
			newName: "new_device",
			response: NewControllerResponse(
				NetStatNoContent,
				map[string]interface{}{},
			),
			err: nil,
		},
		{
			name:    "Failed to rename unexistent device",
			oldName: "doesnt_exists",
			newName: "new_device",
			response: NewControllerError(
				NetStatNotFound,
				failError.Error(),
			),
			err: failError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			deviceStoreTest, _ := device_store.NewDeviceStore(device.NewDevice(&mockAddress{}, "device"))
			controller := NewController(deviceStoreTest)

			response, err := controller.Rename(testCase.oldName, testCase.newName)

			if err != nil {
				if err.Error() != testCase.err.Error() {
					t.Errorf("Test_Controller_Rename().Err = %s, expected %s", err.Error(), testCase.err.Error())
				}
			}

			if !reflect.DeepEqual(response.Content, testCase.response.Content) {
				t.Errorf("Test_Controller_Rename().Response.Content = %s, expected %s", response.Content, testCase.response.Content)
			}

			if response.Code != response.Code {
				t.Errorf("Test_Controller_Rename().Response.Code = %s, expected %s", response.Code, testCase.response.Code)
			}
		})
	}
}
