package main

import (
	"github.com/ISKalsi/boomba-the-sapera/mux"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	m := mux.New()

	r.GET("/", m.HandleIndex)
	r.POST("/start", m.HandleStart)
	r.POST("/move", m.HandleMove)
	r.POST("/end", m.HandleEnd)

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
	router := SetupRoutes()
	port := getPort()

	if e := router.Run(":" + port); e != nil {
		log.Fatal(e)
	}
}
