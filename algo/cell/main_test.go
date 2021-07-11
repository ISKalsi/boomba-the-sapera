package cell

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCell(t *testing.T) {
	c := New()

	assert.Equal(t, models.Coord{X: -1, Y: -1}, c.Parent)
	assert.Equal(t, -1.0, c.F)
	assert.Equal(t, -1.0, c.G)
	assert.Equal(t, -1.0, c.H)
	assert.Equal(t, false, c.IsBlocked)
}
