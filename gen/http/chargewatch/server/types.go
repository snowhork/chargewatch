// Code generated by goa v3.2.6, DO NOT EDIT.
//
// chargewatch HTTP server types
//
// Command:
// $ goa gen chargewatch/design

package server

import (
	chargewatch "chargewatch/gen/chargewatch"

	goa "goa.design/goa/v3/pkg"
)

// UpdateDeviceRequestBody is the type of the "chargewatch" service
// "updateDevice" endpoint HTTP request body.
type UpdateDeviceRequestBody struct {
	// value [0,100]
	ChargeValue *int `form:"chargeValue,omitempty" json:"chargeValue,omitempty" xml:"chargeValue,omitempty"`
}

// ListDevicesResponseBody is the type of the "chargewatch" service
// "listDevices" endpoint HTTP response body.
type ListDevicesResponseBody struct {
	Devices []*DeviceResponseBody `form:"devices" json:"devices" xml:"devices"`
}

// CreateDeviceResponseBody is the type of the "chargewatch" service
// "createDevice" endpoint HTTP response body.
type CreateDeviceResponseBody struct {
	Device *DeviceResponseBody `form:"device" json:"device" xml:"device"`
}

// UpdateChargeResponseBody is the type of the "chargewatch" service
// "updateCharge" endpoint HTTP response body.
type UpdateChargeResponseBody struct {
	Device *DeviceResponseBody `form:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
}

// GetChargeHistoryResponseBody is the type of the "chargewatch" service
// "getChargeHistory" endpoint HTTP response body.
type GetChargeHistoryResponseBody struct {
	ChargeHistory *ChargeHistoryResponseBody `form:"chargeHistory" json:"chargeHistory" xml:"chargeHistory"`
}

// UpdateDeviceResponseBody is the type of the "chargewatch" service
// "updateDevice" endpoint HTTP response body.
type UpdateDeviceResponseBody struct {
	Device *DeviceResponseBody `form:"device" json:"device" xml:"device"`
}

// CreateDeviceStatusInternalServerErrorResponseBody is the type of the
// "chargewatch" service "createDevice" endpoint HTTP response body for the
// "StatusInternalServerError" error.
type CreateDeviceStatusInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateChargeStatusBadRequestResponseBody is the type of the "chargewatch"
// service "updateCharge" endpoint HTTP response body for the
// "StatusBadRequest" error.
type UpdateChargeStatusBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeviceResponseBody is used to define fields on response body types.
type DeviceResponseBody struct {
	ID          string              `form:"id" json:"id" xml:"id"`
	UserID      string              `form:"userID" json:"userID" xml:"userID"`
	Name        string              `form:"name" json:"name" xml:"name"`
	Description string              `form:"description" json:"description" xml:"description"`
	State       string              `form:"state" json:"state" xml:"state"`
	Charge      *ChargeResponseBody `form:"charge" json:"charge" xml:"charge"`
}

// ChargeResponseBody is used to define fields on response body types.
type ChargeResponseBody struct {
	Value     int   `form:"value" json:"value" xml:"value"`
	Charging  bool  `form:"charging" json:"charging" xml:"charging"`
	Timestamp int64 `form:"timestamp" json:"timestamp" xml:"timestamp"`
}

// ChargeHistoryResponseBody is used to define fields on response body types.
type ChargeHistoryResponseBody struct {
	DeviceID string                   `form:"deviceID" json:"deviceID" xml:"deviceID"`
	Logs     []*ChargeLogResponseBody `form:"logs" json:"logs" xml:"logs"`
}

// ChargeLogResponseBody is used to define fields on response body types.
type ChargeLogResponseBody struct {
	DeviceID string              `form:"deviceID" json:"deviceID" xml:"deviceID"`
	Charge   *ChargeResponseBody `form:"charge" json:"charge" xml:"charge"`
}

// ChargeRequestBodyRequestBody is used to define fields on request body types.
type ChargeRequestBodyRequestBody struct {
	Value     *int   `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	Charging  *bool  `form:"charging,omitempty" json:"charging,omitempty" xml:"charging,omitempty"`
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

// NewListDevicesResponseBody builds the HTTP response body from the result of
// the "listDevices" endpoint of the "chargewatch" service.
func NewListDevicesResponseBody(res *chargewatch.ListDevicesResult) *ListDevicesResponseBody {
	body := &ListDevicesResponseBody{}
	if res.Devices != nil {
		body.Devices = make([]*DeviceResponseBody, len(res.Devices))
		for i, val := range res.Devices {
			body.Devices[i] = marshalChargewatchDeviceToDeviceResponseBody(val)
		}
	}
	return body
}

// NewCreateDeviceResponseBody builds the HTTP response body from the result of
// the "createDevice" endpoint of the "chargewatch" service.
func NewCreateDeviceResponseBody(res *chargewatch.CreateDeviceResult) *CreateDeviceResponseBody {
	body := &CreateDeviceResponseBody{}
	if res.Device != nil {
		body.Device = marshalChargewatchDeviceToDeviceResponseBody(res.Device)
	}
	return body
}

// NewUpdateChargeResponseBody builds the HTTP response body from the result of
// the "updateCharge" endpoint of the "chargewatch" service.
func NewUpdateChargeResponseBody(res *chargewatch.UpdateChargeResult) *UpdateChargeResponseBody {
	body := &UpdateChargeResponseBody{}
	if res.Device != nil {
		body.Device = marshalChargewatchDeviceToDeviceResponseBody(res.Device)
	}
	return body
}

// NewGetChargeHistoryResponseBody builds the HTTP response body from the
// result of the "getChargeHistory" endpoint of the "chargewatch" service.
func NewGetChargeHistoryResponseBody(res *chargewatch.GetChargeHistoryResult) *GetChargeHistoryResponseBody {
	body := &GetChargeHistoryResponseBody{}
	if res.ChargeHistory != nil {
		body.ChargeHistory = marshalChargewatchChargeHistoryToChargeHistoryResponseBody(res.ChargeHistory)
	}
	return body
}

// NewUpdateDeviceResponseBody builds the HTTP response body from the result of
// the "updateDevice" endpoint of the "chargewatch" service.
func NewUpdateDeviceResponseBody(res *chargewatch.UpdateDeviceResult) *UpdateDeviceResponseBody {
	body := &UpdateDeviceResponseBody{}
	if res.Device != nil {
		body.Device = marshalChargewatchDeviceToDeviceResponseBody(res.Device)
	}
	return body
}

// NewCreateDeviceStatusInternalServerErrorResponseBody builds the HTTP
// response body from the result of the "createDevice" endpoint of the
// "chargewatch" service.
func NewCreateDeviceStatusInternalServerErrorResponseBody(res *goa.ServiceError) *CreateDeviceStatusInternalServerErrorResponseBody {
	body := &CreateDeviceStatusInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateChargeStatusBadRequestResponseBody builds the HTTP response body
// from the result of the "updateCharge" endpoint of the "chargewatch" service.
func NewUpdateChargeStatusBadRequestResponseBody(res *goa.ServiceError) *UpdateChargeStatusBadRequestResponseBody {
	body := &UpdateChargeStatusBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListDevicesPayload builds a chargewatch service listDevices endpoint
// payload.
func NewListDevicesPayload(userID string) *chargewatch.ListDevicesPayload {
	v := &chargewatch.ListDevicesPayload{}
	v.UserID = userID

	return v
}

// NewCreateDevicePayload builds a chargewatch service createDevice endpoint
// payload.
func NewCreateDevicePayload(body struct {
	// device name
	Name *string `form:"name" json:"name" xml:"name"`
	// description
	Description *string `form:"description" json:"description" xml:"description"`
}, userID string) *chargewatch.CreateDevicePayload {
	v := &chargewatch.CreateDevicePayload{}
	if body.Name != nil {
		v.Name = *body.Name
	}
	if body.Description != nil {
		v.Description = *body.Description
	}
	if body.Description == nil {
		v.Description = ""
	}
	v.UserID = userID

	return v
}

// NewUpdateChargePayload builds a chargewatch service updateCharge endpoint
// payload.
func NewUpdateChargePayload(body struct {
	// charge
	Charge *ChargeRequestBodyRequestBody `form:"charge" json:"charge" xml:"charge"`
}, deviceID string) *chargewatch.UpdateChargePayload {
	v := &chargewatch.UpdateChargePayload{}
	if body.Charge != nil {
		v.Charge = unmarshalChargeRequestBodyRequestBodyToChargewatchCharge(body.Charge)
	}
	v.DeviceID = deviceID

	return v
}

// NewGetChargeHistoryPayload builds a chargewatch service getChargeHistory
// endpoint payload.
func NewGetChargeHistoryPayload(deviceID string) *chargewatch.GetChargeHistoryPayload {
	v := &chargewatch.GetChargeHistoryPayload{}
	v.DeviceID = deviceID

	return v
}

// NewUpdateDevicePayload builds a chargewatch service updateDevice endpoint
// payload.
func NewUpdateDevicePayload(body *UpdateDeviceRequestBody, userID string, deviceID string) *chargewatch.UpdateDevicePayload {
	v := &chargewatch.UpdateDevicePayload{
		ChargeValue: *body.ChargeValue,
	}
	v.UserID = userID
	v.DeviceID = deviceID

	return v
}

// ValidateUpdateDeviceRequestBody runs the validations defined on
// UpdateDeviceRequestBody
func ValidateUpdateDeviceRequestBody(body *UpdateDeviceRequestBody) (err error) {
	if body.ChargeValue == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("chargeValue", "body"))
	}
	if body.ChargeValue != nil {
		if *body.ChargeValue < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.chargeValue", *body.ChargeValue, 0, true))
		}
	}
	if body.ChargeValue != nil {
		if *body.ChargeValue > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.chargeValue", *body.ChargeValue, 100, false))
		}
	}
	return
}

// ValidateChargeRequestBodyRequestBody runs the validations defined on
// ChargeRequestBodyRequestBody
func ValidateChargeRequestBodyRequestBody(body *ChargeRequestBodyRequestBody) (err error) {
	if body.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "body"))
	}
	if body.Charging == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("charging", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	return
}
