package evaluator

import "pisang/internal/pkg/token"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value token.Token
	next  *Element
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Top() interface{} {
	return s.top.value
}

func (s *Stack) Push(value token.Token) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value token.Token) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return token.Token{}
}
