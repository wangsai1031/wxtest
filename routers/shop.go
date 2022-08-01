package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"strconv"
)

func LoadShop(r *gin.Engine) {
	v1 := r.Group("/shop")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
		v1.GET("/json", resJson)
		v1.GET("/protobuf", protobuf)
	}
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("Login %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("Submit %s\n", name))
}

func resJson(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "msg": "OK"})
}

func protobuf(c *gin.Context) {
	c1 := c.DefaultQuery("c1", "1")
	c2 := c.DefaultQuery("c2", "2")
	c3 := c.DefaultQuery("c3", "3")

	c1i, _ := strconv.Atoi(c1)
	c2i, _ := strconv.Atoi(c2)
	c3i, _ := strconv.Atoi(c3)

	data := protoexample.Test{
		Reps: []int64{int64(c1i), int64(c2i), int64(c3i)},
	}
	c.ProtoBuf(200, &data)
}
