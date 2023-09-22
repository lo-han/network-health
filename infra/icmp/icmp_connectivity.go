package icmp

import (
	"errors"
	"fmt"
	entity "network-health/core/entity/device"
	"network-health/core/usecases/check"
	"time"

	"github.com/go-ping/ping"
)

type ICMPConnectivityHandler struct{}

func NewICMPConnectivityHandler() *ICMPConnectivityHandler {
	return &ICMPConnectivityHandler{}
}

func (ICMPConnectivityHandler) PingDevice(device *entity.Device) (deviceStatus entity.Status) {
	pinger, err := ping.NewPinger(device.Address())

	err = errors.New("text")
	if err != nil {
		deviceStatus = entity.Offline
		fmt.Printf(check.HealthErrorCannotConnectToServer("ICMP", err.Error()).Error())
		return
	}

	pinger.Count = 1
	pinger.Timeout = 2000000000
	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, time.Now())
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		if stats.PacketsSent != stats.PacketsRecv {
			deviceStatus = entity.Loaded
		}

		fmt.Printf("\tround-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("\nPING %s (%s):", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		deviceStatus = entity.Offline
		fmt.Printf(check.HealthErrorServerError(err.Error()).Error())
		return
	}

	return
}
