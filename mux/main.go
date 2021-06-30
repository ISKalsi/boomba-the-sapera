package mux

import (
	"github.com/ISKalsi/boomba-the-sapera/algorithm"
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Mux struct {
	getter *algorithm.NextMoveGetter
}

func New(nextMoveGetter *algorithm.NextMoveGetter) *Mux {
	return &Mux{getter: nextMoveGetter}
}

func decodeGameRequest(context *gin.Context) *models.GameRequest {
	request := &models.GameRequest{}
	if e := context.ShouldBindJSON(request); e != nil {
		log.Fatal(e)
	}

	return request
}

func (h *Mux) HandleIndex(context *gin.Context) {
	response := models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "pavandubey",
		Color:      "#F487B6",
		Head:       "shades",
		Tail:       "weight",
	}

	context.IndentedJSON(http.StatusOK, response)
}

func (h *Mux) HandleStart(_ *gin.Context) {
	log.Println("GAME STARTED.")
}

func (h *Mux) HandleMove(context *gin.Context) {
	_ = decodeGameRequest(context)

	response := models.MoveResponse{
		Move: (*h.getter).NextMove(),
	}

	context.JSON(http.StatusOK, &response)
}

func (h *Mux) HandleEnd(_ *gin.Context) {
	log.Println("GAME ENDED.")
}
