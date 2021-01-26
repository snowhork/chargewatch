// Code generated by goa v3.2.6, DO NOT EDIT.
//
// chargewatch HTTP client encoders and decoders
//
// Command:
// $ goa gen chargewatch/design

package client

import (
	"bytes"
	chargewatch "chargewatch/gen/chargewatch"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildListDevicesRequest instantiates a HTTP request object with method and
// path set to call the "chargewatch" service "listDevices" endpoint
func (c *Client) BuildListDevicesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID string
	)
	{
		p, ok := v.(*chargewatch.ListDevicesPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chargewatch", "listDevices", "*chargewatch.ListDevicesPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListDevicesChargewatchPath(userID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chargewatch", "listDevices", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListDevicesResponse returns a decoder for responses returned by the
// chargewatch listDevices endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeListDevicesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListDevicesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chargewatch", "listDevices", err)
			}
			err = ValidateListDevicesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chargewatch", "listDevices", err)
			}
			res := NewListDevicesResultOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chargewatch", "listDevices", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateDeviceRequest instantiates a HTTP request object with method and
// path set to call the "chargewatch" service "createDevice" endpoint
func (c *Client) BuildCreateDeviceRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID string
	)
	{
		p, ok := v.(*chargewatch.CreateDevicePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chargewatch", "createDevice", "*chargewatch.CreateDevicePayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateDeviceChargewatchPath(userID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chargewatch", "createDevice", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateDeviceRequest returns an encoder for requests sent to the
// chargewatch createDevice server.
func EncodeCreateDeviceRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chargewatch.CreateDevicePayload)
		if !ok {
			return goahttp.ErrInvalidType("chargewatch", "createDevice", "*chargewatch.CreateDevicePayload", v)
		}
		body := p
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chargewatch", "createDevice", err)
		}
		return nil
	}
}

// DecodeCreateDeviceResponse returns a decoder for responses returned by the
// chargewatch createDevice endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCreateDeviceResponse may return the following errors:
//	- "StatusInternalServerError" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeCreateDeviceResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateDeviceResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chargewatch", "createDevice", err)
			}
			err = ValidateCreateDeviceResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chargewatch", "createDevice", err)
			}
			res := NewCreateDeviceResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CreateDeviceStatusInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chargewatch", "createDevice", err)
			}
			err = ValidateCreateDeviceStatusInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chargewatch", "createDevice", err)
			}
			return nil, NewCreateDeviceStatusInternalServerError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chargewatch", "createDevice", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateChargeRequest instantiates a HTTP request object with method and
// path set to call the "chargewatch" service "updateCharge" endpoint
func (c *Client) BuildUpdateChargeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		deviceID string
	)
	{
		p, ok := v.(*chargewatch.UpdateChargePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chargewatch", "updateCharge", "*chargewatch.UpdateChargePayload", v)
		}
		deviceID = p.DeviceID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateChargeChargewatchPath(deviceID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chargewatch", "updateCharge", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateChargeRequest returns an encoder for requests sent to the
// chargewatch updateCharge server.
func EncodeUpdateChargeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chargewatch.UpdateChargePayload)
		if !ok {
			return goahttp.ErrInvalidType("chargewatch", "updateCharge", "*chargewatch.UpdateChargePayload", v)
		}
		body := NewStructChargeChargeChargeRequestBodyRequestBodyForm(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chargewatch", "updateCharge", err)
		}
		return nil
	}
}

// DecodeUpdateChargeResponse returns a decoder for responses returned by the
// chargewatch updateCharge endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeUpdateChargeResponse may return the following errors:
//	- "StatusBadRequest" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeUpdateChargeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateChargeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chargewatch", "updateCharge", err)
			}
			err = ValidateUpdateChargeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chargewatch", "updateCharge", err)
			}
			res := NewUpdateChargeResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body UpdateChargeStatusBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chargewatch", "updateCharge", err)
			}
			err = ValidateUpdateChargeStatusBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chargewatch", "updateCharge", err)
			}
			return nil, NewUpdateChargeStatusBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chargewatch", "updateCharge", resp.StatusCode, string(body))
		}
	}
}

// BuildGetChargeHistoryRequest instantiates a HTTP request object with method
// and path set to call the "chargewatch" service "getChargeHistory" endpoint
func (c *Client) BuildGetChargeHistoryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		deviceID string
	)
	{
		p, ok := v.(*chargewatch.GetChargeHistoryPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chargewatch", "getChargeHistory", "*chargewatch.GetChargeHistoryPayload", v)
		}
		deviceID = p.DeviceID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetChargeHistoryChargewatchPath(deviceID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chargewatch", "getChargeHistory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetChargeHistoryResponse returns a decoder for responses returned by
// the chargewatch getChargeHistory endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetChargeHistoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetChargeHistoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chargewatch", "getChargeHistory", err)
			}
			err = ValidateGetChargeHistoryResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chargewatch", "getChargeHistory", err)
			}
			res := NewGetChargeHistoryResultOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chargewatch", "getChargeHistory", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateDeviceRequest instantiates a HTTP request object with method and
// path set to call the "chargewatch" service "updateDevice" endpoint
func (c *Client) BuildUpdateDeviceRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID   string
		deviceID string
	)
	{
		p, ok := v.(*chargewatch.UpdateDevicePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chargewatch", "updateDevice", "*chargewatch.UpdateDevicePayload", v)
		}
		userID = p.UserID
		deviceID = p.DeviceID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateDeviceChargewatchPath(userID, deviceID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chargewatch", "updateDevice", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateDeviceRequest returns an encoder for requests sent to the
// chargewatch updateDevice server.
func EncodeUpdateDeviceRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chargewatch.UpdateDevicePayload)
		if !ok {
			return goahttp.ErrInvalidType("chargewatch", "updateDevice", "*chargewatch.UpdateDevicePayload", v)
		}
		body := NewUpdateDeviceRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chargewatch", "updateDevice", err)
		}
		return nil
	}
}

// DecodeUpdateDeviceResponse returns a decoder for responses returned by the
// chargewatch updateDevice endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeUpdateDeviceResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateDeviceResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chargewatch", "updateDevice", err)
			}
			err = ValidateUpdateDeviceResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chargewatch", "updateDevice", err)
			}
			res := NewUpdateDeviceResultOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chargewatch", "updateDevice", resp.StatusCode, string(body))
		}
	}
}

// unmarshalDeviceResponseBodyToChargewatchDevice builds a value of type
// *chargewatch.Device from a value of type *DeviceResponseBody.
func unmarshalDeviceResponseBodyToChargewatchDevice(v *DeviceResponseBody) *chargewatch.Device {
	res := &chargewatch.Device{
		ID:          *v.ID,
		UserID:      *v.UserID,
		Name:        *v.Name,
		Description: *v.Description,
		State:       *v.State,
	}
	res.Charge = unmarshalChargeResponseBodyToChargewatchCharge(v.Charge)

	return res
}

// unmarshalChargeResponseBodyToChargewatchCharge builds a value of type
// *chargewatch.Charge from a value of type *ChargeResponseBody.
func unmarshalChargeResponseBodyToChargewatchCharge(v *ChargeResponseBody) *chargewatch.Charge {
	res := &chargewatch.Charge{
		Value:     *v.Value,
		Charging:  *v.Charging,
		Timestamp: *v.Timestamp,
	}

	return res
}

// marshalChargewatchChargeToChargeRequestBodyRequestBody builds a value of
// type *ChargeRequestBodyRequestBody from a value of type *chargewatch.Charge.
func marshalChargewatchChargeToChargeRequestBodyRequestBody(v *chargewatch.Charge) *ChargeRequestBodyRequestBody {
	res := &ChargeRequestBodyRequestBody{
		Value:     v.Value,
		Charging:  v.Charging,
		Timestamp: v.Timestamp,
	}

	return res
}

// marshalChargeRequestBodyRequestBodyToChargewatchCharge builds a value of
// type *chargewatch.Charge from a value of type *ChargeRequestBodyRequestBody.
func marshalChargeRequestBodyRequestBodyToChargewatchCharge(v *ChargeRequestBodyRequestBody) *chargewatch.Charge {
	if v == nil {
		return nil
	}
	res := &chargewatch.Charge{
		Value:     v.Value,
		Charging:  v.Charging,
		Timestamp: v.Timestamp,
	}

	return res
}

// unmarshalChargeHistoryResponseBodyToChargewatchChargeHistory builds a value
// of type *chargewatch.ChargeHistory from a value of type
// *ChargeHistoryResponseBody.
func unmarshalChargeHistoryResponseBodyToChargewatchChargeHistory(v *ChargeHistoryResponseBody) *chargewatch.ChargeHistory {
	res := &chargewatch.ChargeHistory{
		DeviceID: *v.DeviceID,
	}
	res.Logs = make([]*chargewatch.ChargeLog, len(v.Logs))
	for i, val := range v.Logs {
		res.Logs[i] = unmarshalChargeLogResponseBodyToChargewatchChargeLog(val)
	}

	return res
}

// unmarshalChargeLogResponseBodyToChargewatchChargeLog builds a value of type
// *chargewatch.ChargeLog from a value of type *ChargeLogResponseBody.
func unmarshalChargeLogResponseBodyToChargewatchChargeLog(v *ChargeLogResponseBody) *chargewatch.ChargeLog {
	res := &chargewatch.ChargeLog{
		DeviceID: *v.DeviceID,
	}
	res.Charge = unmarshalChargeResponseBodyToChargewatchCharge(v.Charge)

	return res
}
