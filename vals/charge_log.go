package vals

import "sort"

type ChargeLog struct {
	DeviceID DeviceID
	Charge   Charge
}

func NewChargeLog(deviceID DeviceID, charge Charge) ChargeLog {
	return ChargeLog{
		DeviceID: deviceID,
		Charge:   charge,
	}
}

type ChargeHistory struct {
	DeviceID DeviceID
	Logs     []ChargeLog
}

func NewChargeHistory(deviceID DeviceID, logs []ChargeLog) ChargeHistory {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Charge.Timestamp <= logs[j].Charge.Timestamp
	})
	return ChargeHistory{DeviceID: deviceID, Logs: logs}
}
