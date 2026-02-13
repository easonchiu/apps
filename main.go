package main

import (
	"log"
	"ysgame/db"
	"ysgame/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB
	if err := db.InitMongoDB(); err != nil {
		log.Printf("Warning: Failed to connect to MongoDB: %v", err)
		log.Println("Application will run without database functionality")
	} else {
		defer db.CloseMongoDB()
	}

	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	routers.RegisterRouters(g)
	g.Run(":9090")
}
