package models

import (
	"math"
	"strconv"
)

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Coord) Str() string {
	return "{" + strconv.Itoa(c.X) + ", " + strconv.Itoa(c.Y) + "}"
}

func (c Coord) CalculateHeuristics(destination Coord) float64 {
	dx := c.X - destination.X
	dy := c.Y - destination.Y
	return math.Abs(float64(dx)) + math.Abs(float64(dy))
}

func (c Coord) IsOutside(w int, h int) bool {
	return c.X < 0 || c.Y < 0 || c.X >= w || c.Y >= h
}

func (c Coord) IsAtEdge(w int, h int) bool {
	return c.X == 0 || c.X == w-1 || c.Y == 0 || c.Y == h-1
}

func (c Coord) Sum(c2 Coord) Coord {
	return Coord{
		X: c.X + c2.X,
		Y: c.Y + c2.Y,
	}
}

func (c Coord) Diff(c2 Coord) Coord {
	return Coord{
		X: c.X - c2.X,
		Y: c.Y - c2.Y,
	}
}

func (c Coord) FindNearestCoordsFrom(coords []Coord) []Coord {
	minH := c.CalculateHeuristics(coords[0])
	distanceSlice := make([]float64, len(coords))
	nearestCoords := make([]Coord, 0)

	for i := range coords {
		h := c.CalculateHeuristics(coords[i])
		distanceSlice[i] = h
		if h < minH {
			minH = h
		}
	}

	for i := range distanceSlice {
		if distanceSlice[i] == minH {
			nearestCoords = append(nearestCoords, coords[i])
		}
	}

	return nearestCoords
}
