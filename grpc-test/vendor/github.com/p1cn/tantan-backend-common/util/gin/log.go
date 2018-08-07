package gin

// func NewLogFromGin(ctx *gin.Context, logSections ...log.LogSectionFlag) *log.Log {

// 	sctx := tracing.GetServiceContext(ctx)
// 	var duration time.Duration
// 	sun, ok := ctx.Request.Context().Value(constant.KeyTracingReqStartTime).(int64)
// 	st := time.Unix(0, sun)
// 	if ok && sun > 0 {
// 		now := time.Now()
// 		duration = now.Sub(st)
// 	}

// 	hlog := &log.Log{
// 		Type: log.LogTypeContext,
// 		//ServiceName:    version.ServiceName(),
// 		//ServiceVersion: version.Version(),
// 		TraceID:       sctx.GetTraceID(),
// 		TraceServices: sctx.GetServiceTrace(),
// 		UserID:        GetUserID(ctx),

// 		Http: &log.HttpLog{
// 			StartTime: st,
// 			Duration:  duration,
// 		},
// 	}

// 	for _, sec := range logSections {
// 		switch sec {
// 		case log.LogSectionReqDetail:
// 			if hlog.Http.Req == nil {
// 				hlog.Http.Req = &log.HttpRequest{}
// 			}
// 			req := GetInnerRequest(ctx)
// 			if req != nil {
// 				hlog.Http.Req.Method = ctx.Request.Method
// 				hlog.Http.Req.Resource = ctx.Request.RequestURI
// 				hlog.Http.Req.ContentType = ctx.ContentType()
// 				hlog.Http.Req.AcceptLanguage = req.AcceptLanguage
// 				hlog.Http.Req.UserAgent = req.UserAgent
// 				hlog.Http.Req.Geolocation = req.Geolocation
// 				hlog.Http.Req.Client = &log.ClientInfo{
// 					IP:     GetClientIpAddress(ctx),
// 					OS:     req.ClientOS,
// 					AppVer: req.AppVersion,
// 				}
// 			}

// 			if ctx.Request.URL != nil {
// 				hlog.Http.Req.QueryString = pointer.ToString(ctx.Request.URL.RawQuery)
// 			}
// 		case log.LogSectionReqBody:
// 			if hlog.Http.Req == nil {
// 				hlog.Http.Req = &log.HttpRequest{}
// 			}
// 			bb, _ := ioutil.ReadAll(ctx.Request.Body)
// 			hlog.Http.Req.Body = string(bb)
// 		}
// 	}

// 	return hlog
// }
