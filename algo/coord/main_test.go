package coord

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateHeuristics(t *testing.T) {
	start := models.Coord{X: 1, Y: 3}
	end := models.Coord{X: 3, Y: 2}

	h := CalculateHeuristics(&start, &end)

	assert.Equal(t, 3.0, h)
}

func TestIsOutOfGrid(t *testing.T) {
	w := 11
	h := 11

	inside := models.Coord{X: 1, Y: 1}
	outside := models.Coord{X: 15, Y: 15}

	assert.False(t, IsOutside(&inside, w, h))
	assert.True(t, IsOutside(&outside, w, h))
}

func TestSum(t *testing.T) {
	a := models.Coord{X: 1, Y: 2}
	b := models.Coord{X: 2, Y: 1}
	c := models.Coord{X: 3, Y: 3}

	sum := Sum(&a, &b)

	assert.Equal(t, c, sum)
}

func TestDiff(t *testing.T) {
	a := models.Coord{X: 1, Y: 2}
	b := models.Coord{X: 2, Y: 1}
	c := models.Coord{X: -1, Y: 1}

	diff := Diff(&a, &b)

	assert.Equal(t, c, diff)
}
