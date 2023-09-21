package icmp

import (
	"fmt"
	"network-health/core/entity"
	"network-health/core/usecases/check"

	"github.com/go-ping/ping"
)

type ICMPConnectivityHandler struct{}

func NewICMPConnectivityHandler() *ICMPConnectivityHandler {
	return &ICMPConnectivityHandler{}
}

func (ICMPConnectivityHandler) PingDevice(device *entity.Device) (entity.Status, error) {
	pinger, err := ping.NewPinger(device.GetAddress())

	if err != nil {
		return entity.Offline, check.HealthErrorCannotConnectToServer("ICMP")
	}

	pinger.Count = 1
	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		return entity.Offline, check.HealthErrorServerError(err.Error())
	}

	return entity.Online, nil
}
