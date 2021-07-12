package algo

import (
	"github.com/stretchr/testify/assert"
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
