package entity

type Status int

type Address interface {
	Set(address string) (err error)
	Get() (address string)
}

const (
	Online Status = iota
	Offline
	Loaded
)

type Device struct {
	address Address
	name    string
	status  Status
}

func NewDevice(address Address, name string) *Device {
	return &Device{
		address: address,
		name:    name,
		status:  Loaded,
	}
}

func (device *Device) GetName() string {
	return device.name
}

func (device *Device) GetAddress() string {
	return device.address.Get()
}

func (device *Device) GetStatus() Status {
	return device.status
}

func (device *Device) Rename(newName string) {
	device.name = newName
}

func (device *Device) SetStatus(status Status) {
	device.status = status
}
