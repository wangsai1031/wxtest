package officialaccount

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/freepublish"
	"sync"
	"time"
	"weixin/common/handlers/log"
	"weixin/common/util"
)

type PublishStatusCheckEvent struct {
	publishID int64
	//事件尝试次数
	Attempts int64
	//事件时间
	T time.Time
}

// 监控微信文章发布状态
var publishStatusCheckChan chan *PublishStatusCheckEvent
var onceE sync.Once

func init() {
	onceE.Do(func() {
		publishStatusCheckChan = make(chan *PublishStatusCheckEvent, 100)
	})
}

func TriggerPublishStatusCheckEvent(publishID int64) {
	publishStatusCheckChan <- &PublishStatusCheckEvent{publishID: publishID, T: time.Now()}
}

// 微信任务监控
func TaskRun() {
	log.Trace.Info("weixin event task starting...")
	// set ticker
	ticker := time.NewTicker(time.Duration(10) * time.Second)
	defer ticker.Stop()

	for {
		select {
		// 监控微信文章发布状态
		case e := <-publishStatusCheckChan:
			e.bootstrap(func() error {
				// 通过 InstanceId 查询 Server 及 附加 信息，填充到 v1.Instance 结构体
				publishStatus, err := e.publishStatusCheck()
				if err != nil {
					return err
				}

				// todo 记录发布状态
				_ = publishStatus

				return nil
			})
		case t := <-ticker.C:
			if len(publishStatusCheckChan) == 0 {
				// todo 从数据表中查出 发布中 状态的发布任务，添加到监测队列中
			}
			fmt.Println(t.Format("01-02 15:04:05") + " 微信任务监控中")
		}
	}
}

func (e *PublishStatusCheckEvent) bootstrap(f func() error) {
	if err := f(); err != nil {
		log.Trace.Error("publishID:", e.publishID, "attempts:", e.Attempts, " failure!!! ", err)

		// 如果失败，则重试三次
		if e.Attempts < 3 {
			go util.SafeGo(func() {
				time.Sleep(time.Duration(3) * time.Second)
				e.Attempts++
				publishStatusCheckChan <- e
			})
		}

	} else {
		log.Trace.Info("publishID:", e.publishID, " attempts:", e.Attempts, " done!!!")
	}

}

func (e *PublishStatusCheckEvent) publishStatusCheck() (publishStatus freepublish.PublishStatusList, err error) {
	publishStatus, err = PublishStatus(e.publishID)
	if err != nil {
		log.Trace.Error("PublishStatus() error = ", err)
		return
	}

	// 发布中，1秒后放回队列
	if publishStatus.PublishStatus == freepublish.PublishStatusPublishing {
		util.SafeGo(func() {
			time.Sleep(time.Duration(1) * time.Second)
			publishStatusCheckChan <- e
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
