package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"team.wphr.vip/technology-group/infrastructure/trace"
	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"
	"weixin/common/handlers/mysql"
	"weixin/common/handlers/redis"
	"weixin/common/server/grpcserv"
	"weixin/common/server/httpserv"
	"weixin/common/util"
	"weixin/libs/officialaccount"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", "./conf/app.dev.toml", "-c set config file path") // default config file is conf/app.toml
	flag.Parse()
	fmt.Printf("confPath is %s\n", confPath)

	conf.InitConf(confPath)

	log.Init()
	mysql.Init()
	redis.Init()
}

func main() {
	// 监控微信任务
	go util.SafeGo(officialaccount.TaskRun)
	go util.SafeGo(httpserv.Run)

	// 启动服务
	if err := grpcserv.Run(); err != nil {
		log.Trace.Fatalf(context.Background(), trace.DLTagUndefined, "grpcserver Run err %v \n", err)
		os.Exit(1)
	}
}
