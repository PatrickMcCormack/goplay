package collections

import "testing"

// List logging helper function
func logList(ll *LinkedList, t *testing.T) {
	iterator := ll.Iterator()
	for v, count := iterator(), 0; v != nil; v, count = iterator(), count+1 {
		t.Logf("element %v: %v\n", count, v)
	}
}

func TestLinkedList(t *testing.T) {

	var previousSize int

	t.Log("Testing Linked List functionality")

	ll := LinkedList{}

	strTestData := []string{"This", "is", "a", "linked-list"}
	afterDeleteStrTestData := []string{"is", "a"}

	// ---------------------------------------------------------------------
	t.Log("==Start: Insert 1 element and delete")
	ll.Insert(strTestData[0])
	if ll.Size() != 1 {
		t.Log("Insert function failed")
		t.Fail()
	}
	ll.Delete(strTestData[0])
	if ll.Size() != 0 {
		t.Log("Delete operation failed")
		t.Fail()
	}
	t.Log("==Completed: Insert 1 element and delete")

	// ---------------------------------------------------------------------
	t.Log("==Start: Insert all test elements")
	for _, w := range strTestData {
		ll.Insert(w)
	}
	if ll.Size() != len(strTestData) {
		t.Log("Insert operation failed")
		t.Fail()
	}
	t.Log("==Compeleted: Insert all tests elements")

	// ---------------------------------------------------------------------
	t.Log("==Start: Delete head from n element list, where n > 1")
	previousSize = ll.Size()
	ll.Delete(strTestData[0])
	if ll.Size() != previousSize-1 {
		t.Log("Delete operation failed")
		t.Fail()
	}
	t.Log("==Compeleted: Delete head from n element list, where n > 1")

	// ---------------------------------------------------------------------
	t.Log("==Start: Delete tail from n element list, where n > 1")
	previousSize = ll.Size()
	ll.Delete(strTestData[len(strTestData)-1])
	if ll.Size() != previousSize-1 {
		t.Log("Delete operation failed")
		t.Fail()
	}
	t.Log("==Compeleted: Delete tail from n element list, where n > 1")

	// ---------------------------------------------------------------------
	t.Log("==Start: Test Closure Iterator")
	count := 0
	iterator := ll.Iterator()
	for v := iterator(); v != nil; v = iterator() {
		t.Logf("element %v: %v, expecting %v\n", count, v, afterDeleteStrTestData[count])
		if v.(string) != afterDeleteStrTestData[count] {
			t.Log("Iterator is not returning the correct elements")
			t.Fail()
		}
		count++
	}
	if ll.Size() != count {
		t.Log("Iterator did not return the correct number of elements")
		t.Fail()
	}
	t.Log("==Compeleted: Test Iterator")

	// ---------------------------------------------------------------------
	t.Log("==Start: Test Closing a Closure Iterator")
	cItr := ll.Iterator()
	val := cItr()
	if val.(string) != afterDeleteStrTestData[0] {
		t.Log("Iterator is not returning the correct elements")
		t.Logf("found %v, expecting %v\n", val, afterDeleteStrTestData[0])
		t.Fail()
	}
	cItr(true) // close the iterator
	val = cItr()
	if val != nil {
		t.Log("Iterator is working a expected")
		t.Logf("found %v, expecting nil\n", val)
		t.Fail()
	}
	t.Log("==Compeleted: Test Closing a Closure Iterator")
	// ---------------------------------------------------------------------
	t.Log("==Start: Test Object Iterator")
	siCount := 0
	sItr := ll.IteratorS()
	for v := sItr.Next(); v != nil; v = sItr.Next() {
		t.Logf("element %v: %v, expecting %v\n", siCount, v, afterDeleteStrTestData[siCount])
		if v.(string) != afterDeleteStrTestData[siCount] {
			t.Log("Iterator is not returning the correct elements")
			t.Fail()
		}
		siCount++
	}
	if ll.Size() != siCount {
		t.Log("Iterator did not return the correct number of elements")
		t.Fail()
	}
	t.Log("==Compeleted: Test Stateful Iterator")

	// ---------------------------------------------------------------------
	t.Log("==Start: Delete remaining elements")
	ll.Delete(strTestData[1])
	ll.Delete(strTestData[2])
	if ll.Size() != 0 {
		t.Log("Delete operation failed")
		t.Fail()
	}
	t.Log("==Compeleted: Delete remaining elements")

	// ---------------------------------------------------------------------
	t.Log("==Start: Insert after deleing all elements")
	ll.Insert(strTestData[0])
	if ll.Size() != 1 {
		t.Log("Insert operation failed")
		t.Fail()
	}
	t.Log("==Compeleted: Insert after deleing all elements")

	logList(&ll, t)

	// ---------------------------------------------------------------------
	t.Log("Finished Testing Linked List functionality")
}
