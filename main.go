package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"team.wphr.vip/technology-group/infrastructure/trace"
	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"
	"weixin/common/handlers/mysql"
	"weixin/common/handlers/redis"
	"weixin/common/util"
	"weixin/libs/officialaccount"
	"weixin/routers"
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
	// context init
	ctx := context.Background()
	// 监控微信任务
	go util.SafeGo(officialaccount.TaskRun)

	// gin
	r := gin.Default()
	routers.LoadNotify(r)

	addr := conf.Viper.GetString("http.addr")

	if err := r.Run(addr); err != nil {
		log.Trace.Fatalf(ctx, trace.DLTagUndefined, "gin Run err %v \n", err)
		os.Exit(1)
	}
}
