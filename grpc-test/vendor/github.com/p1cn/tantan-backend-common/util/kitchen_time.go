package util

import (
	"fmt"
	"time"
)

type KitchenTime time.Time

func (self *KitchenTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + (*time.Time)(self).Format(time.Kitchen) + "\""), nil
}

func (self *KitchenTime) UnmarshalJSON(data []byte) error {
	aLen := len(data)
	if aLen < 2 {
		return fmt.Errorf("Bad Time")
	}

	t, err := time.ParseInLocation(time.Kitchen, string(data[1:aLen-1]), time.Local)

	if err == nil {
		now := time.Now()
		t = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local)
		*self = KitchenTime(t)
	}
	return err
}
