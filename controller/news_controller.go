package controller

import (
	"context"
	"weixin/common/util"
	"weixin/idl/proto"
	"weixin/logic"
)

type NewsController struct {
}

// SendNews 群发文章消息
func (c NewsController) SendNews(ctx context.Context, req *proto.SendNewsReq) (*proto.SendNewsResp, error) {
	data, err := logic.NewsLogicInstance.SendNews(ctx, req)

	return &proto.SendNewsResp{
		Errno:  util.GenOuterErrno(err),
		Errmsg: util.GenOuterErrmsg(err),
		Data:   data,
	}, nil
}
