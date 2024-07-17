package main

import (
	"fmt"
	"ysgame/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	routers.RegisterRouters(g)
	fmt.Println("server start...")
	g.Run(":8000")
}
