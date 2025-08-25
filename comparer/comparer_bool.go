package comparer

// ComparerBool is a predefined ComparerFunc for comparing two bool values.
// It performs a logical comparison where false is considered less than true and returns:
//
//	-1 if the first boolean is false and the second is true
//	 0 if both booleans are equal
//	+1 if the first boolean is true and the second is false
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerBool ComparerFunc[bool] = func(a, b bool) int {
	if !a && b {
		return -1
	} else if a && !b {
		return 1
	}
	return 0
}
