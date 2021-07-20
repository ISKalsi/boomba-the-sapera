package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestParseMoveIndexToString(t *testing.T) {
	tests := []struct {
		expectedString string
		moveIndex      int
	}{
		{"up", 0},
		{"down", 1},
		{"left", 2},
		{"right", 3},
	}

	for _, test := range tests {
		t.Run(strings.ToUpper(test.expectedString), func(t *testing.T) {
			m := parseMoveDirectionToString(test.moveIndex)
			assert.Equal(t, test.expectedString, m)
		})
	}
}

func TestPanicFromParseMoveIndexToString(t *testing.T) {
	assert.Panics(t, func() {
		_ = parseMoveDirectionToString(4)
	})
}

func TestDirectionAndIndexMapEquality(t *testing.T) {
	for dir, idx := range directionToIndex {
		assert.Equal(t, indexToDirection[idx], dir)
	}
}

func TestInsert(t *testing.T) {
	a := models.Coord{X: 0, Y: 0}
	b := models.Coord{X: 0, Y: 1}
	c := models.Coord{X: 1, Y: 0}
	d := models.Coord{X: 1, Y: 1}

	tests := []struct {
		index int
	}{
		{0},
		{1},
		{2},
		{3},
	}

	for _, test := range tests {
		t.Run("index"+strconv.Itoa(test.index), func(t *testing.T) {
			cells := []models.Coord{a, b, c}
			cells = insert(cells, test.index, d)

			assert.Equal(t, d, cells[test.index])
		})
	}
}

func TestReverse(t *testing.T) {
	a := models.Coord{X: 0, Y: 0}
	b := models.Coord{X: 0, Y: 1}
	c := models.Coord{X: 1, Y: 0}
	d := models.Coord{X: 1, Y: 1}

	slice := []models.Coord{a, b, c, d}
	reversedSlice := []models.Coord{d, c, b, a}

	reverse(slice)

	assert.Equal(t, reversedSlice, slice)
}

func TestHeadCollisions(t *testing.T) {
	tests := []struct {
		name          string
		gr            models.GameRequest
		coordsToCheck []models.Coord
		areBlocked    bool
	}{
		{
			"Win condition",
			testdata.WinCollidingSnakesRequest,
			[]models.Coord{
				{6, 6},
				{6, 4},
				{5, 5},
			},
			false,
		},
		{
			"Lose condition",
			testdata.LoseCollidingSnakesRequest,
			[]models.Coord{
				{6, 4},
				{4, 4},
				{5, 5},
			},
			true,
		},
		{
			"Equal length condition",
			testdata.EqualLengthCollidingSnakesRequest,
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

			a.reset(gr.Board, gr.You)

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
