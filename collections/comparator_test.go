package collections

import "testing"

func TestComparators(t *testing.T) {
	t.Log("Testing Comparator functionality")

	// ---------------------------------------------------------------------
	t.Log("==Start: Testing IntComparator")

	result := IntComparator(1,2)
	t.Logf("IntComparator: %v expected, got %v", LessThan, result)
	if result != LessThan {
		t.Log("LessThan failed on IntComparator")
		t.Fail()
	}

	result = IntComparator(1,1)
	t.Logf("IntComparator: %v expected, got %v", Equal, result)
	if result != Equal {
		t.Log("Equal failed on IntComparator")
		t.Fail()
	}

	result = IntComparator(2,1)
	t.Logf("IntComparator: %v expected, got %v", GreaterThan, result)
	if result != GreaterThan {
		t.Log("GreaterThan failed on IntComparator")
		t.Fail()
	}

	t.Log("==Completed: Testing IntComparator")

	// ---------------------------------------------------------------------
	t.Log("==Start: Testing StringComparator")

	result = StringComparator("aaaa","bbbb")
	t.Logf("StringComparator: %v expected, got %v", LessThan, result)
	if result != LessThan {
		t.Log("LessThan failed on StringComparator")
		t.Fail()
	}

	result = StringComparator("aaaa","aaaa")
	t.Logf("StringComparator: %v expected, got %v", Equal, result)
	if result != Equal {
		t.Log("Equal failed on StringComparator")
		t.Fail()
	}

	result = StringComparator("bbbb","aaaa")
	t.Logf("StringComparator: %v expected, got %v", GreaterThan, result)
	if result != GreaterThan {
		t.Log("GreaterThan failed on StringComparator")
		t.Fail()
	}

	t.Log("==Completed: Testing StringComparator")

	// ---------------------------------------------------------------------
	t.Log("Done Testing Comparator functionality")

}
