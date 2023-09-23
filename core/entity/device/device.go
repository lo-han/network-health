package device

type Status int

type Address interface {
	Set(address string) (err error)
	Get() (address string)
}

const (
	Offline Status = iota
	Loaded
	Online
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

func (device *Device) Name() string {
	return device.name
}

func (device *Device) Rename(newName string) {
	device.name = newName
}

func (device *Device) Address() string {
	return device.address.Get()
}

func (device *Device) Status() Status {
	return device.status
}

func (device *Device) SetStatus(status Status) {
	device.status = status
}
