package comparer

// ComparerByte is a predefined ComparerFunc for comparing two byte values.
// It performs a numerical comparison between two unsigned 8-bit integers and returns:
//
//	-1 if the first byte is less than the second
//	 0 if both bytes are equal
//	+1 if the first byte is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerByte ComparerFunc[byte] = func(a, b byte) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
