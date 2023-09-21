package entity

type Iteration []*Device

type listDevice struct {
	id     string
	device *Device
}

type DeviceStore struct {
	size    int
	devices map[string]listDevice
}

func NewDeviceStore(listSize int, devices ...*Device) *DeviceStore {
	deviceStore := &DeviceStore{}
	deviceStore.size = listSize

	for _, device := range devices {
		deviceStore.devices[device.name] = listDevice{
			id:     device.address.Get(),
			device: device,
		}
	}

	return deviceStore
}

func (store *DeviceStore) AddDevices(devices ...*Device) error {
	if len(store.devices) == store.size {
		return HealthErrorFullDeviceList
	}

	for _, device := range devices {
		store.devices[device.name] = listDevice{
			id:     device.address.Get(),
			device: device,
		}
	}

	return nil
}

func (store *DeviceStore) RenameDevice(oldName string, newName string) (err error) {
	renamedDevice, found := store.devices[oldName]

	if !found {
		err = HealthErrorDeviceNotFound
		return
	}

	renamedDevice.device.Rename(newName)

	delete(store.devices, oldName)

	store.devices[newName] = renamedDevice

	return
}

func (store *DeviceStore) IterateDevices() (devices Iteration) {
	for _, listDevice := range store.devices {
		devices = append(devices, listDevice.device)
	}

	return
}

// func (store *DeviceStore) RemoveDeviceByName(name string) error {
// 	if len(store.devices) == 0 {
// 		return HealthErrorEmptyDeviceList
// 	}

// 	delete(store.devices, name)

// 	return nil
// }

// func (store *DeviceStore) RemoveDeviceByAddress(address string) error {
// 	if len(store.devices) == 0 {
// 		return HealthErrorEmptyDeviceList
// 	}

// 	for name, device := range store.devices {
// 		if device.id == address {
// 			delete(store.devices, name)
// 		}
// 	}

// 	return nil
// }
