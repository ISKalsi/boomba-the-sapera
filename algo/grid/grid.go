package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

type Grid map[models.Coord]*cell.Cell

type ObstacleProvider interface {
	GetBlockedCoords() []models.Coord
}

func reverse(s []models.Coord) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (g Grid) MoveVirtualSnakeAlongPath(body []models.Coord, path []models.Coord) []models.Coord {
	for i := range body {
		g[body[i]].IsBlocked = false
	}

	bodyLen := len(body)
	pathLen := len(path)

	var p []models.Coord

	if bodyLen <= pathLen {
		diff := pathLen - bodyLen
		p = path[diff:]
		reverse(p)
	} else {
		diff := bodyLen - pathLen
		p = path[:]
		reverse(p)
		p = append(p, body[:diff]...)
	}

	for i := range p {
		g[p[i]].IsBlocked = true
	}

	return p
}
