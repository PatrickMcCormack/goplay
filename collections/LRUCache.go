package collections

import "sync"

// LRUCache data structure
type LRUCache struct {
	entries      HashTable
	ttls         Dlist // ttls is the eviction sequence for the LRU
	sync.RWMutex       // composite object
}

// LRUCacheEntry keeps track of the value passed by the user and
// the "TTL" for the cache entry
type LRUCacheEntry struct {
	value interface{}
	ttl   *DLElement
}

// Initialize the cache by passing the size of the desired cache.
// Important Note, this method is not idempotent
func (cache *LRUCache) Initialize(size int) {
	cache.entries.Initialize(size)
}

// Add a named value to the cache. Adding an existing named cache
// entry updates the entry with the new value and resets the cache
// entry TTL.
func (cache *LRUCache) Add(name interface{}, cacheValue interface{}) {
	cache.Lock()
	defer cache.Unlock()
	cacheEntry, _ := cache.entries.Find(name)
	if cacheEntry != nil { // updating
		cache.ttls.Delete(cacheEntry.(*LRUCacheEntry).ttl)
		cache.entries.Delete(name)
		newTTL := cache.ttls.Append(name)
		entry := &LRUCacheEntry{value: cacheValue, ttl: newTTL}
		cache.entries.Insert(name, entry)
	} else {
		newTTL := cache.ttls.Append(name)
		entry := &LRUCacheEntry{value: cacheValue, ttl: newTTL}
		cache.entries.Insert(name, entry)
	}
}

// Get a named value from the cache
func (cache *LRUCache) Get(name interface{}) interface{} {
	cache.RLock()
	defer cache.RUnlock()
	entry, _ := cache.entries.Find(name)
	var rval interface{}
	rval = nil
	if entry != nil {
		rval = entry.(*LRUCacheEntry).value
	}
	return rval
}
