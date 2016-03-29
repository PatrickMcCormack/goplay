package collections

// Trie - really crude, initial, incomplete, no-functioning implementation of
// a Trie, lots to do to make this respectable (correct handling of char, runes, etc.)
// Patrick McCormack

import (
	"errors"
	"fmt"
	s "strings"
)

const (
	// AlphabetSize ...
	AlphabetSize = 26
	CharOffset   = 65
)

type TrieNodes [AlphabetSize]*TrieNode

// TrieNode ...
type TrieNode struct {
	// Lets start with the basic English alphabet and then support 8859-1
	// The existanace of a node at a position indicates that character
	// exists.
	trieNodes TrieNodes
}

// Trie ...
type Trie struct {
	root TrieNode
}

// Insert ...
func (trie *Trie) Insert(word string) error {
	var err error
	word = s.ToUpper(word)
	node := &trie.root
	for i := 0; i < len(word); i++ {
		index := word[i] - CharOffset
		if index < 0 || index > 25 {
			err = errors.New("Insert into Trie failed because char is out of range")
			break
		}
		// If the node does not exist create and keep going
		if node.trieNodes[index] == nil {
			node.trieNodes[index] = &TrieNode{}
		}
		// Keep on truckin
		node = node.trieNodes[index]
	}
	return err
}

// Find ...
func (trie *Trie) Exists(word string) error {
	var err error
	word = s.ToUpper(word)
	node := &trie.root
	for i := 0; i < len(word); i++ {
		index := word[i] - CharOffset
		fmt.Printf("%v", string(index+CharOffset))
		if index < 0 || index > 25 {
			err = errors.New("Insert into Trie failed because char is out of range")
			break
		}
		// If the node does not exist create and keep going
		if node.trieNodes[index] == nil {
			err = errors.New("Not found")
			break
		} else {
			// Keep on truckin
			//      fmt.Printf("found %v and index value is %v\n", index, node.trieNodes[index])
			node = node.trieNodes[index]
		}
	}
	return err
}

// FindAll ... print all possibiliies from this node and letter
func (trie *Trie) FindAll() {
	findAll(&trie.root.trieNodes, "", 0)
}

func findAll(nodes *TrieNodes, word string, level int) {
	for i, c := range *nodes {
		if c != nil {
			//fmt.Printf("Level %v, Char=%v\n",level, string(i+CharOffset))
			findAll(&c.trieNodes, word+string(i+CharOffset), level+1)
			// how to not get every version of the word as the stack unwinds
			fmt.Println(word)
		}
	}
}
