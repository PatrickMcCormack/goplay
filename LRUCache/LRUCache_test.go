package collections

import "testing"

func TestLRUCache(t *testing.T) {

	t.Log("Testing LRU Cache functionality")

	var cache LRUCache

	// ---------------------------------------------------------------------
	t.Log("==Start: Initialize Cache")
	cache.Initialize(100)
	t.Log("==Completed: Initialize Cache")

	// ---------------------------------------------------------------------
	t.Log("==Start: Insert first element")
	cache.Add("Key1", "Value1")
	t.Log("==Completed: Insert first element")

	// ---------------------------------------------------------------------
	t.Log("==Start: Find first element")
	entry1 := cache.Get("Key1")
	t.Log(entry1)
	t.Log("==Completed: Find first element")

	// ---------------------------------------------------------------------
	t.Log("==Start: Insert second element")
	cache.Add("Key2", "Value2")
	t.Log("==Completed: Insert second element")

	// ---------------------------------------------------------------------
	t.Log("==Start: Find second element")
	entry2 := cache.Get("Key2")
	t.Log(entry2)
	t.Log("==Completed: Find second element")

	// ---------------------------------------------------------------------
	t.Log("==Start: Test TTL reset on Append to existing entry")

	iterator := cache.ttls.Iterator()
	for v := iterator(); v != nil; v = iterator() {
		t.Logf("element %v", v)
	}

	t.Log("------")

	cache.Add("Key1", "Value1 updated")

	iterator = cache.ttls.Iterator()
	for v := iterator(); v != nil; v = iterator() {
		t.Logf("element %v", v)
	}

	// ---------------------------------------------------------------------
	t.Log("Finished Testing Linked List functionality")
}
