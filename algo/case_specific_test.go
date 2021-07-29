package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/testdata/collide_in_itself"
	"github.com/ISKalsi/boomba-the-sapera/testdata/collide_in_snake"
	"github.com/ISKalsi/boomba-the-sapera/testdata/food_related"
	"github.com/ISKalsi/boomba-the-sapera/testdata/hazard_related"
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

func TestNextMove_OutOfBoundsEdgeCase3(t *testing.T) {
	gr := out_of_bounds.EdgeCaseRequest3
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

func TestNextMove_CollideInItselfEdgeCase2(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest2
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	expectedMove := parseMoveDirectionToString(DOWN)
	assert.Equal(t, expectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase3(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest3
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	expectedMove := parseMoveDirectionToString(RIGHT)
	assert.Equal(t, expectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase4(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest4
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	notExpectedMove := parseMoveDirectionToString(DOWN)
	assert.NotEqual(t, notExpectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase5(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest5
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	notExpectedMove := parseMoveDirectionToString(UP)
	assert.NotEqual(t, notExpectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase6(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest6
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	expectedMove := parseMoveDirectionToString(DOWN)
	assert.Equal(t, expectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase7(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest7
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	expectedMove := parseMoveDirectionToString(DOWN)
	assert.Equal(t, expectedMove, nextMove)
}

func TestNextMove_CollideInItselfEdgeCase8(t *testing.T) {
	gr := collide_in_itself.EdgeCaseRequest8
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	expectedMove := parseMoveDirectionToString(RIGHT)
	assert.Equal(t, expectedMove, nextMove)
}

func TestNextMove_CollideInSnakeEdgeCase1(t *testing.T) {
	gr := collide_in_snake.EdgeCaseRequest1
	a := Init(gr.Board, gr.You)

	nextMove := a.NextMove(&gr)
	notExpectedMove := parseMoveDirectionToString(LEFT)
	assert.NotEqual(t, notExpectedMove, nextMove)
}

func TestNoPathToTailAndAvoidHazard(t *testing.T) {
	gr := hazard_related.EdgeCaseRequest1
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(RIGHT)
	assert.Equal(t, expectedNextMove, actualNextMove)
}

func TestNextMove_OutOfHealthFromHazardEdgeCase2(t *testing.T) {
	gr := hazard_related.EdgeCaseRequest2
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(RIGHT)
	assert.Equal(t, expectedNextMove, actualNextMove)
}

func TestNextMove_OutOfHealthFromHazardEdgeCase3(t *testing.T) {
	gr := hazard_related.EdgeCaseRequest3
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(RIGHT)
	assert.Equal(t, expectedNextMove, actualNextMove)
}

func TestNextMove_OutOfHealthFromHazardEdgeCase4(t *testing.T) {
	gr := hazard_related.EdgeCaseRequest4
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	notExpectedNextMove := parseMoveDirectionToString(UP)
	assert.NotEqual(t, notExpectedNextMove, actualNextMove)
}

func TestNextMove_OutOfHealthFromHazardEdgeCase5(t *testing.T) {
	gr := hazard_related.EdgeCaseRequest5
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(DOWN)
	assert.Equal(t, expectedNextMove, actualNextMove)
}

func TestNextMove_NearestPlausibleFoodEdgeCase1(t *testing.T) {
	gr := food_related.EdgeCaseRequest1
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(RIGHT)
	assert.Equal(t, expectedNextMove, actualNextMove)
}

func TestNextMove_NearestPlausibleFoodEdgeCase2(t *testing.T) {
	gr := food_related.EdgeCaseRequest2
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(UP)
	assert.Equal(t, expectedNextMove, actualNextMove)
}

func TestNextMove_NearestPlausibleFoodEdgeCase3(t *testing.T) {
	gr := food_related.EdgeCaseRequest3
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(DOWN)
	assert.Equal(t, expectedNextMove, actualNextMove)
}

func TestNextMove_NearestPlausibleFoodEdgeCase4(t *testing.T) {
	gr := food_related.EdgeCaseRequest4
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(UP)
	assert.Equal(t, expectedNextMove, actualNextMove)
}

func TestNextMove_NearestPlausibleFoodEdgeCase5(t *testing.T) {
	gr := food_related.EdgeCaseRequest5
	a := Init(gr.Board, gr.You)

	actualNextMove := a.NextMove(&gr)
	expectedNextMove := parseMoveDirectionToString(LEFT)
	assert.Equal(t, expectedNextMove, actualNextMove)
}
