package collections

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {

	t.Log("Testing Heap functionality")

	var h Heap
	h.Initialize(50, IntComparator)

	h.Insert(50)
	fmt.Println(h.heap)
	h.Insert(80)
	fmt.Println(h.heap)
	h.Insert(90)
	fmt.Println(h.heap)
	h.Insert(100)
	fmt.Println(h.heap)

	fmt.Println(len(h.heap))

	t.Log("Finished Testing Heap functionality")
}
