package collections

// Simple Stack Implementation
// Todo:
//  1. Add thread safety

type StackElement struct {
	value interface{}
	next  *StackElement
}

type Stack struct {
	top       *StackElement
	stackSize int
}

func (s *Stack) size() int {
	return s.stackSize
}

func (s *Stack) push(value interface{}) {
	e := &StackElement{value, nil}
	e.next = s.top
	s.top = e
	s.stackSize++
}

func (s *Stack) pop() interface{} {
	if s.stackSize == 0 {
		return nil
	}
	rv := s.top.value
	s.top = s.top.next
	s.stackSize--
	return rv
}

func (s *Stack) peek() interface{} {
	if s.stackSize == 0 {
		return nil
	}
	rv := s.top.value
	return rv
}
