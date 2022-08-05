package grpcmiddleware

import (
	"context"

	"weixin/idl/exterr"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func ValidatorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if req4Validate, ok := req.(interface{ Validate() error }); ok {
			validateErr := req4Validate.Validate()
			if validateErr != nil {
				err = grpc.Errorf(codes.Code(exterr.E_PARAM_ERROR), validateErr.Error())
				return
			}
		}

		return handler(ctx, req)
	}
}
