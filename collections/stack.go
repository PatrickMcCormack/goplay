package collections

import "sync"

// Stack - a thread-safe stack implementation
// PatrickMcCormack

// Stack data structure
type Stack struct {
	top          *StackElement
	stackSize    int
	sync.RWMutex // composite object
}

// StackElement is the contents of the stack
type StackElement struct {
	value interface{}
	next  *StackElement
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	s.RLock()
	defer s.RUnlock()
	return s.stackSize
}

// Push an element on the stack, can be of any type.
func (s *Stack) Push(value interface{}) {
	s.Lock()
	defer s.Unlock()
	e := &StackElement{value, nil}
	e.next = s.top
	s.top = e
	s.stackSize++
}

// Pop and element off the top of the stack
func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()
	if s.stackSize == 0 {
		return nil
	}
	rv := s.top.value
	s.top = s.top.next
	s.stackSize--
	return rv
}

// Peek at the top of the stack without removing an element.
func (s *Stack) Peek() interface{} {
	s.RLock()
	defer s.RUnlock()
	if s.stackSize == 0 {
		return nil
	}
	rv := s.top.value
	return rv
}
