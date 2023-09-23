package check

import (
	"context"
	"errors"
	"network-health/core/entity/connectivity"
	device_entity "network-health/core/entity/device"
	"network-health/core/entity/device_store"
	"network-health/core/entity/logs"
	time_usecase "network-health/core/usecases/time"
	"testing"
)

type mockLogger struct {
}

func (logger *mockLogger) Context(ctx context.Context) logs.Logger {
	return &mockLogger{}
}

func (logger *mockLogger) Error(message string) {
}

func (logger *mockLogger) Fatal(message string) {
}

func (logger *mockLogger) Info(message string) {
}

type mockAddress struct{}

func (ip *mockAddress) Set(address string) (err error) {
	return
}

func (ip *mockAddress) Get() (address string) {
	return "address"
}

type connHandlerOnlineMock struct{}

func (connHandlerOnlineMock) PingDevice(device *device_entity.Device) (stats connectivity.ConnectionStats, err error) {
	stats.PacketsSent = 1
	stats.PacketsRecv = 1
	return
}

type connHandlerOfflineMock struct{}

func (connHandlerOfflineMock) PingDevice(device *device_entity.Device) (stats connectivity.ConnectionStats, err error) {
	err = errors.New("text")
	return
}

type connHandlerLoadedMock struct{}

func (connHandlerLoadedMock) PingDevice(device *device_entity.Device) (stats connectivity.ConnectionStats, err error) {
	stats.PacketsSent = 1
	stats.PacketsRecv = 0
	return
}

func Test_CheckUsecase_Check(t *testing.T) {
	device_1 := device_entity.NewDevice(&mockAddress{}, "1")
	deviceStoreTest, _ := device_store.NewDeviceStore(device_1)

	logs.SetLogger(&mockLogger{})

	testCases := []struct {
		name   string
		conn   *Connectivity
		status string
	}{
		{
			name:   "Succesfull online status check",
			conn:   NewConnectivity(connHandlerOnlineMock{}, &time_usecase.GoTime{}),
			status: "ONLINE",
		},
		{
			name:   "Succesfull offline status check",
			conn:   NewConnectivity(connHandlerOfflineMock{}, &time_usecase.GoTime{}),
			status: "OFFLINE",
		},
		{
			name:   "Succesfull loeaded status check",
			conn:   NewConnectivity(connHandlerLoadedMock{}, &time_usecase.GoTime{}),
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
