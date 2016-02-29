package collections

// Comparator is a type defintion for comparators used to
// insert nodes into a binary tree.
type Comparator func(interface{}, interface{}) int

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
