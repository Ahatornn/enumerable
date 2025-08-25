package comparer

// ComparerString is a predefined ComparerFunc for comparing two string values.
// It performs a lexicographic comparison between two strings and returns:
//
//	-1 if the first string is lexicographically less than the second
//	 0 if both strings are equal
//	+1 if the first string is lexicographically greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerString ComparerFunc[string] = func(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
