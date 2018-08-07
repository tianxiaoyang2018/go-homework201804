package external

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

const (
	ISO8601      = "2006-01-02T15:04:05+0000"
	ISO8601Micro = "2006-01-02T15:04:05.000000+0000"
	TailFlag     = "+0000µ" // identify timestamp field in json

	BotMarkPixel          = "pixel"
	BotMarkXPosed         = "xposed"
	BotMarkPixelAndXPosed = "pixel xposed"
	BotMarkMotion         = "motion"
	BotMarkTime           = "time"
	BotMarkBadPtr         = "bad ptr info"
)

type unixMSCount struct {
	unixMS int64
	count  uint8
}

var userIDTime = struct {
	sync.Mutex
	m map[string]*unixMSCount
}{m: make(map[string]*unixMSCount)}

type Iso8601Time time.Time

func (isoTime *Iso8601Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + (*time.Time)(isoTime).Format(ISO8601Micro) + "µ" + "\""), nil
}

func (isoTime *Iso8601Time) UnmarshalJSON(data []byte) error {
	aLen := len(data)
	if aLen < 2 {
		return fmt.Errorf("Bad Time")
	}
	t, err := ParseIso8601(string(bytes.Trim(data, `µ"'`)))
	if err == nil {
		*isoTime = Iso8601Time(t)
	}
	return err
}

// func (isoTime *Iso8601Time) HasBotMark(userID, clientOS, appVer string, tms int64) (reason string, has bool) {
// 	t := time.Time(*isoTime)
// 	unixMilli := t.UnixNano() / int64(time.Millisecond)

// 	// ios check
// 	if clientOS == "ios" {
// 		if util.VersionGreaterThanOrEqualTo(appVer, "2.3.3") {
// 			if unixMilli%47 == 1 {
// 				return BotMarkPixel, true
// 			}
// 		}

// 		if util.VersionGreaterThanOrEqualTo(appVer, "2.4.0") {
// 			if unixMilli%37 == 1 {
// 				return BotMarkMotion, true
// 			}
// 		}

// 		return "", false
// 	}

// 	// android check
// 	if util.VersionGreaterThanOrEqualTo(appVer, "2.3.2") {
// 		if t.IsZero() {
// 			return BotMarkTime, true
// 		}

// 		if tms != 0 {
// 			if util.AbsInt64(unixMilli-tms)/1000 > int64(config.Conf.MacAccessToken.AllowedTimeDiffInSeconds) {
// 				return BotMarkTime, true
// 			}
// 		}

// 		if unixMilli%47 == 1 {
// 			return BotMarkPixel, true
// 		}
// 		if util.VersionGreaterThanOrEqualTo(appVer, "2.4.0") {
// 			if unixMilli%47 == 1 && unixMilli%7 == 1 {
// 				return BotMarkPixelAndXPosed, true
// 			}

// 			if unixMilli%7 == 1 {
// 				return BotMarkXPosed, true
// 			}
// 		}

// 		if util.VersionGreaterThanOrEqualTo(appVer, "2.5.3") {
// 			if unixMilli%11 == 1 || unixMilli%19 == 1 {
// 				return BotMarkBadPtr, true
// 			}
// 		}

// 		userIDTime.Lock()
// 		defer userIDTime.Unlock()
// 		cachedTime, ok := userIDTime.m[userID]
// 		switch {
// 		case !ok:
// 			userIDTime.m[userID] = &unixMSCount{unixMS: unixMilli}
// 		case cachedTime.unixMS == unixMilli:
// 			cachedTime.count++
// 			if cachedTime.count > 1 {
// 				delete(userIDTime.m, userID)
// 				return BotMarkTime, true
// 			}
// 		default:
// 			delete(userIDTime.m, userID)
// 		}

// 	}

// 	return "", false
// }

func (isoTime *Iso8601Time) Format(fm string) string {
	return time.Time(*isoTime).Format(fm)
}

func ParseIso8601(tm string) (time.Time, error) {
	t, err := time.Parse(ISO8601, tm)
	return t, err
}

func ParseIso8601Micro(tm string) (time.Time, error) {
	t, err := time.Parse(ISO8601Micro, tm)
	return t, err
}
