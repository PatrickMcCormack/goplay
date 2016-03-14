package collections

import "sync"

// todo:
// 1. delete a node implementation
// 2. downheap implementation
// 3. lots of testing

// Heap is a heap data structure. Heap data structures are often used
// as priority queues. The Heap must be initilized before use by calling
// the method Initilize.
type Heap struct {
	heap         []interface{}
	numElements  int
	comparator   Comparator
	sync.RWMutex // composite object
}

// Initialize a heap to an inital size and set the comparator for
// the type stored in the heap. This method is *not* idempotent.
func (aHeap *Heap) Initialize(initialSize int, c Comparator) {
	aHeap.Lock()
	defer aHeap.Unlock()
	aHeap.heap = make([]interface{}, initialSize)
	aHeap.comparator = c
}

// Insert a value in to the heap
func (aHeap *Heap) Insert(value interface{}) {
	aHeap.Lock()
	defer aHeap.Unlock()
	// FIXME -- handle the case where the max heap size has been reached.
	aHeap.heap[aHeap.numElements] = value
	aHeap.upHeap(aHeap.numElements)
	aHeap.numElements++
}

// Delete an element from the Heap
func (aHeap *Heap) Delete(index int) {
	/* Pseudocode
		   1. Delete a node from the array
	        (this creates a "hole" and the tree is no longer "complete")

	     2. Replace the deletion node
	        with the "fartest right node" on the lowest level
	        of the Binary Tree
	        (This step makes the tree into a "complete binary tree")

	     3. Heapify (fix the heap):

	          if ( value in replacement node < its parent node )
	             Filter the replacement node UP the binary tree
	          else
	 	    		  Filter the replacement node DOWN the binary tree
	*/

}

// Size returns the number of elements in the heap
func (aHeap *Heap) Size() int {
	aHeap.RLock()
	defer aHeap.RUnlock()
	return aHeap.numElements
}

// Note internal methods are not thread-safe, they cannot be
// otherwise they would cause deadlocks.

// Calculate the array offset for the parent of a given node
func (aHeap *Heap) parent(index int) int {
	if index == 0 {
		return 0
	}
	return (index - 1) / 2
}

// Calculate the array offset for the left child of a given node
func (aHeap *Heap) leftChildOf(index int) int {
	return 2*index + 1
}

// Calculate the array offset for the right child of a given node
func (aHeap *Heap) rightChildOf(index int) int {
	return 2*index + 2
}

// (Potentially) repair the heap starting at a left node, working
// up to the root node.
func (aHeap *Heap) upHeap(index int) {
	for index != 0 {
		parent := aHeap.parent(index)
		if aHeap.comparator(aHeap.heap[parent], aHeap.heap[index]) == LessThan {
			temp := aHeap.heap[parent]
			aHeap.heap[parent] = aHeap.heap[index]
			aHeap.heap[index] = temp
			index = parent
		} else {
			break
		}
	}
}

// (Potentially) repair the heap starting at a root node, working
// down through the tree.
func (aHeap *Heap) downHeap(index int) {
	// Todo
}
