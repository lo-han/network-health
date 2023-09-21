package check

import (
	"time"
)

type Device struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Status  string `json:"status"`
}

type DeviceStatus struct {
	Devices []Device `json:"devices"`

	Datetime time.Time `json:"datetime"`
}
