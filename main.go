package main

import (
	"github.com/gin-gonic/gin"
	"weixin/routers"
)

func main() {
	r := gin.Default()

	routers.LoadEvent(r)

	r.Run(":8000")
}
