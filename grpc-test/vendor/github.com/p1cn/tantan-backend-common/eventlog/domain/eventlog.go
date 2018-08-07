package domain

import (
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/p1cn/tantan-backend-common/eventlog/inner"
	domain_external "github.com/p1cn/tantan-domain-schema/external/golang"
)

// @TODO Don't forget to run ffjson event_log.go if you modified this file
var (
	// eventName contains '  because of client's problem
	eventNameRE   = regexp.MustCompile("^[a-zA-Z0-9_.'-]{0,512}$")
	deviceTokenRE = regexp.MustCompile("^[a-zA-Z0-9-]{0,40}$")
	userIDRE      = regexp.MustCompile("^[0-9]{0,20}$")
	appBuildRE    = regexp.MustCompile("^[a-zA-Z0-9._-]{0,512}$")
)

const (
	TopicEventLog            = "eventlog"
	TopicEventLogContact     = "eventlog-contact"
	TopicClientEventTracking = "client-event-tracking"
	TopicClientTracking      = "client-tracking"
	TopicClientTrackingP10   = "client-tracking-p10"
	TopicClientCrashLog      = "client-crash-log"
	TopicPush                = "push"
	TopicIdlePush            = "idle-user-push"
	TopicMoments             = "moments"
	TopicClientTrackingV4    = "client_tracking_v4"
)

type EventLogRpcMessage struct {
	Topic string
	Event *Event
	ID    int64
	//sync  bool
}

type SearchSettings struct {
	LookingFor string `json:"lookingFor"`
	Radius     int    `json:"radius"`
	MinAge     int    `json:"minAge"`
	MaxAge     int    `json:"maxAge"`
}

type Geo struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type UserRegion struct {
	Country  int `json:"country"`
	Province int `json:"province"`
	City     int `json:"city"`
	District int `json:"district"`
}

func NewUserRegion(Country, Province, City, District string) *UserRegion {
	countryID, _ := strconv.ParseInt(Country, 10, 32)
	provinceID, _ := strconv.ParseInt(Province, 10, 32)
	cityID, _ := strconv.ParseInt(City, 10, 32)
	districtID, _ := strconv.ParseInt(District, 10, 32)
	return &UserRegion{
		Country:  int(countryID),
		Province: int(provinceID),
		City:     int(cityID),
		District: int(districtID),
	}
}

type UserData struct {
	ID             string                         `json:"id"`
	ActiveDay      int                            `json:"activeDay,omitempty"`
	Gender         string                         `json:"gender,omitempty"`
	Status         string                         `json:"status,omitempty"`
	Age            int                            `json:"age,omitempty"`
	SearchSettings *SearchSettings                `json:"searchSettings,omitempty"`
	Geo            *Geo                           `json:"geo,omitempty"`
	Pictures       []int                          `json:"pictures,omitempty"`
	Popularity     float64                        `json:"popularity,omitempty"`
	Regions        interface{}                    `json:"regions,omitempty"`
	TestGroup      map[string]string              `json:"testGroup,omitempty"`
	Boosting       bool                           `json:"boosting,omitempty"`
	BoostConfig    *domain_external.BoostSettings `json:"boost_config,omitempty"`
}

type Interana struct {
	UserID        int64  `json:"user_id,omitempty"`
	InteractionID uint64 `json:"interaction_id,omitempty"`
	Party         string //FIXME: reingest data in interana before changing this `json"party,omitempty"`
	DeviceToken   string `json:"device_token,omitempty"`
}

type Moderator struct {
	Name string `json:"name"`
}

type MobileNumber struct {
	CountryCode int    `json:"countryCode"`
	Number      string `json:"number"`
}

type ActorReceiver struct {
	User         *UserData     `json:"user,omitempty"`
	Device       *DeviceData   `json:"device,omitempty"`
	MobileNumber *MobileNumber `json:"mobileNumber,omitempty"`
	System       *inner.IdType `json:"system,omitempty"`
	Moderator    *Moderator    `json:"moderator,omitempty"`
	Href         *Href         `json:"href,omitempty"`
}

// swipe event -> push service (event.id = cause for push event log)
type Cause struct { // out of scope, later purpose. skip
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type App struct {
	Version string `json:"version"`
	Build   string `json:"build"`
	Channel string `json:"channel"`
}

type Os struct {
	Version string `json:"version"`
	Name    string `json:"name"`
	Family  string `json:"family"`
}

type Href struct {
	Protocol   string      `json:"protocol"`
	Host       string      `json:"host"`
	Path       string      `json:"path"`
	Parameters interface{} `json:"parameters,omitempty"`
}

type Network struct {
	IP       string `json:"ip"`
	Type     string `json:"type"`
	Provider string `json:"provider"`
}

type Model struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	Name         string `json:"name,omitempty"`
}

type DeviceData struct {
	Id          string            `json:"id,omitempty"`
	Token       string            `json:"token,omitempty"`
	Fingerprint inner.IdType      `json:"fingerprint,omitempty"`
	Ip          string            `json:"Ip,omitempty"`
	App         *App              `json:"app,omitempty"`
	Os          *Os               `json:"os,omitempty"`
	AccessToken string            `json:"accessToken,omitempty"`
	Network     *Network          `json:"network,omitempty"`
	Test        bool              `json:"test"`
	TestGroup   map[string]string `json:"testGroup,omitempty"`
	Model       *Model            `json:"model,omitempty"`
}

type PushDeviceInfo struct {
	Provider   string
	Token      string
	AppVersion string
	Os         inner.DeviceOs
}

type PushData struct {
	Title     string
	Value     string
	Intent    string
	Tokens    []PushDeviceInfo
	GroupName string
}

type SmsResponse struct {
	Code        int
	Message     string
	SmsProvider string
	Operator    string
}

type SmsData struct {
	MobileNumber      MobileNumber
	Language          string
	Type              string
	Template          string
	TemplateVariables map[string]string
	Responses         []SmsResponse
	Error             string
}

type Relationship struct {
	Id                string                      `json:"id"`
	Owner             domain_external.IdType      `json:"owner"`
	OtherUser         domain_external.IdType      `json:"otherUser"`
	State             string                      `json:"state"`
	OtherState        string                      `json:"otherState"`
	Category          string                      `json:"category"`
	Scenarios         []domain_external.IdType    `json:"scenarios,omitempty"`
	Status            []string                    `json:"status"`
	CreatedTime       domain_external.Iso8601Time `json:"createdTime"`
	UpdatedTime       domain_external.Iso8601Time `json:"updateTime"`
	ClientCreatedTime domain_external.Iso8601Time `json:"clientCreatedTime"`
	Type              string                      `json:"type"`
	Additional        RelationshipAdditional      `json:"additional"`
}

type Friendship struct {
	Id          string                      `json:"id"`
	Owner       domain_external.IdType      `json:"owner"`
	OtherUser   domain_external.IdType      `json:"otherUser"`
	State       string                      `json:"state"`
	Friendship  bool                        `json:"friendship"`
	OtherState  string                      `json:"otherState"`
	CreatedTime domain_external.Iso8601Time `json:"createdTime"`
	UpdatedTime domain_external.Iso8601Time `json:"updateTime"`
	Type        string                      `json:"type"`
}

type Followship struct {
	Id            string                      `json:"id"`
	Owner         domain_external.IdType      `json:"owner"`
	OtherUser     domain_external.IdType      `json:"otherUser"`
	State         string                      `json:"state"`
	Status        []string                    `json:"-"`
	OtherState    string                      `json:"-"`
	UserTime      domain_external.Iso8601Time `json:"userTime"`
	OtherUserTime domain_external.Iso8601Time `json:"otherUserTime"`
	Type          string                      `json:"type"`
}

type RelationshipAdditional struct {
	Type            string        `json:"type"`
	SuperLikeQuota  int           `json:"superlikeQuota"`
	UndoQuota       int           `json:"undoQuota"`
	OldRelationship *Relationship `json:"oldRelationship,omitempty"`
}

func ExternalRelationshipToRelationship(r *domain_external.Relationship) *Relationship {
	if r == nil {
		return nil
	}
	return &Relationship{
		Id:                r.Id,
		Owner:             r.Owner,
		OtherUser:         r.OtherUser,
		State:             r.State,
		OtherState:        r.OtherState,
		Category:          r.Category,
		Scenarios:         r.Scenarios,
		Status:            r.Status,
		CreatedTime:       r.CreatedTime,
		UpdatedTime:       r.UpdatedTime,
		ClientCreatedTime: r.ClientCreatedTime,
		Type:              r.Type,
		Additional: RelationshipAdditional{
			Type:            r.Additional.Type,
			SuperLikeQuota:  r.Additional.SuperLikeQuota,
			UndoQuota:       r.Additional.UndoQuota,
			OldRelationship: ExternalRelationshipToRelationship(r.Additional.OldRelationship),
		},
	}
}

func ExternalFriendshipToFriendship(f *domain_external.Friendship) *Friendship {
	if f == nil {
		return nil
	}
	return &Friendship{
		Id:          f.Id,
		Owner:       f.Owner,
		OtherUser:   f.OtherUser,
		State:       f.State,
		Friendship:  f.Friendship,
		OtherState:  f.OtherState,
		CreatedTime: f.CreatedTime,
		UpdatedTime: f.UpdatedTime,
		Type:        f.Type,
	}
}

func ExternalFollowshipToFollowship(f *domain_external.Followship) *Followship {
	if f == nil {
		return nil
	}
	return &Followship{
		Id:            f.Id,
		Owner:         f.Owner,
		OtherUser:     f.OtherUser,
		State:         f.State,
		OtherState:    f.OtherState,
		UserTime:      f.UserTime,
		OtherUserTime: f.OtherUserTime,
		Type:          f.Type,
		Status:        f.Status,
	}
}

type UserStatusChangedData struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type UserRoamingLocationChangedData struct {
	From  *UserRegion `json:"from"`
	To    *UserRegion `json:"to"`
	Reset bool        `json:"reset"`
}

type UserChangedData struct {
	Status          interface{} `json:"status,omitempty"`
	RoamingLocation interface{} `json:"roamingLocation,omitempty"`
}

type AndroidData struct {
	UserFrom  string `json:"userFrom"`
	NewDevice bool   `json:"newDevice"`
}
type IOSData struct {
	UserFrom  string `json:"userFrom"`
	NewDevice bool   `json:"newDevice"`
}

type Data struct {
	//Users things
	User   interface{} `json:"user,omitempty"`
	Device interface{} `json:"device,omitempty"`

	//User To OtherUsers things
	Relationship *Relationship `json:"relationship,omitempty"`
	Conversation interface{}   `json:"conversation,omitempty"`
	Message      interface{}   `json:"message,omitempty"`
	Report       interface{}   `json:"report,omitempty"`

	//User Created Things
	Moment        interface{} `json:"moment,omitempty"`
	MomentLike    interface{} `json:"momentLike,omitempty"`
	MomentComment interface{} `json:"momentComment,omitempty"`

	Push   interface{} `json:"push,omitempty"`
	Sms    interface{} `json:"sms,omitempty"`
	AppLog interface{} `json:"applog,omitempty"`

	UserChanged interface{} `json:"userChanged,omitempty"`

	Order interface{} `json:"order,omitempty"`

	Android interface{} `json:"android,omitempty"`
	IOS     interface{} `json:"ios,omitempty"`

	Contacts []inner.Contact `json:"contacts,omitempty"`

	Friendship *Friendship `json:"friendship,omitempty"`
	Poll       interface{} `json:"poll,omitempty"`
	Followship *Followship `json:"followship,omitempty"`
}

type Kafka struct {
	Partition int32
	Offset    int64
}

type Interaction struct {
	Action *struct {
		Name   string
		Method *struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"method,omitempty"`
	} `json:"action,omitempty"`
	View *struct {
		Current *struct {
			Name  string `json:"name"`
			Modal *struct {
				Name string `json:"name"`
			} `json:"modal,omitempty"`
			Popup *struct {
				Name string `json:"name"`
			} `json:"popup,omitempty"`
			Notification *struct {
				Name string `json:"name"`
			} `json:"notification,omitempty"`
		} `json:"current,omitempty"`
	} `json:"view,omitempty"`
}

//
type Event struct {
	ID              string            `json:"id"`
	Kafka           *Kafka            `json:"kafka,omitempty"`
	Interana        *Interana         `json:"shard,omitempty"`
	Version         string            `json:"version"`           // when compiled. OUT OF SCOPE version=1
	Request         inner.IdType      `json:"request,omitempty"` // out of scope
	Foreground      bool              `json:"foreground"`
	Name            string            `json:"name"`      // swipe, match, etc
	Timestamp       int64             `json:"timestamp"` // should be in ms not in seconds    @CONFIRM to int64
	Source          inner.IdType      `json:"source"`    // hostname + port  , service name
	Cause           *Cause            `json:"cause,omitempty"`
	Actor           ActorReceiver     `json:"actor"`
	Receiver        *ActorReceiver    `json:"receiver,omitempty"`
	Action          string            `json:"action,omitempty"`
	RequestDuration int64             `json:"requestDuration"`
	ResponseCode    int               `json:"responseCode"`
	Data            Data              `json:"data"`
	Client          interface{}       `json:"client,omitempty"`
	Test            bool              `json:"test"`
	TestGroup       map[string]string `json:"testGroup,omitempty"`
	Interaction     *Interaction      `json:"interaction,omitempty"`
}

func (e Event) IsValid() bool {

	if !eventNameRE.Match([]byte(e.Name)) {
		return false
	}

	if e.Actor.Device != nil && !deviceTokenRE.Match([]byte(e.Actor.Device.Token)) {
		return false
	}

	if e.Receiver != nil && e.Receiver.Device != nil && !deviceTokenRE.Match([]byte(e.Receiver.Device.Token)) {
		return false
	}

	if e.Actor.Device == nil || e.Actor.Device.Os == nil ||
		len(e.Actor.Device.Os.Name) == 0 {
		return false
	}

	return true
}

func (e Event) HasUserID() bool {
	if e.Actor.User != nil && len(e.Actor.User.ID) > 0 {
		return true
	}
	return false
}

func NewEvent(sourceID string) (*Event, error) {

	uuID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	event := Event{
		ID:      uuID.String(),
		Version: "1.1.1",
		Request: inner.IdType{
			Id:   "",
			Type: "event",
		},

		Timestamp: time.Now().UTC().UnixNano() / int64(time.Millisecond),
		Source: inner.IdType{ //
			Id: sourceID,
		},
	}
	return &event, nil
}

type EventLogs struct {
	Events []Event `json:"events"`
}

func (c *EventLogs) IsValid() bool {
	for _, e := range c.Events {
		if !e.IsValid() {
			return false
		}
	}
	return true
}

func (c *EventLogs) HasUserID() bool {
	for i := 0; i < len(c.Events); i++ {
		if c.Events[i].Actor.User != nil && len(c.Events[i].Actor.User.ID) > 0 {
			return true
		}
	}
	return false
}
