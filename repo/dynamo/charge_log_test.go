package dynamo

import (
	"chargewatch/entity"
	"chargewatch/vals"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDeviceDynamoRepository_Charge(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	repo := NewDeviceDynamoRepository()

	userID := vals.NewUserID(fmt.Sprintf("%d", rand.Intn(1000000)))
	meta := vals.GenerateDeviceMeta("iphone", "personal device")
	d := entity.GenerateDevice(userID, meta)
	_ = repo.Create(d)

	c1, _ := vals.NewCharge(50, true)
	l1 := vals.NewChargeLog(d.ID, c1)
	time.Sleep(time.Millisecond)
	c2, _ := vals.NewCharge(80, true)
	l2 := vals.NewChargeLog(d.ID, c2)

	err := repo.CreateChargeLog(l1)
	assert.NoError(t, err)
	err = repo.CreateChargeLog(l2)
	assert.NoError(t, err)

	history, err := repo.GetHistory(d.ID)
	assert.NoError(t, err)
	assert.Equal(t, d.ID, history.DeviceID)
	assert.Equal(t, 2, len(history.Logs))
	assert.Equal(t, l1, history.Logs[0])
	assert.Equal(t, l2, history.Logs[1])

}
