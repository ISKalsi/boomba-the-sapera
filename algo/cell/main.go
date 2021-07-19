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

func New(coord models.Coord) *Cell {
	return &Cell{
		Coord:       coord,
		ParentCoord: models.Coord{X: -1, Y: -1},
		F:           -1,
		G:           -1,
		H:           -1,
		IsBlocked:   false,
		IsVisited:   false,
	}
}
