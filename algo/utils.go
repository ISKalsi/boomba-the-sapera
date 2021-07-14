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

var directionToIndex = map[models.Coord]int{
	{0, +1}: UP,
	{0, -1}: DOWN,
	{-1, 0}: LEFT,
	{+1, 0}: RIGHT,
}

var indexToDirection = map[int]models.Coord{
	UP:    {0, +1},
	DOWN:  {0, -1},
	LEFT:  {-1, 0},
	RIGHT: {+1, 0},
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

func insert(a []models.Coord, index int, value models.Coord) []models.Coord {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func reverse(s []models.Coord) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (a *Algorithm) initGrid() grid.Grid {
	snakes := make([]grid.ObstacleProvider, len(a.board.Snakes))
	for i, snake := range a.board.Snakes {
		snakes[i] = snake
	}

	return grid.WithObstacles(a.board.Width, a.board.Height, snakes)
}

func (a *Algorithm) isOkAndNotVisited(cell *cell.Cell) bool {
	return !coord.IsOutside(cell, a.board.Width, a.board.Height) &&
		!cell.IsVisited && !cell.IsBlocked
}
