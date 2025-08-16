package enumerable

// Count returns the number of elements in an enumeration.
// This operation is useful for determining the size of a sequence.
//
// The count operation will:
// - Iterate through all elements in the enumeration
// - Count each element encountered
// - Return the total count
// - Handle nil enumerators gracefully
//
// Returns:
//   The number of elements in the enumeration (0 for empty or nil enumerations)
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to count all elements. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
// - If the enumerator is nil, returns 0
// - If the enumeration is empty, returns 0
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - All elements are processed, even if consumer wants to stop early
func (q Enumerator[T]) Count() int {
	return countInternal(q)
}

// Count returns the number of elements in an enumeration.
// This operation is useful for determining the size of a sequence.
//
// The count operation will:
// - Iterate through all elements in the enumeration
// - Count each element encountered
// - Return the total count
// - Handle nil enumerators gracefully
//
// Returns:
//   The number of elements in the enumeration (0 for empty or nil enumerations)
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to count all elements. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
// - If the enumerator is nil, returns 0
// - If the enumeration is empty, returns 0
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - All elements are processed, even if consumer wants to stop early
func (q AnyEnumerator[T]) Count() int {
	return countInternal(q)
}

func countInternal[T any](enumerator func(func(T) bool)) int {
	if enumerator == nil {
		return 0
	}
	var cnt int
	enumerator(func(T) bool {
		cnt++
		return true
	})
	return cnt
}
