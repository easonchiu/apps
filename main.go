package main

import (
	"ysgame/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	routers.RegisterRouters(g)
	g.Run(":8000")
}
