package main

import (
	"ysgame/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	routers.RegisterRouters(g)
	g.Run(":9090")
}
