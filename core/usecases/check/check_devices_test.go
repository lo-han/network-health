package check

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

type ConnHandlerOnlineMock struct{}

func (ConnHandlerOnlineMock) PingDevice(dev *device.Device) (deviceStatus device.Status) {
	return device.Online
}

type ConnHandlerOfflineMock struct{}

func (ConnHandlerOfflineMock) PingDevice(dev *device.Device) (deviceStatus device.Status) {
	return device.Offline
}

type ConnHandlerLoadedMock struct{}

func (ConnHandlerLoadedMock) PingDevice(dev *device.Device) (deviceStatus device.Status) {
	return device.Loaded
}

func Test_CheckUsecase_Check(t *testing.T) {
	device_1 := device.NewDevice(&mockAddress{}, "1")
	deviceStoreTest, _ := device_store.NewDeviceStore(device_1)

	testCases := []struct {
		name   string
		conn   *Connectivity
		status string
	}{
		{
			name:   "Succesfull online status check",
			conn:   NewConnectivity(ConnHandlerOnlineMock{}),
			status: "ONLINE",
		},
		{
			name:   "Succesfull offline status check",
			conn:   NewConnectivity(ConnHandlerOfflineMock{}),
			status: "OFFLINE",
		},
		{
			name:   "Succesfull loeaded status check",
			conn:   NewConnectivity(ConnHandlerLoadedMock{}),
			status: "LOADED",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			response := testCase.conn.Check(deviceStoreTest)

			device := response.Devices[0]

			if device.Name != device_1.Name() {
				t.Errorf("Test_CheckUsecase_Check().Name = %s, expected %s", device.Name, device_1.Name())
			}

			if device.Address != device_1.Address() {
				t.Errorf("Test_CheckUsecase_Check().Address = %s, expected %s", device.Address, device_1.Address())
			}

			if device.Status != testCase.status {
				t.Errorf("Test_CheckUsecase_Check().Status = %s, expected %s", device.Status, testCase.status)
			}
		})
	}
}
