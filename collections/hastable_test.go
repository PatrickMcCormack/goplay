package collections

import "testing"

var nameTestData=[]string{"Moby"}
var valueTestData=[]string{"Extreme Ways"}

func TestHashtable(t *testing.T) {

	t.Log("Testing Hashtable functionality")

	// ---------------------------------------------------------------------
	t.Log("==Start: Creating HT, insert an element, retrieve & check element")

		var h HashTable
		h.Initialize()
		h.Insert(nameTestData[0], valueTestData[0])
		_, value := h.Find(nameTestData[0])
		t.Logf("Got %v, expecting %v", value, valueTestData[0])
		if value !=  valueTestData[0] {
			t.Log("Hashtable returning wrong value")
			t.Fail()
		}

	t.Log("==Completed: Creating HT, insert an element, retrieve & check element")

	// ---------------------------------------------------------------------
	t.Log("Done testing Hashtable functionality")


}
