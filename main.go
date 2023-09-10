package main

import (
	"ysgame/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	routers.RegisterRouters(g)
	g.Run(":8080")
}
