package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupStack() *Stack {
	s := New()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	return s
}

func TestPush(t *testing.T) {
	s := setupStack()

	assert.Equal(t, 3, s.Len())
}

func TestPop(t *testing.T) {
	s := setupStack()

	num := s.Pop().(int)

	assert.Equal(t, 2, s.Len())
	assert.Equal(t, 3, num)
}

func TestClear(t *testing.T) {
	s := setupStack()

	assert.NotEqual(t, 0, s.Len())
	s.Clear()
	assert.Equal(t, 0, s.Len())
}
