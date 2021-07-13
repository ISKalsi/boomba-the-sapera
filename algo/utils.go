package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/algo/coord"
	"github.com/ISKalsi/boomba-the-sapera/algo/grid"
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

func parseMoveDirectionToString(id int) string {
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
	m := parseMoveDirectionToString(UP)
	println(msg + m)
	return m
}

func (a *Algorithm) initGrid() grid.Grid {
	g := grid.New(a.board.Width, a.board.Height)

	for _, snake := range a.board.Snakes {
		for _, c := range snake.Body {
			g[c].IsBlocked = true
		}
	}

	return g
}

func (a *Algorithm) isOkAndNotVisited(cell *cell.Cell) bool {
	return !coord.IsOutside(cell, a.board.Width, a.board.Height) &&
		!cell.IsVisited && !cell.IsBlocked
}
