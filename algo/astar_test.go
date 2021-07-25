package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
	"github.com/ISKalsi/boomba-the-sapera/testdata/head_on"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAStarSearch(t *testing.T) {
	gr := testdata.StartGameRequest
	a := Init(gr.Board, gr.You)
	pathFound, _ := a.aStarSearch()
	assert.True(t, pathFound)
	assert.True(t, len(a.solvedPath) != 0)
}

func TestBlockedDestinationAStarSearch(t *testing.T) {
	gr := head_on.LoseCollidingSnakesRequest
	a := Init(gr.Board, gr.You)
	a.findPossibleLosingHeadCollisions(gr.You)

	pathFound, _ := a.aStarSearch()
	assert.False(t, pathFound)
}

func TestHazardBlockInBetweenAStarSearch(t *testing.T) {
	gr := testdata.HazardBlockRequest
	a := Init(gr.Board, gr.You)

	expectedPath := []models.Coord{
		{5, 8},
		{6, 8},
		{7, 8},
		{7, 7},
		{7, 6},
		{7, 5},
		{6, 5},
		{5, 5},
	}

	pathFound, _ := a.aStarSearch()
	assert.True(t, pathFound)
	assert.Equal(t, expectedPath, a.solvedPath)
}
