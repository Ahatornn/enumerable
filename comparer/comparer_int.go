package comparer

// ComparerInt is a predefined ComparerFunc for comparing two int values.
// It performs a natural numeric comparison between two integers and returns:
//
//	-1 if the first integer is less than the second
//	 0 if both integers are equal
//	+1 if the first integer is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerInt ComparerFunc[int] = func(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
