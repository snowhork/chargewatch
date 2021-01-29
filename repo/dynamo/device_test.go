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

func TestNewDeviceDynamoRepository(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	repo := NewDeviceDynamoRepository(NewLocalDynamoClient())

	userID := vals.NewUserID(fmt.Sprintf("%d", rand.Intn(1000000)))
	meta := vals.GenerateDeviceMeta("iphone", "personal device")
	d := entity.GenerateDevice(userID, meta)

	assert.NoError(t, repo.Create(d))

	devices, err := repo.ListByUser(d.UserID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(devices.List))
	assert.Equal(t, d, devices.List[0])

	got, err := repo.Get(d.ID)
	assert.NoError(t, err)
	assert.Equal(t, d, got)

	d2 := entity.GenerateDevice(userID, meta)
	_ = repo.Create(d2)
	devices, _ = repo.ListByUser(d.UserID)
	assert.Equal(t, 2, len(devices.List))
}
