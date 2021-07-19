package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

type Grid map[models.Coord]*cell.Cell

type ObstacleProvider interface {
	GetBlockedCoords() []models.Coord
}

func Default(w int, h int) Grid {
	g := Grid{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			c := models.Coord{X: j, Y: i}
			g[c] = cell.New(c)
		}
	}
	return g
}

func WithObstacles(w int, h int, providers []ObstacleProvider) Grid {
	g := Default(w, h)

	for _, p := range providers {
		obstacles := p.GetBlockedCoords()
		for _, c := range obstacles {
			g[c].IsBlocked = true
		}
	}

	return g
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
