package collections

import "sync"

// Dlist - a thread-safe double  linked list implementation
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

// Todo
// func (ll *Dlist) InsertBefore(element *DLElement) {
// func (ll *Dlist) InsertAfter(element *DLElement) {

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
// Todo: consider allowing a direction flag on creation so the iterator
// can be a forward or backward iterator.
func (ll *Dlist) Iterator() func() interface{} {
	ll.RLock()
	current := ll.head
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
