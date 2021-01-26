package dynamo

import (
	"chargewatch/vals"
)

type chargeItem struct {
	Value     int
	Charging  bool
	Timestamp int64
}

type chargeItemFactory struct {
	c vals.Charge
}

func newChargeItemFactory(c vals.Charge) *chargeItemFactory {
	return &chargeItemFactory{c}
}

func (f *chargeItemFactory) Create() chargeItem {
	return chargeItem{
		Value:     int(f.c.Value),
		Charging:  bool(f.c.Charging),
		Timestamp: int64(f.c.Timestamp),
	}
}

type chargeFactory struct {
	c chargeItem
}

func newChargeFactory(c chargeItem) *chargeFactory {
	return &chargeFactory{c: c}
}

func (f *chargeFactory) Create() vals.Charge {
	return vals.Charge{
		Value:     vals.ChargeValue(f.c.Value),
		Charging:  f.c.Charging,
		Timestamp: vals.Timestamp(f.c.Timestamp),
	}
}
