package pque

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockPQImplementation struct {
	data     string
	priority int
}

func (m MockPQImplementation) GetPriority() float64 {
	return float64(m.priority)
}

func TestPriorityQueue(t *testing.T) {
	que := &PriorityQueue{}

	expectedOrder := []MockPQImplementation{
		{"apple", 9},
		{"orange", 12},
		{"banana", 3},
	}

	for _, fruit := range expectedOrder {
		heap.Push(que, fruit)
	}

	assert.Equal(t, 3, que.Len())

	assert.Equal(t, expectedOrder[2], heap.Pop(que))
	assert.Equal(t, expectedOrder[0], heap.Pop(que))
	assert.Equal(t, expectedOrder[1], heap.Pop(que))
}
