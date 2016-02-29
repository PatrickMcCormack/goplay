package collections

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCreateBST(t *testing.T) {

	// fixme - create multiple test scenarios
	// which compare to golden data - needs
	// inorder etc. to return a data structure
	// or better a iterator

	t.Log("Testing the creation of BST")

	tree := &Tree{root: &TreeNode{key: "F"}}
	tree.InsertNode("B", StringComparator)
	tree.InsertNode("A", StringComparator)
	tree.InsertNode("D", StringComparator)
	tree.InsertNode("C", StringComparator)
	tree.InsertNode("E", StringComparator)
	tree.InsertNode("G", StringComparator)
	tree.InsertNode("I", StringComparator)
	tree.InsertNode("H", StringComparator)

	fmt.Println("\nPreOrder")
	tree.PreOrder()
	fmt.Println("")

	fmt.Println("\nInOrder")
	tree.InOrder()
	fmt.Println("")

	fmt.Println("\nPostOrder")
	tree.PostOrder()
	fmt.Println("")

	rand.Seed(time.Now().Unix())

	tree2 := &Tree{root: &TreeNode{key: 5}}
	for i := 0; i < 20; i++ {
		tree2.InsertNode(rand.Intn(1000), IntComparator)
	}

	fmt.Println("InOrder")
	tree2.InOrder()
	fmt.Println("")

	tree3 := &Tree{root: &TreeNode{key: "The"}}
	tree3.InsertNode("quick", StringComparator)
	tree3.InsertNode("fox", StringComparator)
	tree3.InsertNode("jumped", StringComparator)
	tree3.InsertNode("over", StringComparator)
	tree3.InsertNode("the", StringComparator)
	tree3.InsertNode("lazy", StringComparator)
	tree3.InsertNode("Dog", StringComparator)

	fmt.Println("PreOrder")
	tree3.PreOrder()
	fmt.Println("")

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Pre-order - NonRecursivePreOrder %v\n", i)
			tree3.NonRecursivePreOrder()
		}()
	}
	wg.Wait()
	fmt.Println("")

	fmt.Println("Pre-order - ClosureBasedNonRecursivePreOrder")
	iterator := tree3.ClosureBasedNonRecursivePreOrder()
	for v := iterator(); v != nil; v = iterator() {
		fmt.Println(v)
	}
	fmt.Println("")

	t.Log("Finished Testing the creation of BST")
}
