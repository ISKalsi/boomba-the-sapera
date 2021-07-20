package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/testdata/collide_in_itself"
	"github.com/ISKalsi/boomba-the-sapera/testdata/out_of_bounds"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextMove_OutOfBoundsEdgeCase1(t *testing.T) {
	gr := out_of_bounds.EdgeCaseRequest1
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	expectedMove := parseMoveDirectionToString(RIGHT)
	assert.Equal(t, expectedMove, nextMove)
}

func TestNextMove_OutOfBoundsEdgeCase2(t *testing.T) {
	gr := out_of_bounds.EdgeCaseRequest2
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	notExpectedMove := parseMoveDirectionToString(UP)
	assert.NotEqual(t, notExpectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase1(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest1
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	notExpectedMove := parseMoveDirectionToString(UP)
	assert.NotEqual(t, notExpectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase4(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest4
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	notExpectedMove := parseMoveDirectionToString(DOWN)
	assert.NotEqual(t, notExpectedMove, nextMove)
}
