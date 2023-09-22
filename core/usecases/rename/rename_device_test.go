package rename

import (
	device "network-health/core/entity/device"
	"network-health/core/entity/device_store"
	"testing"
)

type mockAddress struct{}

func (ip *mockAddress) Set(address string) (err error) {
	return
}

func (ip *mockAddress) Get() (address string) {
	return "address"
}

func Test_RenameUsecase_Rename(t *testing.T) {
	oldName := "1"

	testCases := []struct {
		name        string
		oldName     string
		newName     string
		errExpected error
		deviceStore *device_store.DeviceStore
	}{
		{
			name:        "Succesfull rename",
			oldName:     oldName,
			newName:     "2",
			errExpected: nil,
		},
		{
			name:        "Failed rename",
			oldName:     oldName,
			newName:     "",
			errExpected: HealthErrorCannotRenameDevice(oldName, device_store.HealthErrorInvalidName.Error()),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.deviceStore, _ = device_store.NewDeviceStore(device.NewDevice(&mockAddress{}, oldName))

			err := Rename(testCase.deviceStore, testCase.oldName, testCase.newName)

			if err != nil {
				if err.Error() != testCase.errExpected.Error() {
					t.Errorf("Test_RenameUsecase_Rename() = %s, expected %s", err.Error(), testCase.errExpected.Error())
				}
			}
		})
	}
}
