package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateHeuristics(t *testing.T) {
	start := Coord{X: 1, Y: 3}
	end := Coord{X: 3, Y: 2}

	h := start.CalculateHeuristics(end)

	assert.Equal(t, 3.0, h)
}

func TestIsOutOfGrid(t *testing.T) {
	w := 11
	h := 11

	inside := Coord{X: 1, Y: 1}
	outside := Coord{X: 15, Y: 15}

	assert.False(t, inside.IsOutside(w, h))
	assert.True(t, outside.IsOutside(w, h))
}

func TestSum(t *testing.T) {
	a := Coord{X: 1, Y: 2}
	b := Coord{X: 2, Y: 1}
	c := Coord{X: 3, Y: 3}

	sum := a.Sum(b)

	assert.Equal(t, c, sum)
}

func TestDiff(t *testing.T) {
	a := Coord{X: 1, Y: 2}
	b := Coord{X: 2, Y: 1}
	c := Coord{X: -1, Y: 1}

	diff := a.Diff(b)

	assert.Equal(t, c, diff)
}

func TestFindNearest(t *testing.T) {
	tests := []struct {
		source  Coord
		nearest Coord
		coords  []Coord
	}{
		{
			Coord{3, 5},
			Coord{2, 4},
			[]Coord{
				{1, 0},
				{2, 4},
				{-1, 3},
				{7, 6},
			},
		},
	}
	for _, test := range tests {
		name := "s: " + test.source.Str() + ", d: " + test.nearest.Str()
		t.Run(name, func(t *testing.T) {
			c := test.source.FindNearest(test.coords)
			assert.Equal(t, test.nearest, c)
		})
	}
}
