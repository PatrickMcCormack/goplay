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

type tree struct {
  value string
  left, right * tree
}

// current, left recursive, right recusive
func preorder(t * tree) {
  if t == nil {
    return
  }
//  fmt.Printf("In PreOrder, t = %v\n", t)
  s := t.value
  fmt.Println(s)
  preorder(t.left)
  preorder(t.right)
}

func bfs(t *tree) {
  
}

// left recursive, current, right recusive
func inorder(t * tree) {
  if t == nil {
    return
  }
  inorder(t.left)
  s := t.value
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
  s := t.value
  fmt.Println(s)
}

func newNode(t *tree, insertLeft bool, value string) *tree {
  nt := new(tree)
  nt.value = value
  if insertLeft == true {
    t.left = nt
  } else {
    t.right = nt
  }
  return nt
}

func main() {

  root  := &tree{"F", nil, nil}
  nodeb := &tree{"B", nil, nil}
  nodea := &tree{"A", nil, nil}
  noded := &tree{"D", nil, nil}
  nodec := &tree{"C", nil, nil}
  nodee := &tree{"E", nil, nil}
  nodeg := &tree{"G", nil, nil}
  nodei := &tree{"I", nil, nil}
  nodeh := &tree{"H", nil, nil}

  root.left = nodeb
  nodeb.left = nodea
  nodeb.right = noded
  noded.left = nodec
  noded.right = nodee
  root.right = nodeg
  nodeg.right = nodei
  nodei.left = nodeh

  fmt.Println("\nPreorder")
  preorder(root)
  fmt.Println("")

  fmt.Println("InOrder")
  inorder(root)
  fmt.Println("")

  fmt.Println("PostOrder")
  postorder(root)

}
