package bst

// Simple Binary Tree
// PatrickMcCormack

// Trees can be traversed in pre-order, in-order, or post-order.
// These searches are referred to as depth-first search (DFS),
// as the search tree is deepened as much as possible on each child
// before going to the next sibling.

import (
	"fmt"
	"sync"
)

type Tree struct {
	root         *Node
	sync.RWMutex // composite object
}

// the key is also the value in this example, dead easy to add a value if needed
type Node struct {
	key         interface{}
	left, right *Node
}

type NodeCompare func(interface{}, interface{}) int

// current, left recursive, right recusive
func (tree *Tree) Preorder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.Preorder()
}

// current, left recursive, right recusive
func (node *Node) Preorder() {
	if node == nil {
		return
	}
	fmt.Println(node.key)
	node.left.Preorder()
	node.right.Preorder()
}

// left recursive, current, right recusive
func (tree *Tree) Inorder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.Inorder()
}

// left recursive, current, right recusive
func (node *Node) Inorder() {
	if node == nil {
		return
	}
	node.left.Inorder()
	fmt.Println(node.key)
	node.right.Inorder()
}

// left recursive, right recusive, current
func (tree *Tree) Postorder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.Postorder()
}

// left recursive, right recusive, current
func (node *Node) Postorder() {
	if node == nil {
		return
	}
	node.left.Postorder()
	node.right.Postorder()
	fmt.Println(node.key)
}

func StringComparator(v1 interface{}, v2 interface{}) int {
	// use same comparator semanitcs as Java
	if v1.(string) < v2.(string) {
		return -1
	} else if v1.(string) > v2.(string) {
		return 1
	} else {
		return 0
	}
}

func IntComparator(v1 interface{}, v2 interface{}) int {
	// use same comparator semanitcs as Java
	if v1.(int) < v2.(int) {
		return -1
	} else if v1.(int) > v2.(int) {
		return 1
	} else {
		return 0
	}
}

func (tree *Tree) InsertNode(insertkey interface{}, compare NodeCompare) *Node {
	tree.Lock()
	defer tree.Unlock()
	return tree.root.InsertNode(insertkey, compare)
}

func (node *Node) InsertNode(insertkey interface{}, compare NodeCompare) *Node {
	if node == nil {
		node = &Node{key: insertkey}
	} else if compare(insertkey, node.key) == -1 {
		node.left = node.left.InsertNode(insertkey, compare)
	} else { // key >= node->key
		node.right = node.right.InsertNode(insertkey, compare)
	}
	return node
}
