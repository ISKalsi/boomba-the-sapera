package algorithm

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNextMove(t *testing.T) {
	tests := []struct {
		lastMove           int
		restrictedNextMove int
	}{
		{0, 2},
		{1, 3},
		{2, 0},
		{3, 1},
	}

	for _, test := range tests {
		moveString := parseMoveIndexToString(test.lastMove)
		name := "Restricted move for " + strings.ToUpper(moveString)

		t.Run(name, func(t *testing.T) {
			g := Snake{lastMove: test.lastMove}
			nextMove := g.NextMove()
			assert.NotEqual(t, test.restrictedNextMove, nextMove)
		})
	}
}

func TestParseMoveIndexToString(t *testing.T) {
	tests := []struct {
		expectedString string
		moveIndex      int
	}{
		{"up", 0},
		{"right", 1},
		{"down", 2},
		{"left", 3},
	}

	for _, test := range tests {
		t.Run(strings.ToUpper(test.expectedString), func(t *testing.T) {
			m := parseMoveIndexToString(test.moveIndex)
			assert.Equal(t, test.expectedString, m)
		})
	}
}

func TestPanicFromParseMoveIndexToString(t *testing.T) {
	assert.Panics(t, func() {
		_ = parseMoveIndexToString(4)
	})
}
