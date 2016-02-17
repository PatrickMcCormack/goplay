package collections

// Todo
// 1. add thread safety

import (
  "fmt"
)

type Element struct {
  value interface{}
  next *Element
}

type Queue struct {
  head *Element
  tail *Element
  queueSize int
}

func (q *Queue) push(v interface{}) {
  oldLast := q.tail
  q.tail = &Element{v, nil}
  if q.queueSize == 0 {
    q.head = q.tail
  } else {
    oldLast.next = q.tail
  }
  q.queueSize++
}

func (q *Queue) pop() interface{} {
  if q.head == nil { // case where there are no elements in the queue
    return nil
  }
  value := q.head.value
  q.head = q.head.next
  q.queueSize--
  return value
}

func (q *Queue) size() int {
  return q.queueSize
}

func test() {
  var queue Queue
  queue.push(1)
  queue.push(2)
  queue.push(3)
  fmt.Printf("size  = %d\n", queue.size())
  fmt.Printf("value = %d\n", queue.pop())
  fmt.Printf("size  = %d\n", queue.size())
  fmt.Printf("value = %d\n", queue.pop())
  fmt.Printf("size  = %d\n", queue.size())
  fmt.Printf("value = %d\n", queue.pop())
  fmt.Printf("size  = %d\n", queue.size())
  if queue.pop() == nil {
    fmt.Println("queue is empty")
    fmt.Printf("queue.head == %v\n", queue.head)
    fmt.Printf("queue.tail == %v\n", queue.tail)
  }
}
