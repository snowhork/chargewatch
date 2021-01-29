package vals

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"math/rand"
	"time"
)

//
//type DeviceProp struct {
//	ID          DeviceID
//	UserID      UserID
//	Name        DeviceName
//	Description DeviceDescription
//	State       DeviceState
//	Charge      Charge
//	CreatedAt   Timestamp
//}

type DeviceMeta struct {
	Name        DeviceName
	Description DeviceDescription
	State       DeviceState
}

type DeviceID string
type DeviceName string
type DeviceDescription string
type DeviceState string

func GenerateDeviceID() DeviceID {
	rand.Seed(time.Now().UnixNano())
	return DeviceID(uuid.New().String())
}

func NewDeviceID(id string) DeviceID {
	return DeviceID(id)
}

func GenerateDeviceMeta(name, desc string) DeviceMeta {
	return DeviceMeta{
		Name:        DeviceName(name),
		Description: DeviceDescription(desc),
		State:       DeviceStateInit,
	}
}

func InitCharge() Charge {
	return Charge{
		Value:    ChargeValue(0),
		Charging: false,
	}
}

var DeviceStateInit = DeviceState("ChargeStateInit")
var DeviceStateRunning = DeviceState("ChargeStateRunning")

type Charge struct {
	Value     ChargeValue
	Charging  bool
	Timestamp Timestamp
}
type ChargeValue int

func NewChargeValue(val int) (ChargeValue, error) {
	if val < 0 || val > 100 {
		return ChargeValue(-1), xerrors.Errorf("%w: value must be between 0 and 100", ErrValidation)
	}
	return ChargeValue(val), nil
}

func NewCharge(val int, charging bool) (Charge, error) {
	v, err := NewChargeValue(val)
	if err != nil {
		return Charge{}, err
	}

	return Charge{Value: v, Charging: charging, Timestamp: GenerateCurrentTimestamp()}, nil
}
