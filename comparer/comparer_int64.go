package comparer

// ComparerInt64 is a predefined ComparerFunc for comparing two int64 values.
// It performs a natural numeric comparison between two int64 integers and returns:
//
//	-1 if the first integer is less than the second
//	 0 if both integers are equal
//	+1 if the first integer is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerInt64 ComparerFunc[int64] = func(a, b int64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
