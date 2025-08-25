package comparer

import "time"

// ComparerTime is a predefined ComparerFunc for comparing two time.Time values.
// It performs a chronological comparison between two time values and returns:
//
//	-1 if the first time is before the second
//	 0 if both times are equal
//	+1 if the first time is after the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
//
// Note: This comparison uses time.Time's built-in Before() and After() methods
// for accurate chronological ordering.
var ComparerTime ComparerFunc[time.Time] = func(a, b time.Time) int {
	if a.Before(b) {
		return -1
	} else if a.After(b) {
		return 1
	}
	return 0
}
