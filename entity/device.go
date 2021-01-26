package entity

import (
	"chargewatch/vals"
	"sort"
)

type Device struct {
	ID        vals.DeviceID
	UserID    vals.UserID
	Meta      vals.DeviceMeta
	Charge    vals.Charge
	CreatedAt vals.Timestamp
}

func GenerateDevice(userID vals.UserID, meta vals.DeviceMeta) *Device {
	return &Device{
		ID:        vals.GenerateDeviceID(),
		UserID:    userID,
		Meta:      meta,
		Charge:    vals.InitCharge(),
		CreatedAt: vals.GenerateCurrentTimestamp(),
	}
}

func (d *Device) UpdateCharge(charge vals.Charge) (vals.ChargeLog, error) {
	d.Charge = charge
	return vals.NewChargeLog(d.ID, charge), nil
}

type DeviceList struct {
	List []*Device
}

func NewDeviceList(list []*Device) *DeviceList {
	res := &DeviceList{List: make([]*Device, len(list))}
	copy(res.List, list)

	sort.Slice(res.List, func(i, j int) bool {
		return res.List[i].CreatedAt >= res.List[j].CreatedAt
	})

	return res
}
