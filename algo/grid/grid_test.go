package grid

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
