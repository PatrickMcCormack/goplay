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
func preorder(t * tree) {
  if t == nil {
    return
  }
  s := t.key
  fmt.Println(s)
  preorder(t.left)
  preorder(t.right)
}

// left recursive, current, right recusive
func inorder(t * tree) {
  if t == nil {
    return
  }
  inorder(t.left)
  s := t.key
  fmt.Println(s)
  inorder(t.right)
}

// left recursive, right recusive, current
func postorder(t * tree) {
  if t == nil {
    return
  }
  postorder(t.left)
  postorder(t.right)
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
  preorder(root)
  fmt.Println("")

  fmt.Println("InOrder")
  inorder(root)
  fmt.Println("")

  fmt.Println("PostOrder")
  postorder(root)

}
