package main

import (
	"os"
	"os/signal"
	"syscall"
	"weixin/log"
	"weixin/serve"
)

func main() {
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	go serve.Wechat()

	<-chSig
	log.Info.Println("process exit")
}
