package httpserv

import (
	"context"
	"net/http"
	"os"
	"team.wphr.vip/technology-group/infrastructure/trace"
	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"
	"weixin/libs/officialaccount"
)

func Run() {
	http.HandleFunc("/event", officialaccount.ServeWechat)

	addr := conf.Viper.GetString("http.addr")
	log.Trace.Info("wechat server listener at", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Trace.Fatalf(context.Background(), trace.DLTagUndefined, "wechat server Run err %v \n", err)
		os.Exit(1)
	}
}
