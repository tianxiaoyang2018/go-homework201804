package constant

import (
	"net/http"
	"strings"
)

const (
	HeaderUserID       = "X-Putong-User-Id"
	HeaderAuthType     = "X-Authorization-Type"
	HeaderTestingGroup = "X-Testing-Group"
)

/*
http 头带的信息传递给下游http服务：
	X-B3-TraceId :  Random.Int63()  << 64 | Random.Int63()
	X-B3-SpanId :   Random.Int63()  << 64 | Random.Int63()
	X-TT-STrace : 服务名的数组  例如： "nginx;gateway;tantan-user-restapi"

例子：

// nginx 发往下游的请求http头应该包含
X-B3-TraceId :  5822107fcfd52d29a0f5f3f164fd
X-B3-SpanId : 5822107fcfd52d29a0f5f3f16411
X-TT-STrace : nginx

// gateway服务接受到nginx请求后，再调用下游的http服务再http头里面添加这些trace信息
// http服务接受到请求应该生成自己的span id，再将上游的span id改成parent span id
X-B3-TraceId :  5822107fcfd52d29a0f5f3f164fd
X-B3-ParentSpanId : 5822107fcfd52d29a0f5f3f16411
X-B3-SpanId : 5822107fcfd52d29a0f5f3f16422
X-TT-STrace : "nginx;gateway"

*/

//
const (
	// http 头里面的trace 信息; 作为内部http调用传递用途
	KeyTracingTrace        = "X-B3-TraceId" //  Random.Int63()
	KeyTracingParentSpan   = "X-B3-ParentSpanId"
	KeyTracingSpan         = "X-B3-SpanId" //  Random.Int63()
	KeyTracingServiceTrace = "X-TT-STrace" // 服务名字 以";" 分开, e.g. : nginx;api-gateway;tantan-backend-user;tantan-backend-device

	//
	KeyTracingReqStartTime = "X-TT-Req-Start" // 请求开始时间 :  nanoseconds
	KeyTracingApiId        = "X-TT-API-ID"    // api id
	KeyTracingBaggage      = "X-TT-Baggage"
	KeyTracingDeviceId     = "X-TT-DeviceId"

	KeyTracingDebug   = "X-TT-Debug"
	KeyServiceContext = "service.context"
	KeyTracingRequest = "X-TT-Request"
	KeyTracingExt     = "X-TT-Ext"
	KeyTracingSampled = "X-TT-Sampled"
)

func ParseIpAddress(req *http.Request) (string, error) {
	var ipAddress string
	addrs := strings.Split(req.Header.Get("X-Forwarded-For"), ",")
	if len(addrs) > 0 {
		ipAddress = addrs[0]
	} else {
		ipAddress = req.Header.Get("X-Real-IP")
	}
	ipAddress = strings.Replace(ipAddress, " ", "", -1)
	return ipAddress, nil
}
