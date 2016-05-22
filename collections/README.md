This is a simple set of thread-safe go collections that I started writing 
to learn Go, this is a work in progress.
  1. Stack
  2. Queue
  3. Single Linked List
  4. Double Linked List
  5. Binary Tree
  6. HashTable  
  7. Heap - ~50% complete
  8. Trie - just stared  
  9. K-d Tree - future
 10. Balanced Tree (red/black or AVL) - future

It's interesting to see how many language features are in the binary tree code which is a relatively small amount of code:

- user defined types
- type inference
- type assertion
- structs
- composition
- struct literal initialization of explicit elements
- pointers
- locking (read/write locks)
- deferred statement execution
- empty interfaces
- methods
- package visibility
- go testing framework
- anonymous functions
- closures
- iota
- go errors
- go docs, try "go doc" or "go doc ClosureBasedNonRecursivePreOrder"
