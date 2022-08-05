package officialaccount

import (
	"github.com/silenceper/wechat/v2/officialaccount/broadcast"
	"github.com/silenceper/wechat/v2/officialaccount/freepublish"
	"sync"
	"time"
	"weixin/common/handlers/log"
	"weixin/common/util"
)

type WeixinEvent struct {
	Id int64
	//事件尝试次数
	Attempts int64
	//事件时间
	T time.Time
}

func (e *WeixinEvent) bootstrap(weixinEventChan chan *WeixinEvent, f func() error) {
	if err := f(); err != nil {
		log.Trace.Error("ID:", e.Id, "attempts:", e.Attempts, " failure!!! ", err)

		// 如果失败，则重试三次
		if e.Attempts < 3 {
			go util.SafeGo(func() {
				time.Sleep(time.Duration(10) * time.Second)
				e.Attempts++
				weixinEventChan <- e
			})
		}

	} else {
		log.Trace.Info("ID:", e.Id, " attempts:", e.Attempts, " done!!!")
	}
}

// 监控群发消息推送状态
var SendMsgStatusCheckChan chan *WeixinEvent

// 监控微信文章发布状态
var PublishStatusCheckChan chan *WeixinEvent
var onceE sync.Once

func init() {
	onceE.Do(func() {
		SendMsgStatusCheckChan = make(chan *WeixinEvent, 100)
		PublishStatusCheckChan = make(chan *WeixinEvent, 100)
	})
}

func TriggerSendMsgStatusCheckEvent(msgID int64) {
	SendMsgStatusCheckChan <- &WeixinEvent{Id: msgID, T: time.Now()}
}

func TriggerPublishStatusCheckEvent(publishID int64) {
	PublishStatusCheckChan <- &WeixinEvent{Id: publishID, T: time.Now()}
}

// 微信任务监控
func TaskRun() {
	log.Trace.Info("weixin event task starting...")
	// set ticker
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		// 监控微信图文消息推送状态
		case e := <-SendMsgStatusCheckChan:
			e.bootstrap(SendMsgStatusCheckChan, func() error {
				// 通过 InstanceId 查询 Server 及 附加 信息，填充到 v1.Instance 结构体
				SendStatus, err := SendMsgStatusCheck(e)
				if err != nil {
					return err
				}

				// todo 记录发布状态
				_ = SendStatus

				return nil
			})
		case e := <-PublishStatusCheckChan:
			e.bootstrap(PublishStatusCheckChan, func() error {
				// 通过 InstanceId 查询 Server 及 附加 信息，填充到 v1.Instance 结构体
				publishStatus, err := PublishStatusCheck(e)
				if err != nil {
					return err
				}

				// todo 记录发布状态
				_ = publishStatus

				return nil
			})
		case t := <-ticker.C:
			if len(SendMsgStatusCheckChan) == 0 {
				// SEND_SUCCESS表示发送成功，SENDING表示发送中，SEND_FAIL表示发送失败，DELETE表示已删除
				// todo 从数据表中查出 发送中 状态的消息推送任务，添加到监测队列中
			}
			if len(PublishStatusCheckChan) == 0 {
				// todo 从数据表中查出 发布中 状态的发布任务，添加到监测队列中
			}
			log.Trace.Info("微信任务监控中 ", t.Format("01-02 15:04:05"))
		}
	}
}

// 监控微信文章发布状态
func PublishStatusCheck(e *WeixinEvent) (publishStatus freepublish.PublishStatusList, err error) {
	publishStatus, err = PublishStatus(e.Id)
	if err != nil {
		log.Trace.Error("PublishStatus() error = ", err)
		return
	}

	// 发布中，3秒后放回队列
	if publishStatus.PublishStatus == freepublish.PublishStatusPublishing {
		util.SafeGo(func() {
			time.Sleep(time.Duration(3) * time.Second)
			PublishStatusCheckChan <- e
		})

		return
	}

	if publishStatus.PublishStatus == freepublish.PublishStatusSuccess {
		log.Trace.Info("PublishStatus() 发布成功 = ", publishStatus)

		return
	}

	log.Trace.Info("PublishStatus() 发布异常 = ", publishStatus)
	return
}

// 监控群发消息推送状态
func SendMsgStatusCheck(e *WeixinEvent) (sendStatus *broadcast.Result, err error) {
	sendStatus, err = GetMassStatus(e.Id)
	if err != nil {
		log.Trace.Error("GetMassStatus() error = ", err)
		return
	}

	// 发布中，3秒后放回队列
	if sendStatus.MsgStatus == string(SendMsgStatusSending) {
		util.SafeGo(func() {
			time.Sleep(time.Duration(5) * time.Second)
			SendMsgStatusCheckChan <- e
		})

		return
	}

	if sendStatus.MsgStatus == string(SendMsgStatusSendSuccess) {
		log.Trace.Info("GetMassStatus() 发布成功 = ", sendStatus)

		return
	}

	log.Trace.Info("GetMassStatus() 发布异常 = ", sendStatus)
	return
}
