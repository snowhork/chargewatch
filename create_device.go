package chargewatchapi

import (
	api "chargewatch/gen/chargewatch"

	"chargewatch/entity"
	"chargewatch/vals"
	"context"

	"golang.org/x/xerrors"
)

func (s *chargewatchsrvc) CreateDevice(ctx context.Context, p *api.CreateDevicePayload) (*api.CreateDeviceResult, error) {
	device := createDevicePayload2Device(p)
	if err := s.repo.Device.Create(device); err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, api.MakeStatusInternalServerError(err)
	}

	return genCreateDeviceResult(device), nil
}

func createDevicePayload2Device(p *api.CreateDevicePayload) *entity.Device {
	userID := vals.NewUserID(p.UserID)
	meta := vals.GenerateDeviceMeta(p.Name, p.Description)

	return entity.GenerateDevice(userID, meta)
}

func genCreateDeviceResult(d *entity.Device) *api.CreateDeviceResult {
	return &api.CreateDeviceResult{
		Device: device2Payload(d),
	}
}
