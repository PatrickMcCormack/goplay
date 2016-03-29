package collections

// hashtable - an implementation of a hashtable/dictionary
// PatrickMcCormack

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
	"sync"
)

type hashBucket struct {
	name  interface{}
	value interface{}
}

// HashTable is a structure which is used to implement a hash table.
// This implementation resolves collisions with chaining.
// The hashtable must be initialized before use either by calling the
// Initialize method or by initlization of the buckets field on creation.
type HashTable struct {
	buckets      []*LinkedList
	numEntries   int
	sync.RWMutex // composite object
}

var (
	// DefaultNumBuckets is  the number of buckets created
	// if <=0 is passed to Initalize
	DefaultNumBuckets = 101 // Bucketsize is best if Prime
)

// Initialize a HashTable, to create a HashTable with a non-default number
// of buckets you must pass a non-zero positive integer to Intialize.
// The Initialize method is not idempotent.
func (h *HashTable) Initialize(size int) {
	h.Lock()
	defer h.Unlock()
	if size <= 0 {
		h.buckets = make([]*LinkedList, DefaultNumBuckets)
	} else {
		h.buckets = make([]*LinkedList, size)
	}
}

func (h *HashTable) calcHash(name interface{}) (int, error) {
	// This code converts a interface{} to a byte array
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(name)
	hash := 0
	// calculate the hash
	if err == nil {
		hasher := fnv.New32a()
		hasher.Write(buf.Bytes())
		hv := hasher.Sum32()
		hash = int(hv % uint32(len(h.buckets)))
	}
	// fmt.Printf("the hash for %v is %v, err is %v\n", name, hash, err)
	return hash, err
}

// Insert a name/value into the hash table.
func (h *HashTable) Insert(name interface{}, value interface{}) error {
	h.Lock()
	defer h.Unlock()
	index, err := h.calcHash(name)
	if err != nil {
		return err
	}
	bucketList := h.buckets[index]
	if bucketList == nil {
		newBucketList := &LinkedList{}
		bucket := &hashBucket{name: name, value: value}
		newBucketList.Insert(bucket)
		h.buckets[index] = newBucketList
	} else {
		var bucket *hashBucket
		iterator := bucketList.Iterator()
		for bucket := iterator(); bucket != nil; bucket = iterator() {
			if bucket.(*hashBucket).name == name {
				break
			}
		}
		if bucket != nil {
			bucketList.Delete(bucket)
			newBucket := &hashBucket{name: name, value: value}
			bucketList.Insert(newBucket)
		} else {
			newBucket := &hashBucket{name: name, value: value}
			bucketList.Insert(newBucket)
		}
	}
	return err
}

// Find a value in the hash table giving a key.
func (h *HashTable) Find(name interface{}) (interface{}, error) {
	h.RLock()
	defer h.RUnlock()
	index, err := h.calcHash(name)
	if err != nil {
		return nil, err
	}
	var rvalue interface{}
	rvalue = nil
	if h.buckets[index] != nil {
		bucketList := h.buckets[index]
		iterator := bucketList.Iterator()
		for bucket := iterator(); bucket != nil; bucket = iterator() {
			if bucket.(*hashBucket).name == name {
				rvalue = bucket.(*hashBucket).value
				break
			}
		}
	}
	return rvalue, err
}

// Delete an entry in the hash table.
func (h *HashTable) Delete(name interface{}) error {
	h.Lock()
	defer h.Unlock()
	index, err := h.calcHash(name)
	if err != nil {
		return err
	}
	bucketList := h.buckets[index]
	if bucketList != nil {
		iterator := bucketList.Iterator()
		bucket := iterator()
		for ; bucket != nil; bucket = iterator() {
			if bucket.(*hashBucket).name == name {
				// horrible hack because of threading problem that needs to be fixed!
				for iterator() != nil {
					// when the iterator is exhaused it gives up the read lock on
					// the linked list, if this does not happen we get a deadlock
					// in the delete from the linked list below.
				}
			}
		}
		if bucket != nil {
			bucketList.Delete(bucket)
		}
	}
	return err
}
