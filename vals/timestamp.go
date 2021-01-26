package vals

import "time"

type Timestamp int64

func GenerateCurrentTimestamp() Timestamp {
	return Timestamp(time.Now().UnixNano())
}
