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

func logToFile() (func(*os.File), *os.File) {
	f, err := os.OpenFile("logs/json.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	deferFunc := func(f *os.File) {
		e := f.Close()
		if e != nil {
			log.Fatalf("error closing file: %v", e)
		}
	}

	log.SetOutput(f)
	return deferFunc, f
}

func main() {
	router := SetupRoutes()
	port := getPort()

	gin.New()

	fn, file := logToFile()
	defer fn(file)

	if e := router.Run(":" + port); e != nil {
		log.Fatal(e)
	}
}
