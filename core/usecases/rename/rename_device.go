package rename

import (
	entity "network-health/core/entity/device_store"
)

func Rename(store *entity.DeviceStore, oldName, newName string) (err RenameUsecaseErrorStack) {
	entityError := store.RenameDevice(oldName, newName)

	if entityError != nil {
		err.Append(entityError)

		usecaseError := HealthErrorCannotRenameDevice(oldName, entityError.Error())

		err.Append(usecaseError)
		return
	}

	return
}
