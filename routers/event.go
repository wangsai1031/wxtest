package routers

import (
	"github.com/gin-gonic/gin"
)

func LoadEvent(r *gin.Engine) {
	r.GET("/event", event)
}

func event(c *gin.Context) {
	echostr := c.DefaultQuery("echostr", "")
	c.String(200, echostr)
}
