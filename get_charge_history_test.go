package chargewatchapi

import (
	api "chargewatch/gen/chargewatch"

	"github.com/stretchr/testify/assert"

	"chargewatch/entity"
	"chargewatch/repo/memory"
	"chargewatch/vals"
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChargewatchsrvc_GetChargeHistory(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := &chargewatchsrvc{repo: &RepositoryAgg{Device: memory.NewDeviceMemoryRepository()}}

	userID := vals.NewUserID(fmt.Sprintf("%d", rand.Intn(1000000)))
	meta := vals.GenerateDeviceMeta("iphone", "personal device")
	d := entity.GenerateDevice(userID, meta)
	_ = s.repo.Device.Create(d)

	c1, _ := vals.NewCharge(50, false)
	l1 := vals.NewChargeLog(d.ID, c1)
	time.Sleep(time.Millisecond)
	c2, _ := vals.NewCharge(80, true)
	l2 := vals.NewChargeLog(d.ID, c2)

	_ = s.repo.Device.CreateChargeLog(l1)
	_ = s.repo.Device.CreateChargeLog(l2)

	cases := []struct {
		name     string
		payload  *api.GetChargeHistoryPayload
		expected *api.GetChargeHistoryResult
	}{
		{"got history",
			&api.GetChargeHistoryPayload{
				DeviceID: string(d.ID),
			},
			&api.GetChargeHistoryResult{
				ChargeHistory: &api.ChargeHistory{
					DeviceID: string(d.ID),
					Logs: []*api.ChargeLog{
						{DeviceID: string(d.ID), Charge: &api.Charge{Value: 50, Charging: false, Timestamp: int64(c1.Timestamp)}},
						{DeviceID: string(d.ID), Charge: &api.Charge{Value: 80, Charging: true, Timestamp: int64(c2.Timestamp)}},
					},
				},
			},
		},
		{"unknown deviceID",
			&api.GetChargeHistoryPayload{
				DeviceID: "unknown",
			},
			&api.GetChargeHistoryResult{
				ChargeHistory: &api.ChargeHistory{
					DeviceID: "unknown",
					Logs:     []*api.ChargeLog{},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := s.GetChargeHistory(context.TODO(), c.payload)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, got)
		})
	}
}
