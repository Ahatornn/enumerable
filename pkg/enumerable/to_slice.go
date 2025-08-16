package enumerable

// ToSlice converts an enumeration to a slice containing all elements.
// This operation is useful for materializing lazy enumerations into concrete collections.
//
// The to slice operation will:
// - Iterate through all elements in the enumeration
// - Collect all elements into a new slice
// - Return the slice containing all elements in order
// - Handle nil enumerators gracefully
//
// Returns:
//   A slice containing all elements from the enumeration, or empty slice if enumerator is nil
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to collect all elements. For large
// enumerations, this may be expensive in both time and memory.
//
// ⚠️ Memory note: This operation buffers all elements in memory.
//
// Notes:
// - If the enumerator is nil, returns an empty slice (not nil)
// - If the enumeration is empty, returns an empty slice
// - For very large enumerations, consider processing elements incrementally.
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(n) - allocates memory for all elements
// - The enumeration stops only when exhausted or if upstream operations stop it
// - Returned slice preserves the order of elements from the enumeration
// - Use with caution for infinite or very large enumerations
func (q Enumerator[T]) ToSlice() []T {
	return toSliceInternal(q)
}

// ToSlice converts an enumeration to a slice containing all elements.
// This operation is useful for materializing lazy enumerations into concrete collections.
//
// The to slice operation will:
// - Iterate through all elements in the enumeration
// - Collect all elements into a new slice
// - Return the slice containing all elements in order
// - Handle nil enumerators gracefully
//
// Returns:
//   A slice containing all elements from the enumeration, or empty slice if enumerator is nil
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to collect all elements. For large
// enumerations, this may be expensive in both time and memory.
//
// ⚠️ Memory note: This operation buffers all elements in memory.
//
// Notes:
// - If the enumerator is nil, returns an empty slice (not nil)
// - If the enumeration is empty, returns an empty slice
// - For very large enumerations, consider processing elements incrementally.
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(n) - allocates memory for all elements
// - The enumeration stops only when exhausted or if upstream operations stop it
// - Returned slice preserves the order of elements from the enumeration
// - Use with caution for infinite or very large enumerations
func (q AnyEnumerator[T]) ToSlice() []T {
	return toSliceInternal(q)
}

func toSliceInternal[T any](enumerator func(func(T) bool)) []T {
	if enumerator == nil {
		return []T{}
	}
	result := make([]T, 0, 16)
	enumerator(func(item T) bool {
		result = append(result, item)
		return true
	})
	return result
}
