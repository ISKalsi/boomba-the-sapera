package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/testdata"
	"testing"
)

func TestLongestPath(t *testing.T) {
	gr := testdata.ThreeLengthSnakeRequest
	a := Init(gr.Board)
	_ = a.longestPath()
}
