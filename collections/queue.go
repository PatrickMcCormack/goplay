package collections

// Queue - a thread-safe queue implementation
// PatrickMcCormack

import "sync"

// Queue is a data structure that represents a simple Queue
type Queue struct {
	head         *Element
	tail         *Element
	queueSize    int
	sync.RWMutex // composite object
}

// Element is the internal queue data structure that holds the
// value pushed on to the queue and a link to the next element
// in the queue.
type Element struct {
	value interface{}
	next  *Element
}

// Push adds the value (inside an Element) to the back of the queue
func (q *Queue) Push(v interface{}) {
	q.Lock()
	defer q.Unlock()
	oldLast := q.tail
	q.tail = &Element{v, nil}
	if q.queueSize == 0 {
		q.head = q.tail
	} else {
		oldLast.next = q.tail
	}
	q.queueSize++
}

// Pop removes the top elemement in the queue and returns the value
func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()
	if q.head == nil { // case where there are no elements in the queue
		return nil
	}
	value := q.head.value
	q.head = q.head.next
	q.queueSize--
	return value
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()
	return q.queueSize
}
