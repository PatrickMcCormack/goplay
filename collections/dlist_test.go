package collections

import "testing"

// List logging helper function
func logDList(ll *Dlist, t *testing.T) {
	iterator := ll.Iterator()
	for v, count := iterator(), 0; v != nil; v, count = iterator(), count+1 {
		t.Logf("element %v: %v\n", count, v)
	}
}

func TestDlist(t *testing.T) {

	t.Log("Testing Double Linked List functionality")

	var ll Dlist

	strTestData := []string{"This", "is", "a", "dlinked-list", "yey!"}
	afterDeleteStrTestData := []string{"is", "a", "dlinked-list"}
	elementPointerArray := []*DLElement{nil, nil, nil, nil, nil}

	// ---------------------------------------------------------------------
	t.Log("==Start: Append 1 element and Delete")
	element := ll.Append(strTestData[0])
	if ll.Size() != 1 {
		t.Log("Append function failed")
		t.Fail()
	}
	ll.Delete(element)
	if ll.Size() != 0 {
		t.Log("Delete operation failed")
		t.Fail()
	}
	t.Log("==Completed: Append 1 element and Delete")

	// ---------------------------------------------------------------------
	t.Log("==Start: Append all test elements")
	for i, w := range strTestData {
		newElement := ll.Append(w)
		elementPointerArray[i] = newElement
	}
	if ll.Size() != len(strTestData) {
		t.Log("Append operation failed")
		t.Fail()
	}
	logDList(&ll, t)
	t.Log("==Compeleted: Append all tests elements")

	// ---------------------------------------------------------------------
	t.Log("==Start: Delete head from n element list, where n > 1")
	var previousSize int
	previousSize = ll.Size()
	ll.Delete(elementPointerArray[0])
	if ll.Size() != previousSize-1 {
		t.Log("Delete operation failed")
		t.Fail()
	}
	logDList(&ll, t)
	t.Log("==Compeleted: Delete head from n element list, where n > 1")

	// ---------------------------------------------------------------------
	t.Log("==Start: Delete tail from n element list, where n > 1")
	previousSize = ll.Size()
	ll.Delete(elementPointerArray[len(strTestData)-1])
	if ll.Size() != previousSize-1 {
		t.Log("Delete operation failed")
		t.Fail()
	}
	logDList(&ll, t)
	t.Log("==Compeleted: Delete tail from n element list, where n > 1")

	// ---------------------------------------------------------------------
	t.Log("==Start: Test Iterator")
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
	t.Log("==Start: Delete remaining elements")
	ll.Delete(elementPointerArray[1])
	ll.Delete(elementPointerArray[2])
	ll.Delete(elementPointerArray[3])
	if ll.Size() != 0 {
		t.Log("Delete operation failed")
		t.Fail()
	}
	logDList(&ll, t)
	t.Log("==Compeleted: Delete remaining elements")

	// ---------------------------------------------------------------------
	t.Log("==Start: Append after deleing all elements")
	ll.Append(elementPointerArray[0])
	if ll.Size() != 1 {
		t.Log("Append operation failed")
		t.Fail()
	}
	logDList(&ll, t)
	t.Log("==Compeleted: Append after deleing all elements")

	// ---------------------------------------------------------------------
	t.Log("Finished Testing Double Linked List functionality")
}
