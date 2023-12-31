package device_store

import "network-health/core/entity/device"

type Iterator struct {
	device []*device.Device
}

func (i *Iterator) List() []*device.Device {
	return i.device
}

type DeviceStore struct {
	size    int
	devices map[string]*device.Device
}

func NewDeviceStore(devices ...*device.Device) (*DeviceStore, error) {
	deviceStore := &DeviceStore{}
	deviceStore.size = len(devices)
	deviceStore.devices = make(map[string]*device.Device)

	for _, device := range devices {
		if _, alreadyExists := deviceStore.devices[device.Name()]; alreadyExists {
			return nil, HealthErrorDuplicatedName
		}
		if device.Name() == "" {
			return nil, HealthErrorInvalidName
		}
		deviceStore.devices[device.Name()] = device
	}

	return deviceStore, nil
}

func (store *DeviceStore) RenameDevice(oldName string, newName string) (err error) {
	renamedDevice, found := store.devices[oldName]

	if !found {
		err = HealthErrorDeviceNotFound
		return
	}

	if newName == "" {
		return HealthErrorInvalidName
	}

	if _, alreadyExists := store.devices[newName]; alreadyExists {
		return HealthErrorDuplicatedName
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

// func (store *DeviceStore) AddDevices(devices ...*device.Device) error {
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
