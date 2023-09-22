package device_store

import (
	"network-health/core/entity/device"
	"testing"
)

type mockAddress struct{}

func (ip *mockAddress) Set(address string) (err error) {
	return
}

func (ip *mockAddress) Get() (address string) {
	return "address"
}

func Test_DeviceStore_RenameDevice(t *testing.T) {
	deviceStoreTest, _ := NewDeviceStore(
		device.NewDevice(&mockAddress{}, "device_2"),
		device.NewDevice(&mockAddress{}, "device_3"),
	)

	t.Run("Couldn't instantiate same name devices", func(t *testing.T) {
		_, err := NewDeviceStore(
			device.NewDevice(&mockAddress{}, "device_2"),
			device.NewDevice(&mockAddress{}, "device_2"),
		)

		if err == nil {
			t.Error("Test_DeviceStore_RenameDevice() allowed same name devices!")
		}
	})

	t.Run("Couldn't instantiate invalid device name", func(t *testing.T) {
		_, err := NewDeviceStore(
			device.NewDevice(&mockAddress{}, ""),
		)

		if err == nil {
			t.Error("Test_DeviceStore_RenameDevice() allowed invalid device name!")
		}
	})

	successTestCases := []struct {
		name    string
		store   *DeviceStore
		newName string
		oldName string
	}{
		{
			name:    "Succesfull rename",
			store:   deviceStoreTest,
			newName: "another_device",
			oldName: "device_3",
		},
	}

	for _, testCase := range successTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			_ = testCase.store.RenameDevice(testCase.oldName, testCase.newName)

			device, exists := testCase.store.devices[testCase.newName]
			if !exists {
				t.Errorf("Test_DeviceStore_RenameDevice() didn't store device index!")
			}

			if !(device.Name() == testCase.newName) {
				t.Errorf("Test_DeviceStore_RenameDevice() didn't renamed device name!")
			}
		})
	}

	failTestCases := []struct {
		name    string
		store   *DeviceStore
		newName string
		oldName string
	}{
		{
			name:    "Device doesn't exists",
			store:   deviceStoreTest,
			oldName: "doesnt_exists",
			newName: "another_device",
		},
		{
			name:    "Name already exists",
			store:   deviceStoreTest,
			oldName: "device_3",
			newName: "device_2",
		},
	}

	for _, testCase := range failTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.store.RenameDevice(testCase.oldName, testCase.newName)

			if err == nil {
				t.Errorf("Test_DeviceStore_RenameDevice() = nil, want %v", err)
			}
		})
	}

}

func Test_DeviceStore_IterateDevices(t *testing.T) {
	device_2 := device.NewDevice(&mockAddress{}, "device_2")
	device_3 := device.NewDevice(&mockAddress{}, "device_3")
	deviceStoreTest, _ := NewDeviceStore(
		device_2,
		device_3,
	)

	testCase := []struct {
		name  string
		store *DeviceStore
	}{
		{
			name:  "Successfully returned a iterable object",
			store: deviceStoreTest,
		},
	}

	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			list := testCase.store.IterateDevices().List()

			for _, device := range list {
				if device != device_2 && device != device_3 {
					t.Errorf("Test_DeviceStore_IterateDevices() '%s' object not in original list!", device.Name())
				}
			}
		})
	}
}
