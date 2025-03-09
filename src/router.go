package src

import "github.com/gin-gonic/gin"

func LoadRouter(engine *gin.Engine) *gin.Engine {

	defaultGroup := engine.Group("/")

	defaultGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return engine
}
