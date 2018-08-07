package gin

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/p1cn/tantan-backend-common/config"
	log "github.com/p1cn/tantan-backend-common/contextlog"
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/util"
	"github.com/p1cn/tantan-backend-common/util/constant"
	service "github.com/p1cn/tantan-domain-schema/golang/common"
)

var (
	ErrBadAccessToken    = fmt.Errorf("Bad access token")
	ErrBadMethodOverride = fmt.Errorf("Bad Method Override")
	ErrUnknownResource   = fmt.Errorf("Unknown resource")
)

const (
	ISO8601      = "2006-01-02T15:04:05+0000"
	ISO8601Micro = "2006-01-02T15:04:05.000000+0000"
)

const (
	AuthTypeBasic  = "Basic"
	AuthTypeBearer = "Bearer"
	AuthTypeMac    = "MAC"

	MagicUserIDForAntispam              = 53060000
	MaxFailedTimesForBuildInfoHashCheck = 7
)

func NewRequest(req *http.Request) (*Request, error) {
	var request Request
	err := request.parse(req)
	if err != nil {
		slog.Notice("%v", err)
		return &request, err
	}
	return &request, nil
}

type Geolocation struct {
	Longitude float64
	Latitude  float64
}

type Pagination struct {
	Since  string
	Until  string
	Offset int
	Limit  int
}

type Parameters struct {
	Search string
	Filter string
	Query  string
	Sort   string
}

type Request struct {
	UserID         string
	Method         string
	URI            string
	ContentType    string
	AcceptLanguage string
	Referer        string
	UserAgent      string
	ClientOS       string
	ClientState    string
	AppVersion     string
	Geolocation    *log.GeoUri
	AuthType       string
	ETag           string

	TraceID string
	//Parameters  Parameters
	//BuildInfoHash      string
	//SecretVersion      string
	//ContentLength      int
	//HMACTimestampMs    int64
	//Pagination         Pagination
	//With               int
	AccessToken     string
	RawTestingGroup string
}

func (self *Request) parse(req *http.Request) error {
	self.Method = req.Method

	ctx, _ := req.Context().Value(constant.KeyServiceContext).(*service.Context)
	if ctx != nil {
		self.TraceID = ctx.GetTraceID()
	}

	err := self.parseHeaders(req)
	if err != nil {
		return err
	}
	err = self.parseQuery(req)
	if err != nil {
		return err
	}

	return nil
}

func (self *Request) parseHeaders(req *http.Request) error {
	self.UserID = req.Header.Get(constant.HeaderUserID)
	self.AuthType = req.Header.Get(constant.HeaderAuthType)

	self.ContentType = req.Header.Get("Content-Type")
	self.UserAgent = req.Header.Get("User-Agent")
	self.ETag = req.Header.Get("If-None-Match")

	os, appVersion := util.ParseUserAgent(self.UserAgent)
	self.ClientOS = os
	self.AppVersion = appVersion

	self.AccessToken = ParseAuthBearerAccessToken(req.Header)
	self.RawTestingGroup = req.Header.Get(constant.HeaderTestingGroup)

	geolocation := req.Header.Get("Geolocation")
	if geolocation != "" {
		var g util.GeoUri
		err := g.Parse(geolocation)
		if err != nil {
			return err
		}
		if !g.IsDefaultLocation() {
			self.Geolocation = &log.GeoUri{
				Latitude:  g.Latitude,
				Longitude: g.Longitude,
				CordC:     g.CordC,
			}
			// set geo precision
			self.Geolocation.Latitude = util.SetFloatPrecision(self.Geolocation.Latitude, 4)
			self.Geolocation.Longitude = util.SetFloatPrecision(self.Geolocation.Longitude, 4)
		}
	}

	self.AcceptLanguage = util.ParseLanguage(req.Header.Get("Accept-Language"))

	self.Referer = req.Header.Get("Referer")
	self.ClientState = req.Header.Get("Client-State")

	return nil
}

func (self *Request) parseQuery(req *http.Request) error {
	values := req.URL.Query()

	if config.GetCommonConfig().Debug.BackdoorEnabled {
		userId := values.Get("user_id")
		if userId != "" {
			self.UserID = userId
		}
	}

	return nil
}

func (self *Request) recursiveParseResource(path []string, resource *Resource) error {
	if len(path) == 1 {
		return nil
	}
	value := path[len(path)-1]
	if value == "" {
		return ErrUnknownResource
	}
	resourceType, found := allowedResources[value]
	if !found {
		if resource.Id != nil {
			return ErrUnknownResource
		}
		resource.Id = &value
		return self.recursiveParseResource(path[0:len(path)-1], resource)
	}
	resource.Type = resourceType
	if len(path) == 2 {
		return nil
	}
	resource.Parent = &Resource{}
	return self.recursiveParseResource(path[0:len(path)-1], resource.Parent)
}
func ParseIso8601(tm string) (time.Time, error) {
	t, err := time.Parse(ISO8601, tm)
	return t, err
}

func ParseIso8601Micro(tm string) (time.Time, error) {
	t, err := time.Parse(ISO8601Micro, tm)
	return t, err
}

func ParseAuthBearerAccessToken(header http.Header) string {
	return parseAuthHeader(header, AuthTypeBearer)
}

func parseAuthHeader(header http.Header, name string) string {
	var val string
	str := strings.TrimSpace(header.Get("Authorization"))
	arr := strings.SplitN(str, " ", 2)
	for k, v := range arr {
		arr[k] = strings.TrimSpace(v)
	}
	if len(arr) == 2 && arr[0] == name {
		val = arr[1]
	}

	return val
}
