package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"log"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

var directions = map[models.Coord]int{
	{0, +1}: UP,
	{0, -1}: DOWN,
	{-1, 0}: LEFT,
	{+1, 0}: RIGHT,
}

func parseMoveIndexToString(id int) string {
	switch id {
	case UP:
		return "up"
	case RIGHT:
		return "right"
	case DOWN:
		return "down"
	case LEFT:
		return "left"
	default:
		log.Panicf("index out of range: \"%v\"\n", id)
		return ""
	}
}

func getRandomMove(msg string) string {
	m := parseMoveIndexToString(UP)
	println(msg + m)
	return m
}
