package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wxtest/log"
	"wxtest/util"
)

func LoadEvent(r *gin.Engine) {
	r.GET("/event", event)
}

func event(c *gin.Context) {
	inputs, err := util.RequestInputs(c)
	if err != nil {
		log.Error.Println(err.Error())
		c.String(400, err.Error())
		return
	}

	str := fmt.Sprintf("%+v", inputs)

	log.Info.Println(str)
	c.String(200, str)
}
