package src

import (
	"demo-go/src/demo"

	"github.com/gin-gonic/gin"
)

func LoadRouter(engine *gin.Engine) *gin.Engine {

	defaultGroup := engine.Group("/")

	defaultGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	demoGroup := engine.Group("/demo")
	{
		demoGroup.GET("/helloworld", demo.HelloWorld)
		demoGroup.GET("/ws", demo.HelloWorld)
	}

	return engine
}
