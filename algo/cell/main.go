package cell

import "github.com/ISKalsi/boomba-the-sapera/models"

const (
	WeightNormal = 1.0
	WeightHazard = 15.0
)

type Cell struct {
	models.Coord
	ParentCoord     models.Coord
	F, G, H         float64
	Weight          float64
	IsBlocked       bool
	ShouldBeBlocked bool
	IsVisited       bool
}

func (c Cell) GetPriority() float64 {
	return c.F
}

func New(coord models.Coord) *Cell {
	return &Cell{
		Coord:           coord,
		ParentCoord:     models.Coord{X: -1, Y: -1},
		F:               -1,
		G:               -1,
		H:               -1,
		Weight:          WeightNormal,
		IsBlocked:       false,
		ShouldBeBlocked: false,
		IsVisited:       false,
	}
}

func (c Cell) IsOk() bool {
	return !(c.IsBlocked || c.ShouldBeBlocked)
}

func (c Cell) IsOkAndNotVisited() bool {
	return !c.IsVisited && c.IsOk()
}
