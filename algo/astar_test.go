package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAStarSearch(t *testing.T) {
	gr := testdata.StartGameRequest
	a := Init(gr.Board)
	assert.True(t, a.aStarSearch())
	assert.True(t, a.solvedPath.Len() != 0)
}
