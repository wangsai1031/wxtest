package httpmiddleware

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"weixin/common/handlers/log"

	"team.wphr.vip/technology-group/infrastructure/trace"
)

type ctxKey struct {
	name string
}

var (
	// RequestTimeKey ...
	RequestTimeKey = ctxKey{"requestTime"}
	// ExtraInfoReqOut ...
	ExtraInfoReqOut = ctxKey{"extraReqOut"}
)

// TraceConfig ...
type TraceConfig struct {
	Log log.CommonLog
	//parseMultiFrom时限制的内存使用大小
	MaxMemory int64
	//限制request in中打印输出的body长度最大值
	MaxBody int64
}

// TraceWithConfig 打印access log, 采用把脉日志格式，request_in
func TraceWithConfig(c TraceConfig) func(next http.Handler) http.Handler {
	if c.MaxMemory == 0 {
		c.MaxMemory = 1 << 20
	}
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			body := "null"
			tracer := trace.New(req)

			req = req.WithContext(context.WithValue(req.Context(), RequestTimeKey, time.Now()))
			req = req.WithContext(trace.SetCtxTrace(req.Context(), tracer))

			switch req.Method {
			case "POST", "PUT", "PATCH":
				if !strings.Contains(strings.ToLower(req.Header.Get("Content-Type")), "multipart/form-data") {
					b, err := ioutil.ReadAll(req.Body)
					if err != nil && c.Log != nil {
						c.Log.Warnf(req.Context(), trace.DLTagUndefined,
							"errmsg=Trace middleware read request body error:%s", err)
					}
					req.Body.Close()
					req.Body = ioutil.NopCloser(bytes.NewReader(b))

					maxBody := c.MaxBody

					if maxBody > 0 {
						if maxBody > int64(len(b)) {
							maxBody = int64(len(b))
						}
						body = strconv.Quote(string(b[:maxBody]))
					} else {
						body = strconv.Quote(string(b))
					}
				}
			}
			c.Log.Infof(req.Context(), trace.DLTagRequestIn,
				"proto=%s||user_agent=%s||content_type=%s||args=%s",
				req.Proto,
				req.UserAgent(),
				req.Header.Get("Content-Type"),
				body)

			rec := &bytes.Buffer{}
			writer := &traceWriter{ctx: req.Context(), log: c.Log, basicWriter: w, rec: rec, code: http.StatusOK}
			next.ServeHTTP(writer, req)
			traceRequestOut(writer, rec)
		})
	}
}

func traceRequestOut(b *traceWriter, rec *bytes.Buffer) {
	var (
		output string
		format string
	)
	requestTime, ok := b.ctx.Value(RequestTimeKey).(time.Time)
	var duration time.Duration
	if ok {
		duration = time.Since(requestTime)
	}

	header := b.Header()
	contentType := strings.ToLower(header.Get("Content-Type"))
	ms := int64(duration / time.Millisecond)

	if strings.Contains(contentType, "application/json") {
		output = string(rec.Bytes())
	} else {
		output = ""
	}

	format = "status=%d||response=%s||proc_time=%v"

	if b.log != nil {
		b.log.Infof(b.ctx, trace.DLTagRequestOut,
			format,
			b.code, output, ms)
	}
}
