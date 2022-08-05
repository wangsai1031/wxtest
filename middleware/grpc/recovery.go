package grpcmiddleware

import (
	"context"
	"runtime/debug"

	"weixin/common/handlers/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"team.wphr.vip/technology-group/infrastructure/trace"
)

func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Trace.Errorf(ctx, string(trace.DLTagUndefined), "PANIC:%s\n%s", r, debug.Stack())
				err = grpc.Errorf(codes.Internal, "500 Server internal error")
				return
			}
		}()

		return handler(ctx, req)
	}
}
