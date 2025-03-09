package main

import (
	"demo-go/src"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	src.LoadRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
