package collections

import "sync"

// todo:
// 1. delete a node, lots of testing
// 2. MT safety needs more work

// Heap respresent a heap data structure. Heap data structures are often used
// as priority queues.
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

// Calculate the array offset for the parent of a given node
func (aHeap *Heap) parent(index int) int {
	//	aHeap.RLock()
	//	defer aHeap.RUnlock()
	if index == 0 {
		return 0
	}
	return (index - 1) / 2
}

func (aHeap *Heap) leftChildOf(index int) int {
	//	aHeap.RLock()
	//	defer aHeap.RUnlock()
	return 2*index + 1
}

func (aHeap *Heap) rightChildOf(index int) int {
	//	aHeap.RLock()
	//	defer aHeap.RUnlock()
	return 2*index + 2
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

func (aHeap *Heap) upHeap(index int) {
	for index != 0 {
		parent := aHeap.parent(index)

		if aHeap.comparator(aHeap.heap[parent], aHeap.heap[index]) == -1 {
			temp := aHeap.heap[parent]
			aHeap.heap[parent] = aHeap.heap[index]
			aHeap.heap[index] = temp
			index = parent
		} else {
			break
		}
	}
}
