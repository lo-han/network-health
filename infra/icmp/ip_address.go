package icmp

import (
	"network-health/core/entity"
	"strconv"
	"strings"
)

type IPv4Address struct {
	addr    string
	version string
}

func NewIPv4Address(address string) *IPv4Address {
	return &IPv4Address{
		addr:    address,
		version: "IPv4",
	}
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

	return
}

func (ip *IPv4Address) Get() (address string) {
	return ip.addr
}
