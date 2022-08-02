package util

import (
	"context"
	"runtime/debug"

	"team.wphr.vip/technology-group/infrastructure/trace"
	"weixin/common/handlers/log"
)

func SafeGo(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Trace.Errorf(context.Background(), string(trace.DLTagUndefined), "PANIC:%s\n%s", err, debug.Stack())
		}
	}()

	f()
}
