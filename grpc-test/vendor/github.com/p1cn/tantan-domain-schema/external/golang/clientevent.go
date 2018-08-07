package external

import (
	"fmt"
	"regexp"
	"time"
)

type ClientEventTime time.Time

func (isoTime *ClientEventTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + (*time.Time)(isoTime).Format(ISO8601Micro) + "\""), nil
}

func (isoTime *ClientEventTime) UnmarshalJSON(data []byte) error {
	aLen := len(data)
	if aLen < 2 {
		return fmt.Errorf("Bad Time")
	}
	t, err := ParseIso8601(string(data[1 : aLen-1]))
	if err == nil {
		*isoTime = ClientEventTime(t)
	}
	return err
}

func (isoTime *ClientEventTime) Format(fm string) string {
	return time.Time(*isoTime).Format(fm)
}

type CrashDetail struct {
	MemoryFree int32
	CpuUsed    float32
	DiskFree   int32
	Battery    float32
	Stack      string
}

type PageDetail struct {
	Name       string
	LoadTime   int64
	RenderTime int64
	LiveTime   int64
}
type OperationDetail struct {
	Name        string
	AbnormalFPS int32
}

type PermissionDetail struct {
	Locate      int8
	Push        int8
	AddressBook int8
	Camera      int8
	Album       int8
	MicroPhone  int8
}

type NetworkDetail struct {
	Url       string
	Method    string
	Provider  string
	Type      string
	ErrorCode int32
	Duration  int64
}

type MediaDetail struct {
	Url          string
	Method       string
	Provider     string
	Type         string
	Duration     int64
	ErrorCode    int32
	FailProgress float32
}

type PushDetail struct {
	StartTime       ClientEventTime
	EndTime         ClientEventTime
	ForegroundCount int32
	BackgroundCount int32
	Sdk             string
	Delay2sCount    int32
	Delay5sCount    int32
	Delay10sCount   int32
	Delay1mCount    int32
	Delay10mCount   int32
	Delay20mCount   int32
	Delay1hCount    int32
	Delay4hCount    int32
	Delay12hCount   int32
}

type IdlePushDetail struct {
	StartTime       ClientEventTime
	EndTime         ClientEventTime
	ForegroundCount int32
	BackgroundCount int32
	Sdk             string
	Delay2sCount    int32
	Delay5sCount    int32
	Delay10sCount   int32
	Delay1mCount    int32
	Delay10mCount   int32
	Delay20mCount   int32
	Delay1hCount    int32
	Delay4hCount    int32
	Delay12hCount   int32
}

type EventTrack struct {
	Common struct {
		UserID       string
		DeviceID     string
		Manufacturer string
		Model        string
		Resolution   struct {
			W int16
			H int16
		}
		Platform      string
		SystemVersion string
		AppVersion    string
		Channel       string
		Time          ClientEventTime
		Location      struct {
			Longitude float64
			Latitude  float64
		}
		Language string
		Locale   string
	}
	Details struct {
		Crash      *CrashDetail      `json:"crash,omitempty"`
		Page       *PageDetail       `json:"page,omitempty"`
		Operation  *OperationDetail  `json:"operation,omitempty"`
		Network    *NetworkDetail    `json:"network,omitempty"`
		Permission *PermissionDetail `json:"permission,omitempty"`
		Media      *MediaDetail      `json:"media,omitempty"`
		Push       *PushDetail       `json:"push,omitempty"`
		IdlePush   *IdlePushDetail   `json:"idlepush,omitempty"`
	}
}

var (
	// eventName contains '  because of client's problem
	eventNameRE   = regexp.MustCompile("^[a-zA-Z0-9_.'-]{1,512}$")
	deviceTokenRE = regexp.MustCompile("^[a-zA-Z0-9-]{1,40}$")
	userIDRE      = regexp.MustCompile("^[0-9]{0,20}$")
	appBuildRE    = regexp.MustCompile("^[a-zA-Z0-9._-]{1,512}$")
)

func ValidateEventName(e string) bool {
	return eventNameRE.MatchString(e)
}

func ValidateDeviceToken(t string) bool {
	return deviceTokenRE.MatchString(t)
}

func ValidateUserID(u string) bool {
	return userIDRE.MatchString(u)
}

func ValidateAppBuild(a string) bool {
	return appBuildRE.MatchString(a)
}

// @TODO not enough
func (e *EventTrack) Validate() bool {
	return ValidateDeviceToken(e.Common.DeviceID) && ValidateUserID(e.Common.UserID)
}

type EventTracks struct {
	Events []EventTrack
}

func (e *EventTracks) Validate() bool {
	for _, ee := range e.Events {
		if !ee.Validate() {
			return false
		}
	}
	return true
}

////////////////////////////////

type ClientEventApp struct {
	Version string `json:"version"`
	Build   string `json:"build"`
	Channel string `json:"channel"`
}
type ClientEventOS struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}
type ClientEventNetwork struct {
	IP       string `json:"ip"`
	Provider string `json:"provider"`
	Type     string `json:"type"`
}

type ClientEventModel struct {
	Name string `json:"name"`
}

type ClientEventDevice struct {
	Token     string              `json:"token"`
	APP       *ClientEventApp     `json:"app,omitempty"`
	Model     *ClientEventModel   `json:"model,omitempty"` // reserved
	OS        *ClientEventOS      `json:"os,omitempty"`
	Network   *ClientEventNetwork `json:"network,omitempty"`
	Test      bool                `json:"test"`
	TestGroup map[string]string   `json:"testgroup,omitempty"`
	Data      map[string]string   `json:"data,omitempty"` // ignore it for now.
}

type ClientEventMobileNumber struct {
	CountryCode int    `json:"countrycode"`
	Number      string `json:"number"`
}
type ClientEventUser struct {
	ID string `json:"id"`
}

type ClientEventActor struct {
	User         *ClientEventUser         `json:"user,omitempty"`
	Device       *ClientEventDevice       `json:"device,omitempty"`
	MobileNumber *ClientEventMobileNumber `json:"mobilenumber,omitempty"`
}

type ClientEventTrack struct {
	Name           string    `json:"name"`
	ForeGround     bool      `json:"foreground"`
	Time           time.Time `json:"time"`
	RequestTime    time.Time `json:"requesttime"`
	BackendRequest struct {
		ID string `json:"id"`
	} `json:"backendrequest"`
	Actor    ClientEventActor  `json:"actor"`
	Receiver *ClientEventActor `json:"receiver,omitempty"`
}

// @TODO not enough
func (c *ClientEventTrack) Validate() bool {
	if !ValidateEventName(c.Name) {
		return false
	}

	if c.Actor.Device != nil && !ValidateDeviceToken(c.Actor.Device.Token) {
		return false
	}

	if c.Receiver != nil && c.Receiver.Device != nil && !ValidateDeviceToken(c.Receiver.Device.Token) {
		return false
	}

	if c.Actor.Device == nil || c.Actor.Device.OS == nil ||
		len(c.Actor.Device.OS.Name) == 0 {
		return false
	}

	return true
}

type ClientEventTracks struct {
	Events []ClientEventTrack `json:"events"`
}

func (c *ClientEventTracks) Validate() bool {
	for _, e := range c.Events {
		if !e.Validate() {
			return false
		}
	}
	return true
}

func (c *ClientEventTracks) HasUserID() bool {
	for i := 0; i < len(c.Events); i++ {
		if c.Events[i].Actor.User != nil && len(c.Events[i].Actor.User.ID) > 0 {
			return true
		}
	}
	return false
}
