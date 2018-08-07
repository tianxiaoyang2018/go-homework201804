package gin

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/p1cn/tantan-backend-common/util/constant"
)

var (
	DefaultExpire time.Time
)

const (
	DefaultCacheControl              = "private, no-cache, no-store, must-revalidate"
	DefaultAccessControlAllowOrigin  = "*"
	DefaultAccessControlAllowHeaders = "Accept-language, Accept, Authorization, Content-Type, Geolocation, Content-Length"
	DefaultAccessControlMaxAge       = int64(86400)

	ContentTypeJson = "application/json"
)

func init() {
	DefaultExpire, _ = time.Parse(http.TimeFormat, "Wed, 31 Dec 2014 23:59:59 GMT")
}

func GetClientIpAddress(ctx *gin.Context) string {
	if ctx.Request == nil {
		return ""
	}
	var ipAddress string
	addrs := strings.Split(ctx.Request.Header.Get("X-Forwarded-For"), ",")
	if len(addrs) > 0 {
		ipAddress = addrs[0]
	} else {
		ipAddress = ctx.Request.Header.Get("X-Real-IP")
	}
	return strings.Trim(ipAddress, " ")
}

func GetUserID(ctx *gin.Context) string {
	if ctx.Request == nil {
		return ""
	}
	s, _ := ctx.Request.Context().Value(constant.HeaderUserID).(string)
	return s
}

func SetCacheControl(c *gin.Context, v string) {
	c.Header("Cache-Control", v)
}

func SetExpires(c *gin.Context, t time.Time) {
	c.Header("Expires", t.Format(http.TimeFormat))
}

func SetTantanExpire(c *gin.Context) {
	c.Header("Expires", "Wed, 31 Dec 2014 23:59:59 GMT")
}

func SetAccessControlAllowOrigin(c *gin.Context, v string) {
	c.Header("Access-Control-Allow-Origin", v)
}

func SetAccessControlAllowHeaders(c *gin.Context, v string) {
	c.Header("Access-Control-Allow-Headers", v)
}

func SetAccessControlMaxAge(c *gin.Context, v int64) {
	c.Header("Access-Control-Max-Age", strconv.FormatInt(v, 10))
}

func GetUserAgent(c *gin.Context) string {
	return c.GetHeader("User-Agent")
}

func SetContentType(c *gin.Context, ht string) {
	c.Header("Content-Type", ht)
}

func GetInnerRequest(c *gin.Context) *Request {
	if c.Request == nil {
		return nil
	}
	req, _ := c.Request.Context().Value(constant.KeyTracingRequest).(*Request)
	return req
}
