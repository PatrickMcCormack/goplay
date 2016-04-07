package collections

import "sync"

// LinkedList - a thread-safe linked list implementation
// PatrickMcCormack

// LinkedList data structure.
type LinkedList struct {
	head         *LLElement
	tail         *LLElement
	listSize     int
	sync.RWMutex // composite object
}

// LLElement is the contents of the linked list
type LLElement struct {
	value interface{}
	next  *LLElement
}

// IteratorState - a stateful iterator for a LinkedList
type IteratorState struct {
	current *LLElement
	list    *LinkedList
	closed  bool
}

// Size returns the number of elements in the linked list
func (ll *LinkedList) Size() int {
	ll.RLock()
	defer ll.RUnlock()
	return ll.listSize
}

// Insert a value into a linked list
func (ll *LinkedList) Insert(v interface{}) {
	// FIXME commented out because of deadlock in hashtable in LRU
	// review lock strategy for all data structures in collections
	//	ll.Lock()
	//	defer ll.Unlock()
	newele := &LLElement{value: v, next: nil}
	if ll.listSize == 0 {
		ll.head = newele
		ll.tail = newele
	} else {
		ll.tail.next = newele
		ll.tail = newele
	}
	ll.listSize++
}

// Delete a value from a linked list
func (ll *LinkedList) Delete(v interface{}) {
	ll.Lock()
	defer ll.Unlock()
	if ll.listSize == 0 {
		return
	}
	if ll.listSize == 1 {
		ll.head = nil
		ll.tail = nil
		ll.listSize = 0
		return
	}
	current := ll.head
	previous := current
	for ; current != nil; previous, current = current, current.next {
		if current.value == v {
			break
		}
	}
	if current == ll.head {
		ll.head = ll.head.next
	} else if current == ll.tail {
		ll.tail = previous
	}
	// remove the node
	previous.next = current.next
	ll.listSize--
}

// IteratorS returns an iterator state object.
func (ll *LinkedList) IteratorS() *IteratorState {
	ll.RLock()
	return &IteratorState{current: ll.head, list: ll, closed: false}
}

// Next returns the value at the current iterator position and moves to
// the next position. Returns nil if the end of the collection has been
// reached.
func (is *IteratorState) Next() interface{} {
	if is.closed == true {
		return nil
	}
	if is.current == nil {
		is.Close()
		return nil
	}
	rvalue := is.current.value
	is.current = is.current.next
	return rvalue
}

// Close the Iterator, cannot be used again. This also releases there
// readlock on the underlying linked list.
func (is *IteratorState) Close() {
	if is.closed == false {
		is.closed = true
		is.list.RUnlock()
	}
}

// Iterator returns a closure that allows iteration over the linked list.
// If there are no more values to return the iterator closure returns nil
// and the readloack on the linked list is released.
// To close a CIterator the iterator shold be called with the parameter true,
// this releases the read lock and marks the iterator as closed.
func (ll *LinkedList) Iterator() func(...bool) interface{} {
	ll.RLock()
	current := ll.head
	closed := false
	// return a closure over the variables in scope
	return func(close ...bool) interface{} {
		if closed == true {
			return nil
		}
		// Using a varadic parameter as a hacky way to implement default values
		// so that the closure can be called with zero or more parameter.
		closeItr := false
		for _, val := range close {
			closeItr = val
		}
		if closeItr == true || current == nil {
			ll.RUnlock()
			closed = true
			return nil
		}
		rvalue := current.value
		current = current.next
		return rvalue
	}
}
