package main

// Simple Binary Tree
// PatrickMcCormack

// Trees can be traversed in pre-order, in-order, or post-order.
// These searches are referred to as depth-first search (DFS),
// as the search tree is deepened as much as possible on each child
// before going to the next sibling.

import (
  "fmt"
)

// the key is also the value in this example, dead easy to add a value if needed
type tree struct {
  key interface{}
  left, right * tree
}

type nodeCompare func(interface{}, interface{}) int

// current, left recursive, right recusive
func (t *tree) preorder() {
  if t == nil {
    return
  }
  s := t.key
  fmt.Println(s)
  t.left.preorder()
  t.right.preorder()
}

// left recursive, current, right recusive
func (t * tree) inorder() {
  if t == nil {
    return
  }
  t.left.inorder()
  s := t.key
  fmt.Println(s)
  t.right.inorder()
}

// left recursive, right recusive, current
func (t *tree) postorder() {
  if t == nil {
    return
  }
  t.left.postorder()
  t.right.postorder()
  s := t.key
  fmt.Println(s)
}

func compareString(v1 interface{}, v2 interface{}) int {
    // use same comparator semanitcs as Java
    if v1.(string) < v2.(string) {
      return -1
    } else if v1.(string) > v2.(string) {
      return 1
    } else {
      return 0
    }
}

func (node *tree) insertNode(insertkey interface{}, compare nodeCompare) *tree {
  if node == nil {
    node = &tree{insertkey, nil, nil}
  } else if compare(insertkey, node.key) == -1 {
      node.left = node.left.insertNode(insertkey, compare)
  } else { // key >= tree->key
      node.right = node.right.insertNode(insertkey, compare)
  }
  return node
}

func main() {

  root  := &tree{"F", nil, nil}
  root.insertNode("B", compareString)
  root.insertNode("A", compareString)
  root.insertNode("D", compareString)
  root.insertNode("C", compareString)
  root.insertNode("E", compareString)
  root.insertNode("G", compareString)
  root.insertNode("I", compareString)
  root.insertNode("H", compareString)

  fmt.Println("\nPreorder")
  root.preorder()
  fmt.Println("")

  fmt.Println("InOrder")
  root.inorder()
  fmt.Println("")

  fmt.Println("PostOrder")
  root.postorder()

}
