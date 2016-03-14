package collections

import "testing"

// Int Test Data Set
var intTestData = [...]int{10, 1, 5, 15, 20, 4, 13}
var intPreOrderTestData = [...]int{10, 1, 5, 4, 15, 13, 20}
var intInOrderTestData = [...]int{1, 4, 5, 10, 13, 15, 20}
var intPostOrderTestData = [...]int{4, 5, 1, 13, 20, 15, 10}

// ---------------------------------------------------------------------
func TestInsertRootNode(t *testing.T) {
	t.Log("==Start: Insert Root Node, test Size and Depth")
	tree := &Tree{comparator: IntComparator}
	tree.Insert(10)
	t.Logf("The size of the tree is: %v\n", tree.Size())
	t.Logf("The max depth is: %v\n", tree.Depth())
	if tree.Size() != 1 {
		t.Log("Tree was not created correctly, size is incorrect.")
		t.Fail()
	}
	if tree.Depth() != 1 {
		t.Log("Tree was not created correctly, depth is incorrect.")
		t.Fail()
	}
	// Log the tree for quick debug purposes
	// tree.BreadthFirstTraversalWithLevels()
	t.Log("==Completed: Insert Root Node, test Size and Depth")
}

// ---------------------------------------------------------------------
func TestSmallTreeForSizeAndDepth(t *testing.T) {
	t.Log("==Start: Insert more nodes and test size and depth")
	tree := &Tree{comparator: IntComparator}
	for _, e := range intTestData {
		tree.Insert(e)
	}
	t.Logf("The size of the tree is: %v\n", tree.Size())
	t.Logf("The max depth is: %v\n", tree.Depth())
	if tree.Size() != len(intTestData) {
		t.Log("Tree was not created correctly, size is incorrect.")
		t.Fail()
	}
	if tree.Depth() != 4 {
		t.Log("Tree was not created correctly, depth is incorrect.")
		t.Fail()
	}
	// Log the tree for quick debug purposes
	// tree.BreadthFirstTraversalWithLevels()
	t.Log("==Completed: Insert more nodes and test size and depth")
}

// ---------------------------------------------------------------------
func TestIteratorsDriver(t *testing.T) {
	t.Log("==Start: TestIteratorsDriver")
	testIterator(t, "PreOrder", PreOrder)
	testIterator(t, "InOrder", InOrder)
	testIterator(t, "PostOrder", PostOrder)
	t.Log("==Completed: TestIteratorsDriver")
}

// ---------------------------------------------------------------------
func testIterator(t *testing.T, name string, order IteratorOrder) {
	t.Logf("==Start: Insert data, iterate in %v, check order is correct", name)
	tree := &Tree{comparator: IntComparator}
	for _, e := range intTestData {
		tree.Insert(e)
	}
	index := 0
	goldenValue := 0
	iterator := tree.Iterator(order)
	for v := iterator(); v != nil; v = iterator() {
		switch {
		case order == PreOrder:
			goldenValue = intPreOrderTestData[index]
		case order == InOrder:
			goldenValue = intInOrderTestData[index]
		case order == PostOrder:
			goldenValue = intPostOrderTestData[index]
		}
		t.Logf("iterator: %v, GoldenData: %v\n", v, goldenValue)
		if v.(int) != goldenValue {
			t.Logf("%v iterator not returning data correctly", name)
			t.Fail()
			break
		}
		index++
	}
	t.Logf("==Completed: Insert data, iterate in %v, check order is correct", name)
}

// ---------------------------------------------------------------------
func TestRecursiveFunctionsDriver(t *testing.T) {
	t.Log("==Start: TestRecursiveFunctionsDriver")
	testRecursiveFunctions(t, "PreOrder", PreOrder)
	testRecursiveFunctions(t, "InOrder", InOrder)
	testRecursiveFunctions(t, "PostOrder", PostOrder)
	t.Log("==Completed: TestRecursiveFunctionsDriver")
}

// ---------------------------------------------------------------------
func testRecursiveFunctions(t *testing.T, name string, order IteratorOrder) {
	t.Logf("==Start: Insert data, iterate in %v, check order is correct", name)
	tree := &Tree{comparator: IntComparator}
	for _, e := range intTestData {
		tree.Insert(e)
	}
	var list LinkedList
	tree.RecursiveTraversal(order, &list)
	if list.Size() == 0 {
		t.Logf("Recursive %v returned 0 results", name)
		t.Fail()
	}
	index := 0
	goldenValue := 0
	iterator := list.Iterator()
	for v := iterator(); v != nil; v = iterator() {
		switch {
		case order == PreOrder:
			goldenValue = intPreOrderTestData[index]
		case order == InOrder:
			goldenValue = intInOrderTestData[index]
		case order == PostOrder:
			goldenValue = intPostOrderTestData[index]
		}
		t.Logf("Iterator: %v, GoldenData: %v\n", v, goldenValue)
		if v.(int) != goldenValue {
			t.Log("Recursive pre-order not correct")
			t.Fail()
			break
		}
		index++
	}
	t.Log("==Completed: Insert data, iterate in pre-order, check order is correct")
}

/*
Todo
 . Breadth First Traversal testing
 . MT testing
*/
