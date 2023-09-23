package connectivity

import (
	"context"
	"errors"
	device_entity "network-health/core/entity/device"
	"network-health/core/entity/logs"
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

func (connHandlerOnlineMock) PingDevice(device *device_entity.Device) (stats ConnectionStats, err error) {
	stats.PacketsSent = 1
	stats.PacketsRecv = 1
	return
}

type connHandlerOfflineMock struct{}

func (connHandlerOfflineMock) PingDevice(device *device_entity.Device) (stats ConnectionStats, err error) {
	err = errors.New("text")
	return
}

type connHandlerLoadedMock struct{}

func (connHandlerLoadedMock) PingDevice(device *device_entity.Device) (stats ConnectionStats, err error) {
	stats.PacketsSent = 1
	stats.PacketsRecv = 0
	return
}

func Test_Connect(t *testing.T) {
	device_1 := device_entity.NewDevice(&mockAddress{}, "1")

	logs.SetLogger(&mockLogger{})

	testCases := []struct {
		name    string
		handler ConnectionHandler
		status  device_entity.Status
	}{
		{
			name:    "Online if packets received is equal to sent",
			status:  device_entity.Online,
			handler: connHandlerOnlineMock{},
		},
		{
			name:    "Loaded if packets received is different from sent",
			status:  device_entity.Loaded,
			handler: connHandlerLoadedMock{},
		},
		{
			name:    "Online if returned error",
			status:  device_entity.Offline,
			handler: connHandlerOfflineMock{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			status := Connect(testCase.handler, device_1)

			if testCase.status != status {
				t.Errorf("Test_Connect().Status = %d, expected %d", status, testCase.status)
			}
		})
	}
}
