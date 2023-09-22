package entity

type Iterator struct {
	device []*Device
}

func (i *Iterator) List() []*Device {
	return i.device
}

type DeviceStore struct {
	size    int
	devices map[string]*Device
}

func NewDeviceStore(listSize int, devices ...*Device) *DeviceStore {
	deviceStore := &DeviceStore{}
	deviceStore.size = listSize
	deviceStore.devices = make(map[string]*Device)

	for _, device := range devices {
		deviceStore.devices[device.name] = device
	}

	return deviceStore
}

func (store *DeviceStore) RenameDevice(oldName string, newName string) (err error) {
	renamedDevice, found := store.devices[oldName]

	if !found {
		err = HealthErrorDeviceNotFound
		return
	}

	renamedDevice.Rename(newName)

	delete(store.devices, oldName)

	store.devices[newName] = renamedDevice

	return
}

func (store *DeviceStore) IterateDevices() (iteration *Iterator) {
	iteration = new(Iterator)

	for _, device := range store.devices {
		iteration.device = append(iteration.device, device)
	}

	return
}

// func (store *DeviceStore) AddDevices(devices ...*Device) error {
// 	if len(store.devices) == store.size {
// 		return HealthErrorFullDeviceList
// 	}

// 	for _, device := range devices {
// 		store.devices[device.name] = device
// 	}

// 	return nil
// }
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
