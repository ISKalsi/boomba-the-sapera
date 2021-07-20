package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefault(t *testing.T) {
	w := 10
	h := 12

	g := Default(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c := models.Coord{X: i, Y: j}
			assert.NotNil(t, g[c])
		}
	}
}

type MockObstacleProviderImplementation struct {
	obstacles []models.Coord
}

func (m MockObstacleProviderImplementation) GetBlockedCoords() []models.Coord {
	return m.obstacles
}

type MockProbableObstacleProviderImplementation struct {
	obstacles []models.Coord
}

func (m MockProbableObstacleProviderImplementation) GetProbablyBlockedCoords() []models.Coord {
	return m.obstacles
}

func search(s []models.Coord, v models.Coord) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}

	return false
}

func TestWithObstacles(t *testing.T) {
	op := MockObstacleProviderImplementation{
		obstacles: []models.Coord{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		},
	}

	pop := MockProbableObstacleProviderImplementation{
		obstacles: []models.Coord{
			{2, 2},
			{3, 3},
		},
	}

	w := 4
	h := 5
	g := WithObstacles(w, h, []ObstacleProvider{op}, []PotentialObstacleProvider{pop})

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c := models.Coord{X: i, Y: j}

			if isActuallyObstacle := search(op.obstacles, c); isActuallyObstacle {
				assert.True(t, g[c].IsBlocked)
				assert.False(t, g[c].ShouldBeBlocked)
			} else if isProbablyObstacle := search(pop.obstacles, c); isProbablyObstacle {
				assert.True(t, g[c].ShouldBeBlocked)
			} else {
				assert.False(t, g[c].IsBlocked)
				assert.False(t, g[c].ShouldBeBlocked)
			}
		}
	}
}

func TestMoveBigVirtualSnakeAlongPath(t *testing.T) {
	/*
		0 -> HEAD
		O -> BODY
		X -> DESTINATION

		  BEFORE		 AFTER
		. X . . .	   . 0 O O O
		. . O O 0	   . . . . O
		. . O . .	   . . . . .
		. . O . .	   . . . . .
		. . . . .	   . . . . .
	*/
	g := Default(5, 5)
	b := []models.Coord{
		{4, 3},
		{3, 3},
		{2, 3},
		{2, 2},
		{2, 1},
	}
	expectedB := []models.Coord{
		{1, 4},
		{2, 4},
		{3, 4},
		{4, 4},
		{4, 3},
	}
	p := []models.Coord{
		{4, 4},
		{3, 4},
		{2, 4},
		{1, 4},
	}

	newB := g.MoveVirtualSnakeAlongPath(b, p)

	assert.Equal(t, expectedB, newB)

	diff := len(b) - len(p)
	for _, coord := range b[diff:] {
		assert.False(t, g[coord].IsBlocked)
	}

	for _, coord := range expectedB {
		assert.True(t, g[coord].IsBlocked)
	}
}

func TestMoveSmallVirtualSnakeAlongPath(t *testing.T) {
	/*
		0 -> HEAD
		O -> BODY
		X -> DESTINATION

		  BEFORE		 AFTER
		. X . . .	   . 0 O O .
		. . O O 0	   . . . . .
		. . . . .	   . . . . .
		. . . . .	   . . . . .
		. . . . .	   . . . . .
	*/
	g := Default(5, 5)
	b := []models.Coord{
		{4, 3},
		{3, 3},
		{2, 3},
	}
	expectedB := []models.Coord{
		{1, 4},
		{2, 4},
		{3, 4},
	}
	p := []models.Coord{
		{4, 4},
		{3, 4},
		{2, 4},
		{1, 4},
	}

	newB := g.MoveVirtualSnakeAlongPath(b, p)

	assert.Equal(t, expectedB, newB)

	for _, coord := range b {
		assert.False(t, g[coord].IsBlocked)
	}

	for _, coord := range expectedB {
		assert.True(t, g[coord].IsBlocked)
	}
}
