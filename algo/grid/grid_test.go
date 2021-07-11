package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCellBlocked(t *testing.T) {
	coord := models.Coord{X: 1, Y: 2}
	g := New(5, 5)

	g[coord].IsBlocked = true
	assert.True(t, g.IsBlocked(&coord))
}
