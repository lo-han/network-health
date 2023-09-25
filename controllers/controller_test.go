package controllers

import (
	"network-health/core/entity/connectivity"
	device "network-health/core/entity/device"
	"network-health/core/entity/device_store"
	"network-health/core/entity/logs"
	"network-health/core/usecases/check"
	"network-health/core/usecases/rename"
	"network-health/infra/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/fatih/structs"
)

type timeMock struct {
	now time.Time
}

func newTimeMock() *timeMock {
	return &timeMock{
		now: time.Now(),
	}
}

func (mock *timeMock) Now() time.Time {
	return mock.now
}

type mockAddress struct{}

func (ip *mockAddress) Set(address string) (err error) {
	return
}

func (ip *mockAddress) Get() (address string) {
	return "address"
}

type connHandlerMock struct{}

func (connHandlerMock) PingDevice(device *device.Device) (stats connectivity.ConnectionStats, err error) {
	stats.PacketsSent = 1
	stats.PacketsRecv = 1
	return
}

func Test_Controller_Check(t *testing.T) {
	logs.SetLogger(&mocks.MockLogger{})
	timeNow := newTimeMock()
	deviceStoreTest, _ := device_store.NewDeviceStore(device.NewDevice(&mockAddress{}, "device"))
	content := structs.Map(check.DeviceStatus{
		Devices: []check.Device{
			{
				Name:    "device",
				Address: "address",
				Status:  "ONLINE",
			},
		},
		Datetime: timeNow.Now(),
	})

	testCases := []struct {
		name       string
		controller *Controller
		response   *ControllerResponse
	}{
		{
			name:       "Succesfull check",
			controller: NewController(deviceStoreTest, timeNow),
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
	failErrorNotFound := rename.HealthErrorCannotRenameDevice("doesnt_exists", device_store.HealthErrorDeviceNotFound.Error())
	failErrorBadRequestInvalid := rename.HealthErrorCannotRenameDevice("device", device_store.HealthErrorInvalidName.Error())
	failErrorBadRequestDuplicated := rename.HealthErrorCannotRenameDevice("device", device_store.HealthErrorDuplicatedName.Error())

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
				failErrorNotFound.Error(),
			),
			err: failErrorNotFound,
		},
		{
			name:    "Failed to rename invalid device name",
			oldName: "device",
			newName: "",
			response: NewControllerError(
				NetStatBadRequest,
				failErrorBadRequestInvalid.Error(),
			),
			err: failErrorBadRequestInvalid,
		},
		{
			name:    "Failed to rename duplicated device name",
			oldName: "device",
			newName: "device_2",
			response: NewControllerError(
				NetStatBadRequest,
				failErrorBadRequestDuplicated.Error(),
			),
			err: failErrorBadRequestDuplicated,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			deviceStoreTest, _ := device_store.NewDeviceStore(device.NewDevice(&mockAddress{}, "device"), device.NewDevice(&mockAddress{}, "device_2"))
			controller := NewController(deviceStoreTest, newTimeMock())

			response, err := controller.Rename(testCase.oldName, testCase.newName)

			if err == nil {
				if err != testCase.err {
					t.Error("Test_Controller_Rename() failed on success test case")
				}
			}

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
