package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"math"
)

type CoordProvider interface {
	GetX() int
	GetY() int
}

func CalculateHeuristics(source CoordProvider, destination CoordProvider) float64 {
	dx := source.GetX() - destination.GetX()
	dy := source.GetY() - destination.GetY()
	return math.Abs(float64(dx)) + math.Abs(float64(dy))
}

func IsOutOfGrid(c CoordProvider, w int, h int) bool {
	if c == nil {
		return true
	} else {
		return c.GetX() < 0 || c.GetY() < 0 || c.GetX() >= w || c.GetY() >= h
	}
}

func Sum(a CoordProvider, b CoordProvider) models.Coord {
	return models.Coord{
		X: a.GetX() + b.GetX(),
		Y: a.GetY() + b.GetY(),
	}
}

func Diff(a CoordProvider, b CoordProvider) models.Coord {
	return models.Coord{
		X: a.GetX() - b.GetX(),
		Y: a.GetY() - b.GetY(),
	}
}
