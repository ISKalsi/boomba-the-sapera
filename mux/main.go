package mux

import (
	"github.com/ISKalsi/boomba-the-sapera/algo"
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Mux struct {
	algorithm algo.PathSolver
}

func New() *Mux {
	return &Mux{}
}

func (m *Mux) HandleIndex(context *gin.Context) {
	response := models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "pavandubey",
		Color:      "#F487B6",
		Head:       "shades",
		Tail:       "weight",
	}

	context.IndentedJSON(http.StatusOK, response)
}

func (m *Mux) HandleStart(context *gin.Context) {
	request := decodeGameRequest(context)
	m.algorithm = algo.Init(request.Board, request.You)

	log.Println("GAME STARTED.")
}

func (m *Mux) HandleMove(context *gin.Context) {
	request := decodeGameRequest(context)

	response := models.MoveResponse{
		Move: m.algorithm.NextMove(request),
	}

	context.JSON(http.StatusOK, response)
}

func (m *Mux) HandleEnd(_ *gin.Context) {
	log.Println("GAME ENDED.")
}
