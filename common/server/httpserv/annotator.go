package httpserv

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"team.wphr.vip/technology-group/infrastructure/trace"
)

const (
	HTTPTraceIdTag      = "grpcgateway-http-traceid"
	HTTPSpanIdTag       = "grpcgateway-http-spanid"
	HTTPHintCodeTag     = "grpcgateway-http-hintCode"
	HTTPHintContenttTag = "grpcgateway-http-hintContent"
	HTTPMethodTag       = "grpcgateway-http-method"
	HTTPURLag           = "grpcgateway-http-url"
	HTTPParamsTag       = "grpcgateway-http-params"
	HTTPHostTag         = "grpcgateway-http-host"
	HTTPFromTag         = "grpcgateway-http-from"
	HTTPHeaderTag       = "grpcgateway-http-header"
	UserInfoTag         = "grpcgateway-http-user-info"
)

// AnnotatorHTTPReq http req -> grpc metadata
func AnnotatorHTTPReq(ctx context.Context, req *http.Request) metadata.MD {
	var md metadata.MD

	if TraceStruct, ok := trace.GetCtxTrace(req.Context()); ok && TraceStruct != nil {
		md = metadata.Pairs(
			HTTPTraceIdTag, TraceStruct.TraceId,
			HTTPSpanIdTag, TraceStruct.SpanId,
			HTTPMethodTag, TraceStruct.Method,
			HTTPURLag, TraceStruct.URL,
			HTTPParamsTag, TraceStruct.Params,
			HTTPHostTag, TraceStruct.Host,
			HTTPFromTag, TraceStruct.From,
		)
	}

	//userInfo := util.GetUserInfo(ctx)
	//val, _ := json.Marshal(userInfo)
	//if !userInfo.IsDefault() {
	//	md.Set(UserInfoTag, string(val))
	//}

	return md
}
