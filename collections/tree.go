package collections

// Tree - a thread-safe tree implementation
// PatrickMcCormack

// A BST works by ensuring the following conditons:
// 1. The left branch of the tree is always < it's parent
// 2/ The right branch is always <= to it's parent

// Todo - add Delete

import (
	"fmt"
	"sync"
)

// Tree is a container for binary trees.
type Tree struct {
	root         *TreeNode
	comparator   Comparator
	treeSize     int
	sync.RWMutex // composite object
}

// TreeNode represents an element in a tree.
type TreeNode struct {
	key         interface{} // key is also the value simple to add a seperate value if needed
	left, right *TreeNode
}

// IteratorOrder is a type that represents the 3 ways to iterate over a BST.
type IteratorOrder int

const (
	// PreOrder when passed to Interator(...) returns a pre-order iterator
	PreOrder IteratorOrder = iota // 1 (i.e. 1 << 0)
	// InOrder when passed to Interator(...) returns an in-order iterator
	InOrder // 2 (i.e. 1 << 1)
	// PostOrder when passed to Interator(...) returns a post-order iterator
	PostOrder // 4 (i.e 1 << 2)
)

// Insert takes a key and inserts it into the tree.
// Note: In this implementation the key and the tree element value are the same.
func (tree *Tree) Insert(insertkey interface{}) *TreeNode {
	tree.Lock()
	defer tree.Unlock()
	var rval *TreeNode
	if tree.root == nil {
		tree.root = &TreeNode{key: insertkey}
		rval = tree.root
	} else {
		rval = tree.root.insert(insertkey, tree.comparator)
	}
	tree.treeSize++
	return rval
}

// insert is a recursive helper function to insert a node in a tree
func (node *TreeNode) insert(insertkey interface{}, comparator Comparator) *TreeNode {
	if node == nil {
		node = &TreeNode{key: insertkey}
	} else if comparator(insertkey, node.key) == LessThan {
		node.left = node.left.insert(insertkey, comparator)
	} else { // key >= TreeNode->key
		node.right = node.right.insert(insertkey, comparator)
	}
	return node
}

// Size returns the number of elements in the tree
func (tree *Tree) Size() int {
	tree.RLock()
	defer tree.RUnlock()
	return tree.treeSize
}

// Depth returns the maximum depth fo the tree.
func (tree *Tree) Depth() int {
	tree.RLock()
	defer tree.RUnlock()
	return tree.root.depth()
}

func (node *TreeNode) depth() int {
	if node == nil {
		return 0
	}
	leftdepth := node.left.depth()
	rightdepth := node.right.depth()
	if leftdepth > rightdepth {
		return leftdepth + 1
	}
	return rightdepth + 1
}

// BreadthFirst populates a linkedlist with a breadth first list of the
// BST elements.
// Algorithm adapted from Wikipedia https://en.wikipedia.org/wiki/Tree_traversal
func (tree *Tree) BreadthFirst(list *LinkedList) {
	tree.RLock()
	defer tree.RUnlock()
	var queue Queue
	queue.Push(tree.root)
	for queue.Size() > 0 {
		node := queue.Pop().(*TreeNode)
		list.Insert(node.key)
		if node.left != nil {
			queue.Push(node.left)
		}
		if node.right != nil {
			queue.Push(node.right)
		}
	}
}

// BreadthFirstTraversalWithLevels prints a breadth first version of the tree
// with the nodes for each level occuping the same line.
// Fixme - not a print but make this return a data struture that preserves
// the levels - linked list of linked lists .
func (tree *Tree) BreadthFirstTraversalWithLevels() {
	tree.RLock()
	defer tree.RUnlock()
	var queue Queue
	// currentCount is the number of nodes to print on the current level
	// nextCount is the number of element to print on the next level
	currentCount, nextCount := 1, 0
	queue.Push(tree.root)
	for queue.Size() > 0 {
		node := queue.Pop().(*TreeNode)
		currentCount--
		fmt.Printf("%v ", node.key)
		if node.left != nil {
			nextCount++
			queue.Push(node.left)
		}
		if node.right != nil {
			nextCount++
			queue.Push(node.right)
		}
		if currentCount == 0 {
			fmt.Println("")
			currentCount, nextCount = nextCount, currentCount
		}
	}
}

// Iterator returns a closure that traverses a BST. This helper method can
// return a pre-order, in-order or post-order iterator.
func (tree *Tree) Iterator(order IteratorOrder) func() interface{} {
	var iterator func() interface{}
	switch {
	case order == PreOrder:
		iterator = tree.preOrderIterator()
	case order == InOrder:
		iterator = tree.inOrderIterator()
	case order == PostOrder:
		iterator = tree.postOrderIterator()
	}
	return iterator
}

// preOrderIterator prints a tree elements in pre-order.
// Algorithm adapted from Wikipedia https://en.wikipedia.org/wiki/Tree_traversal
func (tree *Tree) preOrderIterator() func() interface{} {
	tree.RLock()
	var stack Stack
	node := tree.root
	return func() interface{} {
		var returnValue interface{}
		for stack.Size() > 0 || node != nil {
			if node != nil {
				returnValue = node.key
				if node.right != nil {
					stack.Push(node.right)
				}
				node = node.left
				break
			} else {
				node = stack.Pop().(*TreeNode)
			}
		}
		if returnValue == nil {
			tree.RUnlock()
		}
		return returnValue
	}
}

// inOrderIterator prints a tree elements in-order.
// Algorithm adapted from https://en.wikipedia.org/wiki/Tree_traversal
func (tree *Tree) inOrderIterator() func() interface{} {
	tree.RLock()
	var stack Stack
	node := tree.root
	return func() interface{} {
		var returnValue interface{}
		for stack.Size() > 0 || node != nil {
			if node != nil {
				stack.Push(node)
				node = node.left
			} else {
				node = stack.Pop().(*TreeNode)
				returnValue = node.key
				node = node.right
				break
			}
		}
		if returnValue == nil {
			tree.RUnlock()
		}
		return returnValue
	}
}

// PostOrderIterator prints a tree elements in post-order.
// Algorithm adapted from https://en.wikipedia.org/wiki/Tree_traversal
func (tree *Tree) postOrderIterator() func() interface{} {
	tree.RLock()
	var stack Stack
	var lastNodeVisited *TreeNode
	node := tree.root
	return func() interface{} {
		var returnValue interface{}
		for stack.Size() > 0 || node != nil {
			if node != nil {
				stack.Push(node)
				node = node.left
			} else {
				peekNode := stack.Peek().(*TreeNode)
				// if right child exists and traversing node
				// from left child, then move right
				if peekNode.right != nil && lastNodeVisited != peekNode.right {
					node = peekNode.right
				} else {
					returnValue = peekNode.key
					lastNodeVisited = stack.Pop().(*TreeNode)
					break
				}
			}
		}
		if returnValue == nil {
			tree.RUnlock()
		}
		return returnValue
	}
}

// RecursiveTraversal is a helper function for calling the correct
// recursive traversal function
func (tree *Tree) RecursiveTraversal(order IteratorOrder, result *LinkedList) {
	switch {
	case order == PreOrder:
		tree.RecursivePreOrder(result)
	case order == InOrder:
		tree.RecursiveInOrder(result)
	case order == PostOrder:
		tree.RecursivePostOrder(result)
	}
}

// RecursivePreOrder is a thread-safe method that returns a BST in pre-order
func (tree *Tree) RecursivePreOrder(result *LinkedList) {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.recursivePreOrder(result)
}

// preorder is a recursive helper method that prints a binary tree in pre-order
func (node *TreeNode) recursivePreOrder(result *LinkedList) {
	if node == nil {
		return
	}
	result.Insert(node.key)
	node.left.recursivePreOrder(result)
	node.right.recursivePreOrder(result)
}

// RecursiveInOrder is a thread-safe method that returns a BST in in-order
func (tree *Tree) RecursiveInOrder(result *LinkedList) {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.recursiveInOrder(result)
}

// inorder is a recursive helper method that prints a binary tree in in-order
func (node *TreeNode) recursiveInOrder(result *LinkedList) {
	if node == nil {
		return
	}
	node.left.recursiveInOrder(result)
	result.Insert(node.key)
	node.right.recursiveInOrder(result)
}

// RecursivePostOrder is a thread-safe method that returns a BST in post-order
func (tree *Tree) RecursivePostOrder(result *LinkedList) {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.recursivePostOrder(result)
}

// postorder is a recursive helper method that prints a binary tree in post-order
func (node *TreeNode) recursivePostOrder(result *LinkedList) {
	if node == nil {
		return
	}
	node.left.recursivePostOrder(result)
	node.right.recursivePostOrder(result)
	result.Insert(node.key)
}
