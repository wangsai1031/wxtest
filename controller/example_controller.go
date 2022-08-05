package controller

import (
	"context"

	"weixin/common/handlers/log"
	"weixin/common/util"

	"team.wphr.vip/technology-group/infrastructure/trace"
	"weixin/idl/proto"
	exampleLogic "weixin/logic"
)

type ExampleController struct {
}

// Ping Ping
func (c ExampleController) Ping(ctx context.Context, req *proto.PingReq) (*proto.PingRsp, error) {
	resp, err := exampleLogic.ExampleLogicInstance.Ping(ctx, req.Name)

	if err != nil {
		log.Trace.Errorf(ctx, trace.DLTagUndefined, "msg=ping failed||errno=%d||err=%+v", util.GetErrno(err), err)
		return &proto.PingRsp{
			Errno:  util.GenOuterErrno(err),
			Errmsg: util.GenOuterErrmsg(err),
		}, nil
	}

	return &proto.PingRsp{
		Errno:  util.GenOuterErrno(err),
		Errmsg: util.GenOuterErrmsg(err),
		Data:   resp,
	}, nil
}
