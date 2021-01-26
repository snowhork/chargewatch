package repo

import (
	"chargewatch/entity"
	"chargewatch/vals"
)

type DeviceRepository interface {
	Get(id vals.DeviceID) (*entity.Device, error)
	GetHistory(id vals.DeviceID) (vals.ChargeHistory, error)
	ListByUser(id vals.UserID) (*entity.DeviceList, error)
	Create(device *entity.Device) error
	CreateChargeLog(log vals.ChargeLog) error
	Update(device *entity.Device) error
	Delete(device *entity.Device) error
}
