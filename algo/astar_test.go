package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAStarSearch(t *testing.T) {
	gr := testdata.StartGameRequest
	a := Init(gr.Board, gr.You)
	assert.True(t, a.aStarSearch())
	assert.True(t, len(a.solvedPath) != 0)
}

func TestBlockedDestinationAStarSearch(t *testing.T) {
	gr := testdata.LoseCollidingSnakesRequest
	a := Init(gr.Board, gr.You)
	a.reset(gr.Board, gr.You)
	assert.False(t, a.aStarSearch())
}
