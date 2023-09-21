package check

import (
	"network-health/core/entity"
	"time"
)

type Device struct {
	Name    string        `json:"name"`
	Address string        `json:"address"`
	Status  entity.Status `json:"status"`
}

type DeviceStatus struct {
	Devices []Device `json:"devices"`

	Datetime time.Time `json:"datetime"`
}
