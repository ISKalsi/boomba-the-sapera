package stack

import "container/list"

type Stack struct {
	elements *list.List
}

func New() *Stack {
	return &Stack{
		elements: list.New(),
	}
}

func (s *Stack) Len() int {
	return s.elements.Len()
}

func (s *Stack) Push(v interface{}) {
	s.elements.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	e := s.elements.Back()
	s.elements.Remove(e)
	return e.Value
}

func (s *Stack) Clear() {
	s.elements.Init()
}
