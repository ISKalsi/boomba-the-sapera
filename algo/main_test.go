package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/testdata"
	"github.com/ISKalsi/boomba-the-sapera/testdata/collide_in_itself"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextMove_SnakeAtEdge(t *testing.T) {
	gr := testdata.OutOfBoundsEdgeCaseRequest
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	expectedMove := parseMoveDirectionToString(RIGHT)
	assert.Equal(t, expectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase1(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest1
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	notExpectedMove := parseMoveDirectionToString(UP)
	assert.NotEqual(t, notExpectedMove, nextMove)
}
