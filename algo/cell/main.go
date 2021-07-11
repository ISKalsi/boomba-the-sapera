package cell

import "github.com/ISKalsi/boomba-the-sapera/models"

func New() *Cell {
	return &Cell{
		Parent:    models.Coord{X: -1, Y: -1},
		F:         -1,
		G:         -1,
		H:         -1,
		IsBlocked: false,
	}
}
