package bst

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreateBST(t *testing.T) {

	// fixme - create multiple test scenarios
	// which compare to golden data - needs
	// inorder etc. to return a data structure
	// or better a iterator

	t.Log("Testing the creation of BST")

	tree := &Tree{root: &Node{key: "F"}}
	tree.InsertNode("B", StringComparator)
	tree.InsertNode("A", StringComparator)
	tree.InsertNode("D", StringComparator)
	tree.InsertNode("C", StringComparator)
	tree.InsertNode("E", StringComparator)
	tree.InsertNode("G", StringComparator)
	tree.InsertNode("I", StringComparator)
	tree.InsertNode("H", StringComparator)

	fmt.Println("\nPreorder")
	tree.Preorder()
	fmt.Println("")

	fmt.Println("\nInorder")
	tree.Inorder()
	fmt.Println("")

	fmt.Println("\nPostorder")
	tree.Postorder()
	fmt.Println("")

	rand.Seed(time.Now().Unix())

	tree2 := &Tree{root: &Node{key: 5}}
	for i := 0; i < 20; i++ {
		tree2.InsertNode(rand.Intn(1000), IntComparator)
	}

	fmt.Println("Inorder")
	tree2.Inorder()
	fmt.Println("")

	tree3 := &Tree{root: &Node{key: "The"}}
	tree3.InsertNode("quick", StringComparator)
	tree3.InsertNode("fox", StringComparator)
	tree3.InsertNode("jumped", StringComparator)
	tree3.InsertNode("over", StringComparator)
	tree3.InsertNode("the", StringComparator)
	tree3.InsertNode("lazy", StringComparator)
	tree3.InsertNode("Dog", StringComparator)

	fmt.Println("Inorder")
	tree3.Inorder()
	fmt.Println("")

	t.Log("Finished Testing the creation of BST")
}
