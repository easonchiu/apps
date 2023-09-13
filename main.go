package main

import (
	"ysgame/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	gin.SetMode(gin.ReleaseMode)
	routers.RegisterRouters(g)
	g.Run(":9090")
}
