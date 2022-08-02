package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"weixin/common/handlers/conf"
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
}

func main() {
	r := gin.Default()

	routers.LoadEvent(r)

	addr := conf.Viper.GetString("http.addr")

	r.Run(addr)
}
