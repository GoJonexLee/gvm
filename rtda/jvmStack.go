package rtda

import (
	"fmt"
)

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize, 0, nil}
}

func (s *Stack) isEmpty() bool {
	return s._top == nil
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("StackOverflowError")
	}

	if s._top != nil {
		frame.lower = s._top
	}

	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}

	top := s._top
	s._top = top.lower
	top.lower = nil
	s.size--
	return top
}

func (s *Stack) clear() {
	for !s.isEmpty() {
		s.pop()
	}
}

func (s *Stack) top() *Frame {
	if s.top == nil {
		panic("jvm stack is empty!")
	}
	return s._top
}

func (s *Stack) topN(n uint) *Frame {
	if s.size < n {
		panic(fmt.Sprintf("jvm stack size:%v n:%v", s.size, n))
	}

	frame := s._top
	for n > 0 {
		fram = fram.lower
		n--
	}
	return frame
}
