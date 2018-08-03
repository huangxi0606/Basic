package main

import (
	"github.com/gin-gonic/gin"
	"Basic/Middlewares"
	."Basic/Controller"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(Middlewares.CORSMiddleware())//
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hhx",Hhx)
	//r.POST("/v1/email_dispatch/replyEmail",ReplyEmail)
	r.Run(":9005")
}

