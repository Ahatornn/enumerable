package enumerable

// Range generates a sequence of consecutive integers starting at 'start',
// producing exactly 'count' values in ascending order (with step +1).
//
// Parameters:
//   start - initial value of the sequence (inclusive)
//   count - number of values to generate (must be non-negative)
//
// Returns:
//   An Enumerator[int] that can be used in range loops.
//
// Notes:
// - For count = 0, produces an empty sequence (no iterations)
// - For count < 0, behavior is undefined (should be avoided)
func Range(start, count int) Enumerator[int] {
	return func(yield func(int) bool) {
		for i := 0; i < count; i++ {
			if !yield(start + i) {
				return
			}
		}
	}
}

// RangeAny generates a sequence of consecutive integers starting at 'start',
// producing exactly 'count' values in ascending order (with step +1).
//
// Parameters:
//   start - initial value of the sequence (inclusive)
//   count - number of values to generate (must be non-negative)
//
// Returns:
//   An AnyEnumerator[int] that can be used in range loops.
//
// Notes:
// - For count = 0, produces an empty sequence (no iterations)
// - For count < 0, behavior is undefined (should be avoided)
func RangeAny(start, count int) AnyEnumerator[int] {
	return func(yield func(int) bool) {
		for i := 0; i < count; i++ {
			if !yield(start + i) {
				return
			}
		}
	}
}
