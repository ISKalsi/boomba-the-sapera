package cell

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCellPriority(t *testing.T) {
	priority := 10.0
	c := New(models.Coord{})
	c.F = priority
	assert.Equal(t, priority, c.GetPriority())
}
