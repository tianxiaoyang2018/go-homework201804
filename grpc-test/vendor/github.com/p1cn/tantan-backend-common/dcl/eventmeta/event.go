package eventmeta

import "time"

// EventMetaData ...
type EventMetaData struct {
	Timestamp  time.Time
	Key, Value []byte
	Topic      string
	Partition  int32
	Offset     int64
}
