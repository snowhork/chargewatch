package chargewatchapi

import (
	api "chargewatch/gen/chargewatch"
	"chargewatch/repo/memory"
	"chargewatch/vals"

	"github.com/stretchr/testify/assert"

	"context"
	"testing"
)

func TestChargewatchsrvc_CreateDevice(t *testing.T) {
	cases := []struct {
		name     string
		p        *api.CreateDevicePayload
		expected *api.CreateDeviceResult
	}{
		{
			"name",
			&api.CreateDevicePayload{UserID: "nick", Name: "iphone", Description: "desc"},
			&api.CreateDeviceResult{Device: &api.Device{
				UserID:      "nick",
				Name:        "iphone",
				Description: "desc",
				State:       "ChargeStateInit",
				Charge: &api.Charge{
					Value:    0,
					Charging: false,
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := &chargewatchsrvc{repo: &RepositoryAgg{Device: memory.NewDeviceMemoryRepository()}}

			got, err := s.CreateDevice(context.Background(), c.p)
			assert.NoError(t, err)

			deviceID := got.Device.ID
			c.expected.Device.ID = deviceID
			assert.Equal(t, c.expected, got)

			list, _ := s.repo.Device.ListByUser(vals.NewUserID(c.p.UserID))
			device, _ := s.repo.Device.Get(vals.DeviceID(deviceID))
			assert.Equal(t, device, list.List[0])
		})
	}
}
