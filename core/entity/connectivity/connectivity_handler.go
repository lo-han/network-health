package connectivity

import (
	"fmt"
	device_entity "network-health/core/entity/device"
	"network-health/core/entity/logs"
	"time"
)

type ConnectionStats struct {
	PacketsSent  int
	PacketsRecv  int
	NBytes       int
	MinLatency   time.Duration
	AvgLatency   time.Duration
	MaxLatency   time.Duration
	StdDeviation time.Duration
}

type ConnectionHandler interface {
	PingDevice(device *device_entity.Device) (stats ConnectionStats, err error)
}

func Connect(connection ConnectionHandler, device *device_entity.Device) (deviceStatus device_entity.Status) {
	logs.Gateway().Info(fmt.Sprintf("PING %s", device.Address()))

	stats, err := connection.PingDevice(device)

	logs.Gateway().Info(fmt.Sprintf("%d bytes from %s time=%v",
		stats.NBytes, device.Address(), time.Now()))

	logs.Gateway().Info(fmt.Sprintf("\tround-trip min/avg/max/stddev = %v/%v/%v/%v",
		stats.MinLatency, stats.AvgLatency, stats.MaxLatency, stats.StdDeviation))

	if err != nil {
		deviceStatus = device_entity.Offline
		logs.Gateway().Error(fmt.Sprintf("%s", err.Error()))
		return
	}

	if stats.PacketsSent != stats.PacketsRecv {
		deviceStatus = device_entity.Loaded
		return
	}

	deviceStatus = device_entity.Online

	return
}
