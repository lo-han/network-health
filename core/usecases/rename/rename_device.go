package rename

import (
	entity "network-health/core/entity/device_list"
)

func Rename(store *entity.DeviceStore, oldName, newName string) (err error) {
	err = store.RenameDevice(oldName, newName)

	if err != nil {
		err = HealthErrorCannotRenameDevice(oldName)
		return
	}

	return
}
