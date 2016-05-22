package collections

import "sync"

// Dlist - a thread-safe double  linked list implementation
//
// Todo
// 1. func (ll *Dlist) InsertBefore(element *DLElement) {
// 2. func (ll *Dlist) InsertAfter(element *DLElement) {
// 3. Consider allowing a direction flag on creation so
//    the iterator can be a forward or backward iterator.
//
// PatrickMcCormack

// Dlist data structure.
type Dlist struct {
	head         *DLElement
	tail         *DLElement
	listSize     int
	sync.RWMutex // composite object
}

// DLElement is the contents of the linked list
type DLElement struct {
	value    interface{}
	previous *DLElement
	next     *DLElement
}

// Size returns the number of elements in the linked list
func (ll *Dlist) Size() int {
	ll.RLock()
	defer ll.RUnlock()
	return ll.listSize
}

// Append a value to the linked list
func (ll *Dlist) Append(v interface{}) *DLElement {
	ll.Lock()
	defer ll.Unlock()
	newele := &DLElement{value: v, next: nil, previous: nil}
	if ll.listSize == 0 {
		ll.head = newele
		ll.tail = newele
	} else {
		newele.previous = ll.tail
		ll.tail.next = newele
		ll.tail = newele
	}
	ll.listSize++
	return newele
}

// Delete by element address, preferrer and most efficient O(1) way to
// delete an element in the list.
// Todo create error object and return in error conditions
func (ll *Dlist) Delete(element *DLElement) {
	ll.Lock()
	defer ll.Unlock()
	if ll.listSize == 0 {
		return
	}

	if element.previous == nil {
		ll.head = element.next
	} else {
		element.previous.next = element.next
	}

	if element.next == nil {
		ll.tail = element.previous
	} else {
		element.next.previous = element.previous
	}

	ll.listSize--
}

// Iterator returns a closure that allows iteration over the list.
// If there are no more values to return the iterator closure returns nil.
func (ll *Dlist) Iterator() func(...bool) interface{} {
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
		// iterate
		rvalue := current.value
		current = current.next
		return rvalue
	}
}
