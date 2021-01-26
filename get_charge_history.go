package chargewatchapi

import (
	api "chargewatch/gen/chargewatch"
	"chargewatch/vals"
	"context"

	"golang.org/x/xerrors"
)

func (s *chargewatchsrvc) GetChargeHistory(ctx context.Context, p *api.GetChargeHistoryPayload) (*api.GetChargeHistoryResult, error) {
	id := vals.NewDeviceID(p.DeviceID)

	history, err := s.repo.Device.GetHistory(id)
	if err != nil {
		return nil, api.MakeStatusInternalServerError(xerrors.Errorf("%w", err))
	}

	return &api.GetChargeHistoryResult{
		ChargeHistory: history2Payload(history),
	}, nil
}
