package dynamo

import (
	"chargewatch/entity"
	"chargewatch/repo"
	"chargewatch/vals"

	"github.com/guregu/dynamo"

	"golang.org/x/xerrors"
)

type DeviceDynamoRepository struct {
	client *Client
}

func (r *DeviceDynamoRepository) Get(id vals.DeviceID) (*entity.Device, error) {
	var item deviceItem

	if err := r.client.Table.
		Get(hashKeyName, deviceItemHashKey(id)).
		Range(rangeKeyName, dynamo.Equal, deviceItemRangeKey()).
		One(&item); err != nil {
		if err == dynamo.ErrNotFound {
			return nil, repo.ErrNotFound
		}
		return nil, xerrors.Errorf("err dynamo: %w", err)
	}
	return newDeviceFactory(&item).Create(), nil
}

func (r *DeviceDynamoRepository) ListByUser(userID vals.UserID) (*entity.DeviceList, error) {
	var items []deviceItem

	if err := r.client.Table.
		Get(userItemIndexHashKeyName, userID).
		Range(userItemIndexRangeKeyName, dynamo.Equal, UserItemTypeDevice).
		Index(userItemIndexName).
		All(&items); err != nil {
		return nil, xerrors.Errorf("err dynamo: %w", err)
	}

	res := make([]*entity.Device, len(items))
	for i, item := range items {
		res[i] = newDeviceFactory(&item).Create()
	}

	return entity.NewDeviceList(res), nil
}

func (r *DeviceDynamoRepository) Delete(*entity.Device) error {
	panic("implement me")
}

func (r *DeviceDynamoRepository) Update(d *entity.Device) error {
	_, err := r.Get(d.ID)
	if err != nil {
		return err
	}

	item := newDeviceItemFactory(d).Create()
	if err := r.client.Table.Put(item).Run(); err != nil {
		return xerrors.Errorf("err dynamo: %w", err)
	}

	return nil
}

func NewDeviceDynamoRepository(client *Client) repo.DeviceRepository {
	return &DeviceDynamoRepository{
		client: client,
	}
}

func (r *DeviceDynamoRepository) Create(d *entity.Device) error {
	item := newDeviceItemFactory(d).Create()
	if err := r.client.Table.Put(item).Run(); err != nil {
		return xerrors.Errorf("err dynamo: %w", err)
	}

	return nil
}
