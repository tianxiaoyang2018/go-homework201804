package contextlog

import (
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
)

const (
	LogTypeContext = "context"
	LogTypeSpan    = "span"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// NewLog 是初始化一个上下文日志的结构
func NewLog(traceId string) *Log {
	return &Log{
		Type:    LogTypeContext,
		TraceID: traceId,
		Ext:     make(map[string]interface{}),
	}
}

var (
	hostName string
)

func init() {
	hostName, _ = os.Hostname()
}

// NewLogWithType ： 根据指定类型（span还是context）初始化一个上下文日志
func NewLogWithType(logType string, traceId string) *Log {

	return &Log{
		Type:    logType,
		TraceID: traceId,
		Ext:     make(map[string]interface{}),
		//Time:        time.Now(),
		//ServiceHost: hostName,
	}
}

// Log 上下文日志结构
type Log struct {
	Type    string
	TraceID string

	// 扩展字段，可以加入各种需要的字段
	Ext map[string]interface{} `json:"Ext,omitempty"`

	TraceServices []string `json:"TraceServices,omitempty"` // nginx;gateway;user;device;user

	ServiceName    string   `json:"ServiceName,omitempty"`
	ServiceVersion string   `json:"ServiceVer,omitempty"`
	ServiceTags    []string `json:"ServiceTags,omitempty"`
	ServiceHost    string   `json:"Host,omitempty"`
	LogLevel       string   `json:"Level,omitempty"`
	//Time           time.Time `json:"Time"`

	ServerName string `json:"ServerName,omitempty"`
}

func (self *Log) SetExt(key string, value interface{}) {
	if self == nil {
		return
	}
	if self.Ext == nil {
		self.Ext = make(map[string]interface{})
	}
	self.Ext[key] = value
}

func (self *Log) SetMessage(msg interface{}) {
	if self == nil {
		return
	}

	self.SetExt("msg", msg)
}

func (self *Log) SetDebug(db interface{}) {
	if self == nil {
		return
	}

	if es, ok := db.(error); ok {
		self.SetExt("debug", es.Error())
	} else {
		self.SetExt("debug", db)
	}
}

func (self *Log) SetInfo(info interface{}) {
	if self == nil {
		return
	}

	if es, ok := info.(error); ok {
		self.SetExt("info", es.Error())
	} else {
		self.SetExt("info", info)
	}
}

func (self *Log) ToJson() string {
	dd, _ := json.Marshal(self)
	return string(dd)
}

func (self *Log) SetError(err interface{}) {
	if self == nil {
		return
	}

	if es, ok := err.(error); ok {
		self.SetExt("error", es.Error())
	} else {
		self.SetExt("error", err)
	}
}

func (self *Log) SetWarning(warn interface{}) {
	if self == nil {
		return
	}

	if es, ok := warn.(error); ok {
		self.SetExt("warning", es.Error())
	} else {
		self.SetExt("warning", warn)
	}
}

func (self *Log) SetNotice(notice interface{}) {
	if self == nil {
		return
	}

	if es, ok := notice.(error); ok {
		self.SetExt("notice", es.Error())
	} else {
		self.SetExt("notice", notice)
	}
}

func (self *Log) SetAlert(alert interface{}) {
	if self == nil {
		return
	}

	if es, ok := alert.(error); ok {
		self.SetExt("alert", es.Error())
	} else {
		self.SetExt("alert", alert)
	}
}

//   http log

type HttpLog struct {
	StartTime time.Time     `json:"StartTime,omitempty"` // 服务接受到的请求开始时间
	Duration  time.Duration `json:"Duration,omitempty"`  // 服务处理时间,

	Req  *HttpRequest  `json:"Req,omitempty"`
	Resp *HttpResponse `json:"Resp,omitempty"`
}

/*
logData := FromContextcontex)
logData.SetMessage("hello")
log.Warning("%+v", logData)
*/
type HttpRequest struct {
	//Method string `json:"Method,omitempty"`
	//Resource string `json:"Res,omitempty"`

	//ContentType    string  `json:"CT,omitempty"`
	//AcceptLanguage string  `json:"Lang,omitempty"`
	//UserAgent      string  `json:"UA,omitempty"`
	//Geolocation    *GeoUri `json:"Geo,omitempty"`

	//Client *ClientInfo `json:"Client,omitempty"`

	//QueryString *string     `json:"Query,omitempty"`
	//Pagination  *Pagination `json:"Pag,omitempty"`
	//Parameters  *Parameters `json:"Params,omitempty"`

	// 扩展字段，可以加入需要打印的结构，如：http 响应body
	Body interface{}
	Ext  map[string]interface{}
}

type GeoUri struct {
	Latitude  float64
	Longitude float64
	CordC     *float64 `json:"CordC,omitempty"`

	ss     string
	parsed string
}

type ClientInfo struct {
	IP     string
	OS     string
	AppVer string
}

type Pagination struct {
	Since  string
	Util   string
	Offset int
	Limit  int
}

type Parameters struct {
	Search string
	Filter string
	Query  string
	Sort   string
}

type HttpResponse struct {
	Code int
	Body interface{}
	Ext  map[string]interface{}
}
