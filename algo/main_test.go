package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
	"github.com/ISKalsi/boomba-the-sapera/testdata/head_on"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeadCollisions(t *testing.T) {
	tests := []struct {
		name          string
		gr            models.GameRequest
		coordsToCheck []models.Coord
		areBlocked    bool
	}{
		{
			"Win condition",
			head_on.WinCollidingSnakesRequest,
			[]models.Coord{
				{6, 6},
				{6, 4},
				{5, 5},
			},
			false,
		},
		{
			"Lose condition",
			head_on.LoseCollidingSnakesRequest,
			[]models.Coord{
				{6, 4},
				{4, 4},
				{5, 5},
			},
			true,
		},
		{
			"Equal length condition",
			head_on.EqualLengthCollidingSnakesRequest,
			[]models.Coord{
				{6, 4},
				{4, 4},
				{5, 5},
			},
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gr := test.gr
			a := Init(gr.Board, gr.You)
			assert.NotNil(t, a.headCollisions.coords)

			a.findPossibleLosingHeadCollisions(gr.You)
			g := a.initGrid()

			for _, c := range test.coordsToCheck {
				assert.Equal(t, test.areBlocked, g[c].ShouldBeBlocked)
				assert.False(t, false, g[c].IsBlocked)
			}
		})
	}
}

func TestFindNearestFood(t *testing.T) {
	gr := testdata.StartGameRequest
	a := Init(gr.Board, gr.You)

	expectedNearest := models.Coord{X: 7, Y: 10}
	assert.Equal(t, expectedNearest, a.destination)

	gr = testdata.ThreeLengthSnakeRequest
	a.reset(gr.Board, gr.You)

	expectedNearest = models.Coord{X: 6, Y: 10}
	assert.Equal(t, expectedNearest, a.destination)
}
