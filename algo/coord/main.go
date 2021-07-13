package coord

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"math"
)

type PositionProvider interface {
	GetX() int
	GetY() int
}

func CalculateHeuristics(source PositionProvider, destination PositionProvider) float64 {
	dx := source.GetX() - destination.GetX()
	dy := source.GetY() - destination.GetY()
	return math.Abs(float64(dx)) + math.Abs(float64(dy))
}

func IsOutside(c PositionProvider, w int, h int) bool {
	if c == nil {
		return true
	} else {
		return c.GetX() < 0 || c.GetY() < 0 || c.GetX() >= w || c.GetY() >= h
	}
}

func Sum(a PositionProvider, b PositionProvider) models.Coord {
	return models.Coord{
		X: a.GetX() + b.GetX(),
		Y: a.GetY() + b.GetY(),
	}
}

func Diff(a PositionProvider, b PositionProvider) models.Coord {
	return models.Coord{
		X: a.GetX() - b.GetX(),
		Y: a.GetY() - b.GetY(),
	}
}
