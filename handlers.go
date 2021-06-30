package handlers

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleMove(context *gin.Context) {
	request := models.GameRequest{}
	if e := context.ShouldBindJSON(request); e != nil {
		log.Fatal(e)
	}



	println(request)
}
