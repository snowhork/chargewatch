package dynamo

import (
	"chargewatch/vals"

	"github.com/guregu/dynamo"

	"golang.org/x/xerrors"
)

func (r *DeviceDynamoRepository) GetHistory(id vals.DeviceID) (vals.ChargeHistory, error) {
	var items []chargeLogItem

	if err := r.client.Table.
		Get(hashKeyName, chargeLogItemHashKey(id)).
		Range(rangeKeyName, dynamo.GreaterOrEqual, chargeLogItemRangeKey(0)).
		All(&items); err != nil {
		return vals.ChargeHistory{}, xerrors.Errorf("err dynamo: %w", err)
	}

	logs := make([]vals.ChargeLog, len(items))
	for i, item := range items {
		logs[i] = newChargeLogFactory(item).Create()
	}

	return vals.NewChargeHistory(id, logs), nil
}

func (r *DeviceDynamoRepository) CreateChargeLog(log vals.ChargeLog) error {
	item := newChargeLogItemFactory(log).Create()
	if err := r.client.Table.Put(item).Run(); err != nil {
		return xerrors.Errorf("err dynamo: %w", err)
	}

	return nil
}
