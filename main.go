package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func SetupRoutes(g *Game) *gin.Engine {
	r := gin.Default()

	r.GET("/", HandleIndex)
	r.POST("/start", HandleStart)
	r.POST("/move", g.HandleMove)
	r.POST("/end", HandleEnd)

	return r
}

func getPort() string {
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "8080"
	}
	return port
}

func main() {
	g := Game{lastMove: 0}

	router := SetupRoutes(&g)
	port := getPort()

	if e := router.Run(":" + port); e != nil {
		log.Fatal(e)
	}
}
