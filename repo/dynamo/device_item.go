package dynamo

import (
	"chargewatch/entity"
	"chargewatch/vals"
	"fmt"
)

type deviceItem struct {
	HashKey  HashKey
	RangeKey RangeKey

	DeviceID    string
	UserID      UserID // UserItemIndex,hash
	Name        string
	Description string
	State       string
	CreatedAt   int64
	Charge      chargeItem

	UserItemType UserItemType // UserItemIndex,range
}

var UserItemTypeDevice = UserItemType("device")

type deviceItemFactory struct {
	d *entity.Device
}

func newDeviceItemFactory(d *entity.Device) *deviceItemFactory {
	return &deviceItemFactory{d}
}

func (f *deviceItemFactory) Create() deviceItem {
	d := f.d

	return deviceItem{
		HashKey:  deviceItemHashKey(d.ID),
		RangeKey: deviceItemRangeKey(),

		DeviceID:    string(d.ID),
		UserID:      UserID(d.UserID),
		Name:        string(d.Meta.Name),
		Description: string(d.Meta.Description),
		State:       string(d.Meta.State),
		CreatedAt:   int64(d.CreatedAt),
		Charge:      newChargeItemFactory(d.Charge).Create(),

		UserItemType: UserItemTypeDevice,
	}
}

func deviceItemHashKey(id vals.DeviceID) HashKey {
	return HashKey(fmt.Sprintf("device:%s", id))
}

func deviceItemRangeKey() RangeKey {
	return 0
}

type deviceFactory struct {
	i *deviceItem
}

func newDeviceFactory(i *deviceItem) *deviceFactory {
	return &deviceFactory{i}
}

func (f *deviceFactory) Create() *entity.Device {
	i := f.i

	return &entity.Device{
		ID:        vals.DeviceID(i.DeviceID),
		UserID:    vals.UserID(i.UserID),
		CreatedAt: vals.Timestamp(i.CreatedAt),
		Meta: vals.DeviceMeta{
			Name:        vals.DeviceName(i.Name),
			Description: vals.DeviceDescription(i.Description),
			State:       vals.DeviceState(i.State),
		},
		Charge: newChargeFactory(i.Charge).Create(),
	}
}
