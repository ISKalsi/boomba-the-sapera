package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

func New(w int, h int) Grid {
	m := Grid{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			c := models.Coord{X: i, Y: j}
			m[c] = cell.New(c)
		}
	}
	return m
}
