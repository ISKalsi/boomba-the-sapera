package mux

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/gin-gonic/gin"
	"log"
)

func decodeGameRequest(context *gin.Context) *models.GameRequest {
	request := &models.GameRequest{}
	if e := context.ShouldBindJSON(request); e != nil {
		log.Fatal(e)
	}

	return request
}
