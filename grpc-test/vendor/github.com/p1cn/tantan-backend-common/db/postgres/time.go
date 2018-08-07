package postgres

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/pg.v3/pgutil"
)

func toDbTime(t time.Time) time.Time {
	// We dont allow dates before year 1000, since those times might
	// not be valid dates in the database.
	unix := time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC)
	if t.Before(unix) {
		t = unix
	}
	_, offset := time.Now().Zone()
	t = t.UTC().Add(time.Second * time.Duration(-offset))
	return t
}

func fromDbTime(t time.Time) time.Time {
	_, offset := time.Now().Zone()
	t = t.UTC().Add(time.Second * time.Duration(offset))
	return t
}

func toDbTimes(args ...interface{}) []interface{} {
	for i, arg := range args {
		switch t := arg.(type) {
		case Time:
			args[i] = toDbTime(time.Time(t))
		case *Time:
			if t != nil {
				t2 := toDbTime(time.Time(*t))
				args[i] = &t2
			}
		case time.Time:
			args[i] = toDbTime(t)
		case *time.Time:
			if t != nil {
				t2 := toDbTime(*t)
				args[i] = &t2
			}
		default:
		}
	}
	return args
}

type Time time.Time

func (t *Time) Scan(src interface{}) error {
	if src == nil {
		*t = Time{}
		return nil
	}
	switch b := src.(type) {
	case []byte:
		srct, err := pgutil.ParseTime(b)
		if err != nil {
			return err
		}
		srct = fromDbTime(srct)
		*t = Time(srct)
	default:
		return fmt.Errorf("Wrong type")
	}

	return nil
}

func (t *Time) String() string {
	return time.Time(*t).String()
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*t))
}

func (t Time) UnixNano() int64 {
	return time.Time(t).UnixNano()
}

func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

func toInterval(seconds int) string {
	return strconv.Itoa(seconds)
}
