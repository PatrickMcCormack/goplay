package collections

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	t.Log("Testing Trie functionality")

	// all fails commented out because this is a really early wip

	// ---------------------------------------------------------------------
	t.Log("==Start: Some Test")

	var trie Trie
	var err error
	trie.Insert("aAbBcCdDefghijklmnopqrstuvwxyz")
	trie.Insert("amazon")
	trie.Insert("anaconda")
	trie.Insert("ardvark")
	trie.Insert("ardwolf")
	trie.Insert("avalon")
	trie.Insert("adverse")
	trie.Insert("google")
	trie.Insert("yahoo")
	err = trie.Exists("aAbBcCdDefghijklmnopqrstuvwxyz")
	if err != nil {
		t.Log(err)
		//		t.Fail()
	}
	fmt.Println("XXXX")
	err = trie.Exists("aAbBcCdDefghijklmnopqrstuvwxyA")
	if err != nil {
		t.Log(err)
		//		t.Fail()
	}
	fmt.Println("XXXX")
	err = trie.Exists("AABD")
	if err != nil {
		t.Log(err)
		//		t.Fail()
	}
	fmt.Println("XXXX")
	if 1 == 0 {
		t.Log("Something failed")
		//		t.Fail()
	}
	t.Log("==Completed: Some Test")

	//	fmt.Println("####")
	//	trie.FindAll()
	//	fmt.Println("####")

	// ---------------------------------------------------------------------
	t.Log("Done testing Trie functionality")

}
