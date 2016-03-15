package collections

// ComparatorResult encodes the result of a Comparator, see associated consts
type ComparatorResult int

const (
	// LessThan - first parameter is less than the second parameter
	LessThan ComparatorResult = -1
	// Equal - first parameter is equal to the second parameter
	Equal ComparatorResult = 0
	// GreaterThan - first parameter is greater than  the second parameter
	GreaterThan ComparatorResult = 1
)

// Comparator functions are used in collections where < and > are needed.
// Equality of interface{} is handled just fine by Go 
type Comparator func(interface{}, interface{}) ComparatorResult

// StringComparator is a helper function to compare two strings.
func StringComparator(v1 interface{}, v2 interface{}) ComparatorResult {
	// use same comparator semanitcs as Java
	if v1.(string) < v2.(string) {
		return LessThan
	} else if v1.(string) > v2.(string) {
		return GreaterThan
	} else {
		return Equal
	}
}

// IntComparator is a helper function to compare two integers.
func IntComparator(v1 interface{}, v2 interface{}) ComparatorResult {
	// use same comparator semanitcs as Java
	if v1.(int) < v2.(int) {
		return LessThan
	} else if v1.(int) > v2.(int) {
		return GreaterThan
	} else {
		return Equal
	}
}
