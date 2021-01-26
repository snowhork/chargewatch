package memory

import (
	"chargewatch/entity"
	"chargewatch/repo"
	"chargewatch/vals"
)

type DeviceMemoryRepository struct {
	deviceMap map[vals.DeviceID]*entity.Device
	userMap   map[vals.UserID][]*entity.Device
	logsMap   map[vals.DeviceID][]vals.ChargeLog
}

func (r *DeviceMemoryRepository) GetHistory(id vals.DeviceID) (vals.ChargeHistory, error) {
	logs := r.logsMap[id]
	return vals.NewChargeHistory(id, logs), nil
}

func (r *DeviceMemoryRepository) CreateChargeLog(log vals.ChargeLog) error {
	r.logsMap[log.DeviceID] = append(r.logsMap[log.DeviceID], log)
	return nil
}

func (r *DeviceMemoryRepository) Get(id vals.DeviceID) (*entity.Device, error) {
	if d, ok := r.deviceMap[id]; ok {
		return d, nil
	}
	return nil, repo.ErrNotFound
}

func (r *DeviceMemoryRepository) ListByUser(id vals.UserID) (*entity.DeviceList, error) {
	return entity.NewDeviceList(r.userMap[id]), nil
}

func (r *DeviceMemoryRepository) Create(d *entity.Device) error {
	r.deviceMap[d.ID] = d
	r.userMap[d.UserID] = append(r.userMap[d.UserID], d)

	return nil
}

func (*DeviceMemoryRepository) Update(*entity.Device) error {
	panic("implement me")
}

func (*DeviceMemoryRepository) Delete(*entity.Device) error {
	panic("implement me")
}

func NewDeviceMemoryRepository() repo.DeviceRepository {
	return &DeviceMemoryRepository{
		deviceMap: map[vals.DeviceID]*entity.Device{},
		userMap:   map[vals.UserID][]*entity.Device{},
		logsMap:   map[vals.DeviceID][]vals.ChargeLog{},
	}
}
