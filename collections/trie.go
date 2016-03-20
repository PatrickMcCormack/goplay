package collections

import "fmt"

const (
	// AlphabetSize ...
	AlphabetSize = 26
)

// TrieNode ...
type TrieNode struct {
	// Lets start with the basic English alphabet and then support 8859-1
	trieNodes [AlphabetSize]*TrieNode
}

// Trie ...
type Trie struct {
	nodes TrieNode
}

// Insert ...
func (t *Trie) Insert(s string) {
	fmt.Println(s)
}
