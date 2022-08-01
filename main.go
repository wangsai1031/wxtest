package main

import (
	"github.com/gin-gonic/gin"
	"wxtest/routers"
)

func main() {
	r := gin.Default()

	routers.LoadEvent(r)

	r.Run(":8000")
}
