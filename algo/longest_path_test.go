package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPath(t *testing.T) {
	gr := testdata.ThreeLengthSnakeRequest
	a := Init(gr.Board)
	assert.Equal(t, true, a.longestPath())
}
