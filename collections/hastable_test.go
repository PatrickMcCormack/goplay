package collections

import "testing"

var nameTestData = []string{"Moby", "Radiohead"}
var valueTestData = []string{"Extreme Ways", "The Bends"}

func TestHashtable(t *testing.T) {

	t.Log("Testing Hashtable functionality")

	// ---------------------------------------------------------------------
	t.Log("==Start: Creating HT, insert an element, retrieve & check element")

	var h HashTable
	h.Initialize(0)

	h.Insert(nameTestData[0], valueTestData[0])
	value, err := h.Find(nameTestData[0])
	t.Logf("Got %v, expecting %v\n", value, valueTestData[0])
	if value != valueTestData[0] {
		t.Log("Hashtable returning wrong value")
		t.Fail()
	}
	if err != nil {
		t.Logf("Find returning error - %v\n", err)
		t.Fail()
	}

	h.Insert(nameTestData[1], valueTestData[1])
	value, err = h.Find(nameTestData[1])
	t.Logf("Got %v, expecting %v", value, valueTestData[1])
	if value != valueTestData[1] {
		t.Log("Hashtable returning wrong value")
		t.Fail()
	}
	if err != nil {
		t.Logf("Find returning error - %v\n", err)
		t.Fail()
	}

	value, err = h.Find("Won't find me")
	if value != nil {
		t.Log("Hashtable found an impossible record")
		t.Fail()
	}
	if err != nil {
		t.Logf("Find returning error - %v\n", err)
		t.Fail()
	}

	// Test Delete
	h.Delete(nameTestData[0])
	value, err = h.Find(nameTestData[0])
	t.Logf("Got %v, nil", value)
	if err != nil {
		t.Logf("Delete returning error - %v\n", err)
		t.Fail()
	}

	t.Log("==Completed: Creating HT, insert an element, retrieve & check element")

	// ---------------------------------------------------------------------
	t.Log("Done testing Hashtable functionality")

}
