package algorithm

import (
	"log"
	"math/rand"
)

const (
	UP    = "up"
	DOWN  = "down"
	LEFT  = "left"
	RIGHT = "right"

	MOVES = 4
)

type Snake struct {
	lastMove int
}

type NextMoveGetter interface {
	NextMove() string
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

func (s Snake) NextMove() string {
	moveIndex := rand.Intn(MOVES)
	restrictedIndex := (s.lastMove + 2) % 4

	for {
		if moveIndex != restrictedIndex {
			s.lastMove = moveIndex
			return parseMoveIndexToString(moveIndex)
		} else {
			moveIndex = rand.Intn(MOVES)
		}
	}
}
