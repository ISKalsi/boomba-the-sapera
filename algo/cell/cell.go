package cell

import "github.com/ISKalsi/boomba-the-sapera/models"

type Cell struct {
	models.Coord
	ParentCoord models.Coord
	F, G, H     float64
	IsBlocked   bool
	IsVisited   bool
}

func (c Cell) GetPriority() float64 {
	return c.F
}
