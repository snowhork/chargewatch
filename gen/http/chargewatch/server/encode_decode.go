// Code generated by goa v3.2.6, DO NOT EDIT.
//
// chargewatch HTTP server encoders and decoders
//
// Command:
// $ goa gen chargewatch/design

package server

import (
	chargewatch "chargewatch/gen/chargewatch"
	"context"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListDevicesResponse returns an encoder for responses returned by the
// chargewatch listDevices endpoint.
func EncodeListDevicesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*chargewatch.ListDevicesResult)
		enc := encoder(ctx, w)
		body := NewListDevicesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListDevicesRequest returns a decoder for requests sent to the
// chargewatch listDevices endpoint.
func DecodeListDevicesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID string

			params = mux.Vars(r)
		)
		userID = params["userID"]
		payload := NewListDevicesPayload(userID)

		return payload, nil
	}
}

// EncodeCreateDeviceResponse returns an encoder for responses returned by the
// chargewatch createDevice endpoint.
func EncodeCreateDeviceResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*chargewatch.CreateDeviceResult)
		enc := encoder(ctx, w)
		body := NewCreateDeviceResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateDeviceRequest returns a decoder for requests sent to the
// chargewatch createDevice endpoint.
func DecodeCreateDeviceRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body struct {
				// device name
				Name *string `form:"name" json:"name" xml:"name"`
				// description
				Description *string `form:"description" json:"description" xml:"description"`
			}
			err error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			userID string

			params = mux.Vars(r)
		)
		userID = params["userID"]
		payload := NewCreateDevicePayload(body, userID)

		return payload, nil
	}
}

// EncodeCreateDeviceError returns an encoder for errors returned by the
// createDevice chargewatch endpoint.
func EncodeCreateDeviceError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "StatusInternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateDeviceStatusInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "StatusInternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateChargeResponse returns an encoder for responses returned by the
// chargewatch updateCharge endpoint.
func EncodeUpdateChargeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*chargewatch.UpdateChargeResult)
		enc := encoder(ctx, w)
		body := NewUpdateChargeResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateChargeRequest returns a decoder for requests sent to the
// chargewatch updateCharge endpoint.
func DecodeUpdateChargeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body struct {
				// charge
				Charge *ChargeRequestBodyRequestBody `form:"charge" json:"charge" xml:"charge"`
			}
			err error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			deviceID string

			params = mux.Vars(r)
		)
		deviceID = params["deviceID"]
		payload := NewUpdateChargePayload(body, deviceID)

		return payload, nil
	}
}

// EncodeUpdateChargeError returns an encoder for errors returned by the
// updateCharge chargewatch endpoint.
func EncodeUpdateChargeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "StatusBadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateChargeStatusBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "StatusBadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetChargeHistoryResponse returns an encoder for responses returned by
// the chargewatch getChargeHistory endpoint.
func EncodeGetChargeHistoryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*chargewatch.GetChargeHistoryResult)
		enc := encoder(ctx, w)
		body := NewGetChargeHistoryResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetChargeHistoryRequest returns a decoder for requests sent to the
// chargewatch getChargeHistory endpoint.
func DecodeGetChargeHistoryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			deviceID string

			params = mux.Vars(r)
		)
		deviceID = params["deviceID"]
		payload := NewGetChargeHistoryPayload(deviceID)

		return payload, nil
	}
}

// EncodeUpdateDeviceResponse returns an encoder for responses returned by the
// chargewatch updateDevice endpoint.
func EncodeUpdateDeviceResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*chargewatch.UpdateDeviceResult)
		enc := encoder(ctx, w)
		body := NewUpdateDeviceResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateDeviceRequest returns a decoder for requests sent to the
// chargewatch updateDevice endpoint.
func DecodeUpdateDeviceRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateDeviceRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateDeviceRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			userID   string
			deviceID string

			params = mux.Vars(r)
		)
		userID = params["userID"]
		deviceID = params["deviceID"]
		payload := NewUpdateDevicePayload(&body, userID, deviceID)

		return payload, nil
	}
}

// marshalChargewatchDeviceToDeviceResponseBody builds a value of type
// *DeviceResponseBody from a value of type *chargewatch.Device.
func marshalChargewatchDeviceToDeviceResponseBody(v *chargewatch.Device) *DeviceResponseBody {
	res := &DeviceResponseBody{
		ID:          v.ID,
		UserID:      v.UserID,
		Name:        v.Name,
		Description: v.Description,
		State:       v.State,
	}
	if v.Charge != nil {
		res.Charge = marshalChargewatchChargeToChargeResponseBody(v.Charge)
	}

	return res
}

// marshalChargewatchChargeToChargeResponseBody builds a value of type
// *ChargeResponseBody from a value of type *chargewatch.Charge.
func marshalChargewatchChargeToChargeResponseBody(v *chargewatch.Charge) *ChargeResponseBody {
	res := &ChargeResponseBody{
		Value:     v.Value,
		Charging:  v.Charging,
		Timestamp: v.Timestamp,
	}

	return res
}

// unmarshalChargeRequestBodyRequestBodyToChargewatchCharge builds a value of
// type *chargewatch.Charge from a value of type *ChargeRequestBodyRequestBody.
func unmarshalChargeRequestBodyRequestBodyToChargewatchCharge(v *ChargeRequestBodyRequestBody) *chargewatch.Charge {
	if v == nil {
		return nil
	}
	res := &chargewatch.Charge{
		Value:     *v.Value,
		Charging:  *v.Charging,
		Timestamp: *v.Timestamp,
	}

	return res
}

// marshalChargewatchChargeHistoryToChargeHistoryResponseBody builds a value of
// type *ChargeHistoryResponseBody from a value of type
// *chargewatch.ChargeHistory.
func marshalChargewatchChargeHistoryToChargeHistoryResponseBody(v *chargewatch.ChargeHistory) *ChargeHistoryResponseBody {
	res := &ChargeHistoryResponseBody{
		DeviceID: v.DeviceID,
	}
	if v.Logs != nil {
		res.Logs = make([]*ChargeLogResponseBody, len(v.Logs))
		for i, val := range v.Logs {
			res.Logs[i] = marshalChargewatchChargeLogToChargeLogResponseBody(val)
		}
	}

	return res
}

// marshalChargewatchChargeLogToChargeLogResponseBody builds a value of type
// *ChargeLogResponseBody from a value of type *chargewatch.ChargeLog.
func marshalChargewatchChargeLogToChargeLogResponseBody(v *chargewatch.ChargeLog) *ChargeLogResponseBody {
	res := &ChargeLogResponseBody{
		DeviceID: v.DeviceID,
	}
	if v.Charge != nil {
		res.Charge = marshalChargewatchChargeToChargeResponseBody(v.Charge)
	}

	return res
}
