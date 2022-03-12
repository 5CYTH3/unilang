package util

import "scythe.com/uni/src/tokens"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value tokens.Tokens
	next  *Element
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Top() tokens.Tokens {
	return s.top.value
}

func (s *Stack) Push(value tokens.Tokens) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value any) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
