package routers

import (
	"github.com/gin-gonic/gin"
)

func LoadBlob(r *gin.Engine) {
	v2 := r.Group("/blob")
	{
		v2.GET("/login", login)
		v2.GET("/submit", submit)
	}
}
