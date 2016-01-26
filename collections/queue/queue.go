package main

// Unfinished Queue Implementation
// PatrickMcCormack

import (
  "fmt"
)

type Element struct {
  value interface{}
  next *Element // points to the tail
}

type Queue struct {
  head *Element
  tail *Element
  queueSize int
}

func (q *Queue) push(v interface{}) {
  q.Tail := &Element{v, q.tail}
}

func (q *Queue) pop() interface{} {
  return "xx"
}

func (q *Queue) size() int {
  return q.queueSize
}

func main() {
  var queue Queue
  queue.push("xx")
  fmt.Printf("Hello")
}
