package dynamo

import (
	"chargewatch/vals"
	"fmt"
)

type chargeLogItem struct {
	HashKey  HashKey
	RangeKey RangeKey

	DeviceID string
	Charge   chargeItem
}

type chargeLogItemFactory struct {
	l vals.ChargeLog
}

func newChargeLogItemFactory(l vals.ChargeLog) *chargeLogItemFactory {
	return &chargeLogItemFactory{l: l}
}

func (f *chargeLogItemFactory) Create() chargeLogItem {
	l := f.l

	return chargeLogItem{
		HashKey:  chargeLogItemHashKey(l.DeviceID),
		RangeKey: chargeLogItemRangeKey(l.Charge.Timestamp),

		DeviceID: string(l.DeviceID),
		Charge:   newChargeItemFactory(l.Charge).Create(),
	}
}

func chargeLogItemHashKey(id vals.DeviceID) HashKey {
	return HashKey(fmt.Sprintf("deviceChargeLog:%s", id))
}

func chargeLogItemRangeKey(t vals.Timestamp) RangeKey {
	return RangeKey(t)
}

type chargeLogFactory struct {
	i chargeLogItem
}

func newChargeLogFactory(i chargeLogItem) *chargeLogFactory {
	return &chargeLogFactory{i: i}
}

func (f *chargeLogFactory) Create() vals.ChargeLog {
	i := f.i

	return vals.ChargeLog{
		DeviceID: vals.DeviceID(i.DeviceID),
		Charge:   newChargeFactory(i.Charge).Create(),
	}
}
