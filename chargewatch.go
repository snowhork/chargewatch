package chargewatchapi

import (
	"chargewatch/entity"
	api "chargewatch/gen/chargewatch"
	"chargewatch/repo/dynamo"
	"chargewatch/vals"
	"context"
	"log"

	"golang.org/x/xerrors"
)

type chargewatchsrvc struct {
	logger *log.Logger
	repo   *RepositoryAgg
}

func device2Payload(d *entity.Device) *api.Device {
	return &api.Device{
		ID:          string(d.ID),
		UserID:      string(d.UserID),
		Name:        string(d.Meta.Name),
		Description: string(d.Meta.Description),
		State:       string(d.Meta.State),
		Charge:      charge2Payload(d.Charge),
	}
}

func charge2Payload(c vals.Charge) *api.Charge {
	return &api.Charge{
		Value:     int(c.Value),
		Charging:  bool(c.Charging),
		Timestamp: int64(c.Timestamp),
	}
}

func history2Payload(h vals.ChargeHistory) *api.ChargeHistory {
	return &api.ChargeHistory{
		DeviceID: string(h.DeviceID),
		Logs:     logs2Payload(h.Logs),
	}
}

func logs2Payload(logs []vals.ChargeLog) []*api.ChargeLog {
	res := make([]*api.ChargeLog, len(logs))
	for i, l := range logs {
		res[i] = &api.ChargeLog{
			DeviceID: string(l.DeviceID),
			Charge:   charge2Payload(l.Charge),
		}
	}

	return res
}

func (s *chargewatchsrvc) ListDevices(ctx context.Context, p *api.ListDevicesPayload) (*api.ListDevicesResult, error) {
	devices, err := s.repo.Device.ListByUser(vals.NewUserID(p.UserID))
	if err != nil {
		return nil, api.MakeStatusInternalServerError(xerrors.Errorf("%w", err))
	}

	return genListDevicesResult(devices.List), nil
}

func (s *chargewatchsrvc) UpdateCharge(ctx context.Context, p *api.UpdateChargePayload) (*api.UpdateChargeResult, error) {
	charge, err := updateChargePayload2Charge(p)
	if err != nil {
		return nil, api.MakeStatusBadRequest(xerrors.Errorf("%w", err))
	}

	device, err := s.repo.Device.Get(vals.NewDeviceID(p.DeviceID))
	if err != nil {
		return nil, api.MakeStatusInternalServerError(xerrors.Errorf("%w", err))
	}

	chargeLog, err := device.UpdateCharge(charge)
	if err != nil {
		return nil, api.MakeStatusInternalServerError(xerrors.Errorf("%w", err))
	}

	if err := s.repo.Device.Update(device); err != nil {
		return nil, api.MakeStatusInternalServerError(xerrors.Errorf("%w", err))
	}

	if err := s.repo.Device.CreateChargeLog(chargeLog); err != nil {
		return nil, api.MakeStatusInternalServerError(xerrors.Errorf("%w", err))
	}

	return &api.UpdateChargeResult{
		Device: device2Payload(device),
	}, nil
}

func updateChargePayload2Charge(p *api.UpdateChargePayload) (vals.Charge, error) {
	return vals.NewCharge(p.ChargeValue, p.Charging)
}

func genListDevicesResult(ds []*entity.Device) *api.ListDevicesResult {
	res := &api.ListDevicesResult{
		Devices: make([]*api.Device, len(ds)),
	}

	for i, d := range ds {
		res.Devices[i] = device2Payload(d)
	}

	return res
}

func (*chargewatchsrvc) UpdateDevice(context.Context, *api.UpdateDevicePayload) (res *api.UpdateDeviceResult, err error) {
	panic("implement me")
}

func NewChargewatch(logger *log.Logger) api.Service {
	return &chargewatchsrvc{
		logger: logger,
		repo: &RepositoryAgg{
			Device: dynamo.NewDeviceDynamoRepository(),
		},
	}
}
