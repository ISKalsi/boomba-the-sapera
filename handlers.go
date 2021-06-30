package main

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
)

const (
	UP    = "up"
	DOWN  = "down"
	LEFT  = "left"
	RIGHT = "right"

	MOVES = 4
)

type Game struct {
	lastMove int
}

func parseMoveIndexToString(id int) string {
	switch id {
	case 0:
		return UP
	case 1:
		return RIGHT
	case 2:
		return DOWN
	case 3:
		return LEFT
	default:
		log.Panicf("index out of range: \"%v\"\n", id)
		return ""
	}
}

func (g *Game) nextMoveIndex() int {
	moveIndex := rand.Intn(MOVES)
	restrictedIndex := (g.lastMove + 2) % 4

	for {
		if moveIndex != restrictedIndex {
			g.lastMove = moveIndex
			return moveIndex
		} else {
			moveIndex = rand.Intn(MOVES)
		}
	}
}

func (g *Game) decodeGameRequest(context *gin.Context) *models.GameRequest {
	request := &models.GameRequest{}
	if e := context.ShouldBindJSON(request); e != nil {
		log.Fatal(e)
	}

	return request
}

func HandleIndex(context *gin.Context) {
	response := models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "pavandubey",
		Color:      "#F487B6",
		Head:       "shades",
		Tail:       "weight",
	}

	context.IndentedJSON(http.StatusOK, response)
}

func HandleStart(_ *gin.Context) {
	log.Println("GAME STARTED.")
}

func (g *Game) HandleMove(context *gin.Context) {
	_ = g.decodeGameRequest(context)

	response := models.MoveResponse{
		Move: parseMoveIndexToString(g.nextMoveIndex()),
	}

	context.JSON(http.StatusOK, &response)
}

func HandleEnd(_ *gin.Context) {
	log.Println("GAME ENDED.")
}
