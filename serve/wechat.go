package serve

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"io"
	"net/http"
	"os"
	"weixin/libs/officialaccount"
	"weixin/log"
	"weixin/util"
)

func ServeWechat(rw http.ResponseWriter, req *http.Request) {

	officialAccount := officialaccount.GetOfficialAccount()

	// 传入request和responseWriter
	server := officialAccount.GetServer(req, rw)

	util.RequestInputs(*req)

	//关闭接口验证，则validate结果则一直返回true
	server.SkipValidate(true)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		fmt.Println("SetMessageHandler", text)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil && err != io.EOF {
		log.Error.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func Wechat() {
	http.HandleFunc("/event", ServeWechat)
	fmt.Println("wechat server listener at", ":8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
		os.Exit(0)
	}
}
