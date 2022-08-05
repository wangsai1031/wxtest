package officialaccount

import (
	"context"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"net/http"
	"weixin/common/handlers/log"
	"weixin/common/util"
	"weixin/dao"
	message2 "weixin/libs/officialaccount/message"
	"weixin/libs/officialaccount/server"
)

//ServeWechat 处理消息
func ServeWechat(rw http.ResponseWriter, req *http.Request) {
	officialAccount := GetOfficialAccount()
	// 传入request和responseWriter
	srv := GetServer(officialAccount, req, rw)
	srv.SkipValidate(true)
	//设置接收消息的处理方法
	srv.SetMessageHandler(func(mixMessage *message2.MixMessage) *message.Reply {
		// 保存通知信息到数据库
		go util.SafeGo(func() {
			dao.MessageDaoInstance.Save(req.Context(), mixMessage, string(srv.RequestRawXMLMsg))
		})
		return MessageHandler(req.Context(), mixMessage)
	})

	//处理消息接收以及回复
	err := srv.Serve()
	if err != nil {
		log.Trace.Error("Serve Error, err= ", err)
		return
	}
	//发送回复的消息
	err = srv.Send()
	if err != nil {
		log.Trace.Error("Send Error, err= ", err)
		return
	}
}

// GetServer 消息管理：接收事件，被动回复消息管理
func GetServer(officialAccount *officialaccount.OfficialAccount, req *http.Request, writer http.ResponseWriter) *server.Server {
	srv := server.NewServer(officialAccount.GetContext())
	srv.Request = req
	srv.Writer = writer
	return srv
}

func MessageHandler(ctx context.Context, msg *message2.MixMessage) *message.Reply {
	log.Trace.Info("MessageHandler ", msg)

	switch msg.MsgType {

	case message.MsgTypeEvent:
		return EventHandler(ctx, msg)
	case message.MsgTypeText:

	case message.MsgTypeImage:

	}

	//TODO
	//回复消息：演示回复用户发送的消息
	text := message.NewText("Hello " + msg.Content)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

	//article1 := message.NewArticle("测试图文1", "图文描述", "", "")
	//articles := []*message.Article{article1}
	//news := message.NewNews(articles)
	//return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}

	//voice := message.NewVoice(mediaID)
	//return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}

	//
	//video := message.NewVideo(mediaID, "标题", "描述")
	//return &message.Reply{MsgType: message.MsgTypeVideo, MsgData: video}

	//music := message.NewMusic("标题", "描述", "音乐链接", "HQMusicUrl", "缩略图的媒体id")
	//return &message.Reply{MsgType: message.MsgTypeMusic, MsgData: music}

	//多客服消息转发
	//transferCustomer := message.NewTransferCustomer("")
	//return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: transferCustomer}
}

func EventHandler(ctx context.Context, msg *message2.MixMessage) (reply *message.Reply) {
	switch msg.Event {
	case message.EventPublishJobFinish: // 群发消息推送通知
		//todo 变更群发消息状态，记录发送结果
	case message.EventMassSendJobFinish: // 文章发布任务完成通知
		//todo 变更文章发布状态，记录发布结果
	}

	return
}
