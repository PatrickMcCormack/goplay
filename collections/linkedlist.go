package collections

import "sync"

// LinkedList - a thread-safe linked list implementation
// PatrickMcCormack

// Stack data structure. The stack should be initalized on creation:
// Example: ll := LinkedList{comparator: StringComparator}
type LinkedList struct {
	head         *LLElement
	tail         *LLElement // optimization
	listSize     int
	comparator   Comparator
	sync.RWMutex // composite object
}

// LLElement is the contents of the linked list
type LLElement struct {
	value interface{}
	next  *LLElement
}

// Size returns the number of elements in the linked list
func (ll *LinkedList) Size() int {
	ll.RLock()
	defer ll.RUnlock()
	return ll.listSize
}

// Insert a value into a linked list
func (ll *LinkedList) Insert(v interface{}) {
	ll.Lock()
	defer ll.Unlock()
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

// Iterator returns a closure that allows iteration over the linked list.
// If there are no more values to return the iterator closure returns nil.
func (ll *LinkedList) Iterator() func() interface{} {
	current := ll.head
	ll.RLock()
	// return a closure over the variables in scope
	return func() interface{} {
		//		if current == nil || current.next == nil {
		if current == nil {
			ll.RUnlock()
			return nil
		}
		rvalue := current.value
		current = current.next
		return rvalue
	}
}
