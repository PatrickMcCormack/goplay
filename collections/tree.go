package collections

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
	root         *TreeNode
	sync.RWMutex // composite object
}

type TreeNode struct {
	key         interface{} // key is also the value simple to add a seperate value if needed
	left, right *TreeNode
}

type TreeNodeCompare func(interface{}, interface{}) int

func (tree *Tree) Preorder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.Preorder()
}

func (node *TreeNode) Preorder() {
	if node == nil {
		return
	}
	fmt.Println(node.key)
	node.left.Preorder()
	node.right.Preorder()
}

func (tree *Tree) NonRecursivePreOrder() {
	tree.RLock()
	defer tree.RUnlock()
	var stack Stack
	stack.push(tree.root)
	for stack.size() > 0 {
		node := stack.pop().(*TreeNode)
		fmt.Println(node.key)
		if node.right != nil {
			stack.push(node.right)
		}
		if node.left != nil {
			stack.push(node.left)
		}
	}
}

func (tree *Tree) ClosureBasedNonRecursivePreOrder() func() interface{} {
	var stack Stack
	tree.RLock()
	stack.push(tree.root)
	return func() interface{} {
		if stack.size() <= 0 {
			tree.RUnlock()
			return nil
		}
		node := stack.pop().(*TreeNode)
		if node.right != nil {
			stack.push(node.right)
		}
		if node.left != nil {
			stack.push(node.left)
		}
		return node.key
	}
}

func (tree *Tree) Inorder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.Inorder()
}

func (node *TreeNode) Inorder() {
	if node == nil {
		return
	}
	node.left.Inorder()
	fmt.Println(node.key)
	node.right.Inorder()
}

func (tree *Tree) Postorder() {
	tree.RLock()
	defer tree.RUnlock()
	tree.root.Postorder()
}

func (node *TreeNode) Postorder() {
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

func (tree *Tree) InsertNode(insertkey interface{}, compare TreeNodeCompare) *TreeNode {
	tree.Lock()
	defer tree.Unlock()
	return tree.root.InsertNode(insertkey, compare)
}

func (node *TreeNode) InsertNode(insertkey interface{}, compare TreeNodeCompare) *TreeNode {
	if node == nil {
		node = &TreeNode{key: insertkey}
	} else if compare(insertkey, node.key) == -1 {
		node.left = node.left.InsertNode(insertkey, compare)
	} else { // key >= TreeNode->key
		node.right = node.right.InsertNode(insertkey, compare)
	}
	return node
}
