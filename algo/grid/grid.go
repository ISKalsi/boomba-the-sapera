package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

type Grid map[models.Coord]*cell.Cell

type ObstacleProvider interface {
	GetBlockedCoords() []models.Coord
}
