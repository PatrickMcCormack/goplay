package collections

// Queue - a thread-safe queue implementation
// PatrickMcCormack

import "sync"

// Queue is a data structure that represents a simple Queue
type Queue struct {
	head         *QElement
	tail         *QElement
	queueSize    int
	sync.RWMutex // composite object
}

// QElement is the internal queue data structure that holds the
// value pushed on to the queue and a link to the next QElement
// in the queue.
type QElement struct {
	value interface{}
	next  *QElement
}

// Push adds the value (inside an QElement) to the back of the queue
func (q *Queue) Push(v interface{}) {
	q.Lock()
	defer q.Unlock()
	oldLast := q.tail
	q.tail = &QElement{v, nil}
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
	if q.head == nil { // case where there are no QElements in the queue
		return nil
	}
	value := q.head.value
	q.head = q.head.next
	q.queueSize--
	return value
}

// Peek returns the value at the top of the queue without removing the
// element from the queue.
func (q *Queue) Peek() interface{} {
	q.Lock()
	defer q.Unlock()
	if q.head == nil { // case where there are no QElements in the queue
		return nil
	}
	return q.head.value
}

// Size returns the number of QElements in the queue
func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()
	return q.queueSize
}
