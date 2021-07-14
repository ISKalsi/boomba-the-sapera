package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

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
