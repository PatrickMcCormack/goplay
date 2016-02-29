package collections

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Log("Testing Queue functionality")

	var queue Queue
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Printf("size  = %d\n", queue.Size())
	fmt.Printf("value = %d\n", queue.Pop())
	fmt.Printf("size  = %d\n", queue.Size())
	fmt.Printf("value = %d\n", queue.Pop())
	fmt.Printf("size  = %d\n", queue.Size())
	fmt.Printf("value = %d\n", queue.Pop())
	fmt.Printf("size  = %d\n", queue.Size())
	if queue.Pop() == nil {
		fmt.Println("queue is empty")
		fmt.Printf("queue.head == %v\n", queue.head)
		fmt.Printf("queue.tail == %v\n", queue.tail)
	}

	t.Log("Done testing Queue functionality")

}
