// Code generated by goa v3.2.6, DO NOT EDIT.
//
// chargewatch service
//
// Command:
// $ goa gen chargewatch/design

package chargewatch

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the chargewatch service interface.
type Service interface {
	// ListDevices implements listDevices.
	ListDevices(context.Context, *ListDevicesPayload) (res *ListDevicesResult, err error)
	// CreateDevice implements createDevice.
	CreateDevice(context.Context, *CreateDevicePayload) (res *CreateDeviceResult, err error)
	// UpdateCharge implements updateCharge.
	UpdateCharge(context.Context, *UpdateChargePayload) (res *UpdateChargeResult, err error)
	// GetChargeHistory implements getChargeHistory.
	GetChargeHistory(context.Context, *GetChargeHistoryPayload) (res *GetChargeHistoryResult, err error)
	// UpdateDevice implements updateDevice.
	UpdateDevice(context.Context, *UpdateDevicePayload) (res *UpdateDeviceResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "chargewatch"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"listDevices", "createDevice", "updateCharge", "getChargeHistory", "updateDevice"}

// ListDevicesPayload is the payload type of the chargewatch service
// listDevices method.
type ListDevicesPayload struct {
	// userID
	UserID string
}

// ListDevicesResult is the result type of the chargewatch service listDevices
// method.
type ListDevicesResult struct {
	Devices []*Device
}

// CreateDevicePayload is the payload type of the chargewatch service
// createDevice method.
type CreateDevicePayload struct {
	// userID
	UserID string
	// device name
	Name string
	// description
	Description string
}

// CreateDeviceResult is the result type of the chargewatch service
// createDevice method.
type CreateDeviceResult struct {
	Device *Device
}

// UpdateChargePayload is the payload type of the chargewatch service
// updateCharge method.
type UpdateChargePayload struct {
	// deviceID
	DeviceID string
	// charge
	Charge *Charge
}

// UpdateChargeResult is the result type of the chargewatch service
// updateCharge method.
type UpdateChargeResult struct {
	Device *Device
}

// GetChargeHistoryPayload is the payload type of the chargewatch service
// getChargeHistory method.
type GetChargeHistoryPayload struct {
	// deviceID
	DeviceID string
}

// GetChargeHistoryResult is the result type of the chargewatch service
// getChargeHistory method.
type GetChargeHistoryResult struct {
	ChargeHistory *ChargeHistory
}

// UpdateDevicePayload is the payload type of the chargewatch service
// updateDevice method.
type UpdateDevicePayload struct {
	// userID
	UserID string
	// deviceID
	DeviceID string
	// value [0,100]
	ChargeValue int
}

// UpdateDeviceResult is the result type of the chargewatch service
// updateDevice method.
type UpdateDeviceResult struct {
	Device *Device
}

type Device struct {
	ID          string
	UserID      string
	Name        string
	Description string
	State       string
	Charge      *Charge
}

type Charge struct {
	Value     int
	Charging  bool
	Timestamp int64
}

type ChargeHistory struct {
	DeviceID string
	Logs     []*ChargeLog
}

type ChargeLog struct {
	DeviceID string
	Charge   *Charge
}

// MakeStatusInternalServerError builds a goa.ServiceError from an error.
func MakeStatusInternalServerError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "StatusInternalServerError",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeStatusBadRequest builds a goa.ServiceError from an error.
func MakeStatusBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "StatusBadRequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
