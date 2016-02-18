package collections

// Tree - a thread-safe tree implementation
// PatrickMcCormack

// Trees can be traversed in pre-order, in-order, or post-order.
// These searches are referred to as depth-first search (DFS),
// as the search tree is deepened as much as possible on each child
// before going to the next sibling.

import (
	"fmt"
	"sync"
)

// Tree is a container for binary trees.
type Tree struct {
	root         *TreeNode
	sync.RWMutex // composite object
}

// TreeNode represents an element in a tree.
type TreeNode struct {
	key         interface{} // key is also the value simple to add a seperate value if needed
	left, right *TreeNode
}

// TreeNodeCompare is a type defintion for comparators used to
// insert nodes into a binary tree.
type TreeNodeCompare func(interface{}, interface{}) int

// PreOrder is a thread-safe method that prints a binary tree in pre-order
func (tree *Tree) PreOrder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.preorder()
}

// preorder is a recursive helper method that prints a binary tree in pre-order
func (node *TreeNode) preorder() {
	if node == nil {
		return
	}
	fmt.Println(node.key)
	node.left.preorder()
	node.right.preorder()
}

// NonRecursivePreOrder prints a tree elements in pre-order.
func (tree *Tree) NonRecursivePreOrder() {
	tree.RLock()
	defer tree.RUnlock()
	var stack Stack
	stack.Push(tree.root)
	for stack.Size() > 0 {
		node := stack.Pop().(*TreeNode)
		fmt.Println(node.key)
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}
}

// ClosureBasedNonRecursivePreOrder returns an iterator over a binary
// tree. Calling the iterator returns the tree elements in pre-order.
func (tree *Tree) ClosureBasedNonRecursivePreOrder() func() interface{} {
	var stack Stack
	tree.RLock()
	stack.Push(tree.root)
	// return a closure over the variables in scope
	return func() interface{} {
		if stack.Size() <= 0 {
			tree.RUnlock()
			return nil
		}
		node := stack.Pop().(*TreeNode)
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
		return node.key
	}
}

// InOrder is a thread-safe method that prints a binary tree in in-order
func (tree *Tree) InOrder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.inorder()
}

// inorder is a recursive helper method that prints a binary tree in in-order
func (node *TreeNode) inorder() {
	if node == nil {
		return
	}
	node.left.inorder()
	fmt.Println(node.key)
	node.right.inorder()
}

// PostOrder is a thread-safe method that prints a binary tree in post-order
func (tree *Tree) PostOrder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.postorder()
}

// postorder is a recursive helper method that prints a binary tree in post-order
func (node *TreeNode) postorder() {
	if node == nil {
		return
	}
	node.left.postorder()
	node.right.postorder()
	fmt.Println(node.key)
}

// StringComparator is a helper function to compare two strings.
// This is used to insert string elements into a binary tree.
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

// IntComparator is a helper function to compare two integers.
// This is used to insert integer elements into a binary tree.
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

// InsertNode takes a key and a comparator helper function and inserts
// the value into the tree. In this implementation the key and the tree
// element value are the same.
func (tree *Tree) InsertNode(insertkey interface{}, compare TreeNodeCompare) *TreeNode {
	tree.Lock()
	defer tree.Unlock()
	return tree.root.insertnode(insertkey, compare)
}

// insertnode is a recursive helper function to insert a node in a tree
func (node *TreeNode) insertnode(insertkey interface{}, compare TreeNodeCompare) *TreeNode {
	if node == nil {
		node = &TreeNode{key: insertkey}
	} else if compare(insertkey, node.key) == -1 {
		node.left = node.left.insertnode(insertkey, compare)
	} else { // key >= TreeNode->key
		node.right = node.right.insertnode(insertkey, compare)
	}
	return node
}
