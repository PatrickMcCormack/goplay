package collections

import "testing"

func TestStack(t *testing.T) {
	t.Log("Testing Stack functionality")

	var stack Stack
	intTestData := [...]int{1, 2, 3}
	strTestData := [...]string{"Test Element"}

	// ---------------------------------------------------------------------
	t.Log("==Start: Pushing 1 Element")
	stack.Push(strTestData[0])
	if stack.Size() != 1 {
		t.Log("Push operation failed")
		t.Fail()
	}
	t.Log("==Completed: Pushing 1 Element")

	// ---------------------------------------------------------------------
	t.Log("==Start: Popping an element from a stack of size 1")
	element := stack.Pop()
	t.Logf("Popped %v from stack\n", element)
	if stack.Size() != 0 {
		t.Log("Pop operation failed")
		t.Fail()
	}
	t.Log("==Completed: Popping an element from a stack of size 1")

	// ---------------------------------------------------------------------
	t.Log("==Start: Push several element on to the stack")
	for _, e := range intTestData {
		stack.Push(e)
	}
	if stack.Size() != len(intTestData) {
		t.Log("Push operation failed")
		t.Fail()
	}
	t.Log("==Completed: Push several element on to the stack")

	// ---------------------------------------------------------------------
	t.Log("==Start: Peeking at the head of the stack")
	previousSize := stack.Size()
	element = stack.Peek()
	t.Logf("Peeked \"%v\" from stack\n", element)
	if stack.Size() != previousSize {
		t.Log("Pop operation failed")
		t.Fail()
	}
	if element != intTestData[len(intTestData)-1] {
		t.Log("Peek operation failed")
		t.Fail()
	}
	t.Log("==Completed: Peeking at the head of the stack")

	// ---------------------------------------------------------------------
	t.Log("==Start: Pop all elemement from the stack")
	count := 0
	previousSize = stack.Size()
	for ; stack.Size() > 0; count++ {
		element = stack.Pop()
		sameAs := intTestData[len(intTestData)-1-count]
		if element != sameAs {
			t.Log("Popped element %v, expected %v\n", element, sameAs)
			t.Fail()
		}
		t.Logf("Pop element %v\n", element)
	}
	if stack.Size() != 0 {
		t.Log("Pop operation failed")
		t.Fail()
	}
	if previousSize != count {
		t.Log("Pop operation failed")
		t.Fail()
	}
	t.Log("==Completed: Pop all elemement from the stack")

	// ---------------------------------------------------------------------
	t.Log("Done testing Stack functionality")

}
