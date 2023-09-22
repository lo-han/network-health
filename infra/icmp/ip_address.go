package icmp

import (
	entity "network-health/core/entity/device"
	"strconv"
	"strings"
)

type IPv4Address struct {
	addr    string
	version string
}

func NewIPv4Address(address string) (*IPv4Address, error) {
	ipv4address := &IPv4Address{
		version: "IPv4",
	}

	if err := ipv4address.Set(address); err != nil {
		return nil, err
	}

	return ipv4address, nil
}

func (ip *IPv4Address) Set(address string) (err error) {
	addrBytes := strings.Split(address, ".")

	if len(addrBytes) != 4 {
		return entity.HealthErrorInvalidAddress(ip.version)
	}

	for _, subnet := range addrBytes {
		value, conversionErr := strconv.Atoi(subnet)

		if conversionErr != nil {
			return entity.HealthErrorInvalidAddress(ip.version)
		}

		if value >= 255 || value < 0 {
			return entity.HealthErrorInvalidAddress(ip.version)
		}
	}
	ip.addr = address

	return
}

func (ip *IPv4Address) Get() (address string) {
	return ip.addr
}
