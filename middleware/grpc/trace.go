package grpcmiddleware

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"weixin/common/handlers/log"
	"weixin/common/server/httpserv"
	"weixin/common/util"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"team.wphr.vip/technology-group/infrastructure/trace"
)

var (
	enableGrpcReqInLog  = false
	enableGrpcReqOutLog = false
)

func TraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		requestTime := time.Now()
		md, _ := metadata.FromIncomingContext(ctx)

		// set trace to ctx
		xdTrace := new(trace.Trace)
		xdTrace.TraceId = getFromMD(md, httpserv.HTTPTraceIdTag)
		xdTrace.SpanId = getFromMD(md, httpserv.HTTPSpanIdTag)
		xdTrace.Method = getFromMD(md, httpserv.HTTPMethodTag)
		xdTrace.URL = getFromMD(md, httpserv.HTTPURLag)
		xdTrace.Params = getFromMD(md, httpserv.HTTPParamsTag)
		xdTrace.Host = getFromMD(md, httpserv.HTTPHostTag)
		xdTrace.From = getFromMD(md, httpserv.HTTPFromTag)
		xdTrace.FormatString = genFormatString(xdTrace)
		ctx = trace.SetCtxTrace(ctx, xdTrace)

		//ctx = util.SetUserInfoFromStr(ctx, getFromMD(md, httpserv.UserInfoTag))
		ctx = util.SetHttpHeaderToCtx(ctx, getFromMD(md, httpserv.HTTPHeaderTag))

		// jsonpb, ignore omitempty option
		jsonPb := &runtime.JSONPb{OrigName: true, EmitDefaults: true}

		if enableGrpcReqInLog {
			// request log
			request, _ := jsonPb.Marshal(req)
			log.Trace.Infof(ctx, " _com_grpc_in", "request=%s", string(request))
		}

		// biz handler
		res, err := handler(ctx, req)

		if enableGrpcReqOutLog {

			var response string
			if err != nil {
				s, ok := status.FromError(err)
				if !ok {
					s = status.New(codes.Unknown, err.Error())
				}
				response = "{\"code\":" + strconv.Itoa(int(s.Code())) + ",\"desc\":\"" + s.Message() + "\"}"
			} else {
				b, _ := jsonPb.Marshal(res)
				response = string(b)
			}

			procTime := int64(time.Since(requestTime) / time.Millisecond)
			log.Trace.Infof(ctx, " _com_grpc_out", "response=%s||proc_time=%d", response, procTime)

		}

		return res, err
	}
}

// getFromMD 获取md数据
func getFromMD(md metadata.MD, key string) string {
	values := md.Get(key)
	if len(values) > 0 {
		return values[0]
	}

	return ""
}

// 生成日志格式
func genFormatString(t *trace.Trace) string {
	return fmt.Sprintf("traceid=%s||spanid=%s||method=%s||host=%s||uri=%s||params=%s||from=%s||srcMethod=%s||caller=%s",
		t.TraceId, t.SpanId, t.Method, t.Host, t.URL, t.Params, t.From, t.SrcMethod, t.Caller)
}
