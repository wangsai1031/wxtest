package httpmiddleware

import (
	"bytes"
	"context"
	"net/http"

	"weixin/common/handlers/log"

	"team.wphr.vip/technology-group/infrastructure/trace"
)

type basicWriter interface {
	http.ResponseWriter
}

// traceWriter wraps a http.ResponseWriter that implements the minimal
// http.ResponseWriter interface.
type traceWriter struct {
	rec *bytes.Buffer
	ctx context.Context
	log log.CommonLog
	basicWriter
	code int
}

func (b *traceWriter) Write(buf []byte) (int, error) {
	n, err := b.rec.Write(buf)
	if err != nil {
		if b.log != nil {
			b.log.Errorf(b.ctx, trace.DLTagUndefined, "write buffer error || errormsg=%s", err.Error())
		}
		return n, err
	}
	n, err = b.basicWriter.Write(buf)

	return n, err
}

func (b *traceWriter) WriteHeader(code int) {
	b.basicWriter.WriteHeader(code)
	b.code = code
}
