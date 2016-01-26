package main

// Simple Stack Implementation 
// PatrickMcCormack

import (
  "fmt"
)

type Element struct {
  value interface{}
  next *Element
}

type Stack struct {
  top *Element
  stackSize int
}

func (s *Stack) size() int {
  return s.stackSize
}

func (s *Stack) push(value interface{}) {
  e := &Element{value, nil}
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

func main() {
  var stack Stack
  stack.push("hello")
  stack.push("there")
  fmt.Printf("Stack size   = %d\n", stack.size())
  fmt.Printf("Popped value = %s\n", stack.pop())
  fmt.Printf("Popped value = %s\n", stack.pop())
  fmt.Printf("Popped value = %s\n", stack.pop())
  fmt.Printf("Stack size   = %d\n", stack.size())

  // as long as you know the types in your stack you can pop them off
  // and type convert to the value you want since their type is
  // interface{} on the stack. In a real program you'd need to check
  // ok after the type assertion because if it's false you won't be
  // able to use the v1 value without getting a runtime error.

  stack.push("hello")
  v1, ok := stack.pop().(string)
  v1 += ", there"
  fmt.Printf("Type conversion to string is %t %s\n", ok, v1)
  fmt.Printf("The value is %s\n", v1)
}
