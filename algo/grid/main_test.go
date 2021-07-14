package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefault(t *testing.T) {
	w := 10
	h := 12

	g := Default(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c := models.Coord{X: i, Y: j}
			assert.NotNil(t, g[c])
		}
	}
}

type MockObstacleProviderImplementation struct {
	obstacles []models.Coord
}

func (m MockObstacleProviderImplementation) GetBlockedCoords() []models.Coord {
	return m.obstacles
}

func search(s []models.Coord, v models.Coord) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}

	return false
}

func TestWithObstacles(t *testing.T) {
	op := MockObstacleProviderImplementation{
		obstacles: []models.Coord{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		},
	}

	w := 4
	h := 5
	g := WithObstacles(w, h, []ObstacleProvider{op})

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c := models.Coord{X: i, Y: j}
			isActuallyObstacle := search(op.obstacles, c)
			if isActuallyObstacle {
				assert.True(t, g[c].IsBlocked)
			} else {
				assert.False(t, g[c].IsBlocked)
			}
		}
	}
}
