package collections

// Very rough WIP, nothing to see here, move on ...

import (
	"hash/fnv"
)

type hashElement struct {
	name  string
	value string
}

type HashTable struct {
	htable []*hashElement
	size   int
}

func (h *HashTable) Initialize() {
	h.htable = make([]*hashElement, 100)
}

// Insert
func (h *HashTable) Insert(name string, value string) {
	hasher := fnv.New64a()
	hasher.Write([]byte(name))
	hv := hasher.Sum64()
	index := hv % 100
	h.htable[index] = &hashElement{name, value}
	h.size++
}

// Find
func (h *HashTable) Find(name string) (bool, string) {
	hasher := fnv.New64a()
	hasher.Write([]byte(name))
	hv := hasher.Sum64()
	index := hv % 100
	var rvalue string
	var status bool
	if h.htable[index] != nil {
		rvalue = h.htable[index].value
		status = true
	} else {
		rvalue = ""
		status = false
	}
	return status, rvalue
}

// Delete
func (h *HashTable) Delete(name string) (bool, string) {
	h.size--
	return false, ""
}
