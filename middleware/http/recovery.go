package httpmiddleware

import (
	"net/http"
	"runtime/debug"
	"weixin/common/handlers/log"

	"team.wphr.vip/technology-group/infrastructure/trace"
)

// RecoveryConfig ...
type RecoveryConfig struct {
	Log log.CommonLog
}

// RecoveryWithConfig ...
func RecoveryWithConfig(c RecoveryConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					http.Error(w, "500 Server internal error", http.StatusInternalServerError)

					if c.Log != nil {
						c.Log.Errorf(req.Context(), string(trace.DLTagUndefined), "PANIC:%s\n%s", err, debug.Stack())
					}
				}
			}()
			next.ServeHTTP(w, req)
		})
	}
}
