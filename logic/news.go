package logic

import (
	"context"
	"weixin/idl/proto"
)

var NewsLogicInstance NewsLogic

type NewsLogic struct {
}

// 群发文章消息
func (l NewsLogic) SendNews(ctx context.Context, req *proto.SendNewsReq) (resp []*proto.SendNewsData, err error) {
	return
}
