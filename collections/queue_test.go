package collections

import "testing"

func TestQueue(t *testing.T) {
	t.Log("Testing Queue functionality")

	var queue Queue
	intTestData := [...]int{1, 2, 3}
	strTestData := [...]string{"Test Element"}

	// ---------------------------------------------------------------------
	t.Log("==Start: Pushing 1 Element")
	queue.Push(strTestData[0])
	if queue.Size() != 1 {
		t.Log("Push operation failed")
		t.Fail()
	}
	t.Log("==Completed: Pushing 1 Element")

	// ---------------------------------------------------------------------
	t.Log("==Start: Popping an element from a queue of size 1")
	element := queue.Pop()
	t.Logf("Popped %v from queue\n", element)
	if queue.Size() != 0 {
		t.Log("Pop operation failed")
		t.Fail()
	}
	t.Log("==Completed: Popping an element from a queue of size 1")

	// ---------------------------------------------------------------------
	t.Log("==Start: Push several element on to the queue")
	for _, e := range intTestData {
		queue.Push(e)
	}
	if queue.Size() != len(intTestData) {
		t.Log("Push operation failed")
		t.Fail()
	}
	t.Log("==Completed: Push several element on to the queue")

	// ---------------------------------------------------------------------
	t.Log("==Start: Peeking at the head of the queue")
	previousSize := queue.Size()
	element = queue.Peek()
	t.Logf("Peeked \"%v\" from queue\n", element)
	if queue.Size() != previousSize {
		t.Log("Pop operation failed")
		t.Fail()
	}
	if element != intTestData[0] {
		t.Log("Peek operation failed")
		t.Fail()
	}
	t.Log("==Completed: Peeking at the head of the queue")

	// ---------------------------------------------------------------------
	t.Log("==Start: Pop all elemement from the queue")
	count := 0
	previousSize = queue.Size()
	for ; queue.Size() > 0; count++ {
		element = queue.Pop()
		t.Logf("Pop element %v\n", element)
	}
	if queue.Size() != 0 {
		t.Log("Pop operation failed")
		t.Fail()
	}
	if previousSize != count {
		t.Log("Pop operation failed")
		t.Fail()
	}
	t.Log("==Completed: Pop all elemement from the queue")

	// ---------------------------------------------------------------------
	t.Log("Done testing Queue functionality")

}
