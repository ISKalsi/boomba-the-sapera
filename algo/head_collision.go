package algo

import "github.com/ISKalsi/boomba-the-sapera/models"

type possibleHeadCollisions struct {
	coords []models.Coord
}

func (phc possibleHeadCollisions) GetProbablyBlockedCoords() []models.Coord {
	return phc.coords
}
