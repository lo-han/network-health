package icmp

import (
	"network-health/core/entity/connectivity"
	entity "network-health/core/entity/device"
	"network-health/core/usecases/check"

	"github.com/go-ping/ping"
)

type ICMPConnectivityHandler struct{}

func NewICMPConnectivityHandler() *ICMPConnectivityHandler {
	return &ICMPConnectivityHandler{}
}

func (ICMPConnectivityHandler) PingDevice(device *entity.Device) (stats connectivity.ConnectionStats, err error) {
	pinger, err := ping.NewPinger(device.Address())

	if err != nil {
		err = check.HealthErrorCannotConnectToServer("ICMP", err.Error())
		return
	}

	pinger.Count = 1
	pinger.Timeout = 2000000000
	pinger.OnRecv = func(pkt *ping.Packet) {
		stats.NBytes = pkt.Nbytes
	}

	pinger.OnFinish = func(statistics *ping.Statistics) {
		stats.PacketsSent = statistics.PacketsSent
		stats.PacketsRecv = statistics.PacketsRecv
		stats.MaxLatency = statistics.MaxRtt
		stats.MinLatency = statistics.MinRtt
		stats.AvgLatency = statistics.AvgRtt
		stats.StdDeviation = statistics.StdDevRtt
	}

	err = pinger.Run()
	if err != nil {
		err = check.HealthErrorServerError(err.Error())
		return
	}

	return
}
