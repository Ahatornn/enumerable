package comparer

// ComparerRune is a predefined ComparerFunc for comparing two rune values.
// It performs a numerical comparison between two Unicode code points and returns:
//
//	-1 if the first rune is less than the second
//	 0 if both runes are equal
//	+1 if the first rune is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerRune ComparerFunc[rune] = func(a, b rune) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
