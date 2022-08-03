package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"weixin/common/handlers/log"
	"weixin/common/util"
	"weixin/libs/officialaccount"
)

func LoadEvent(r *gin.Engine) {
	r.Any("/event", ServeWechat)
}

//ServeWechat 处理消息
func ServeWechat(c *gin.Context) {
	requestInput, _ := util.GinRequestInputs(c)
	log.Trace.Info("ServeWechat", requestInput)

	officialAccount := officialaccount.GetOfficialAccount()
	// 传入request和responseWriter
	server := officialAccount.GetServer(c.Request, c.Writer)
	server.SkipValidate(true)
	//设置接收消息的处理方法
	server.SetMessageHandler(MessageHandler)

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		log.Trace.Error("Serve Error, err= ", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		log.Trace.Error("Send Error, err= ", err)
		return
	}
}

func MessageHandler(msg *message.MixMessage) *message.Reply {
	log.Trace.Info("MessageHandler ", msg)

	switch msg.MsgType {

	case message.MsgTypeEvent:
		return EventHandler(msg)
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

func EventHandler(msg *message.MixMessage) (reply *message.Reply) {
	switch msg.Event {
	case message.EventPublishJobFinish: // 群发消息推送通知
		//todo 变更群发消息状态，记录发送结果
	case message.EventMassSendJobFinish: // 文章发布任务完成通知
		//todo 变更文章发布状态，记录发布结果
	}

	return
}
