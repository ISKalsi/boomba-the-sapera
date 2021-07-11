package cell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCellPriority(t *testing.T) {
	priority := 10.0
	c := New()
	c.F = priority
	assert.Equal(t, priority, c.GetPriority())
}
