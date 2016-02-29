package collections

// Comparator is a type defintion for comparators used for collections
type Comparator func(interface{}, interface{}) int

// StringComparator is a helper function to compare two strings.
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
