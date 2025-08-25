package comparer

// ComparerFloat32 is a predefined ComparerFunc for comparing two float32 values.
// It performs a numerical comparison between two float32 values and returns:
//
//	-1 if the first value is less than the second
//	 0 if both values are equal
//	+1 if the first value is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
//
// Note: This comparison does not handle NaN values specially - NaN comparisons
// follow Go's built-in comparison rules where NaN is not equal to anything,
// including itself.
var ComparerFloat32 ComparerFunc[float32] = func(a, b float32) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
