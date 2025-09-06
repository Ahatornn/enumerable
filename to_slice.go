package enumerable

// ToSlice converts an enumeration to a slice containing all elements.
// This operation is useful for materializing lazy enumerations into concrete collections.
//
// The to slice operation will:
//   - Iterate through all elements in the enumeration
//   - Collect all elements into a new slice
//   - Return the slice containing all elements in order
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A slice containing all elements from the enumeration, or empty slice if enumerator is nil
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to collect all elements. For large
// enumerations, this may be expensive in both time and memory.
//
// ⚠️ Memory note: This operation buffers all elements in memory.
//
// Notes:
//   - If the enumerator is nil, returns an empty slice (not nil)
//   - If the enumeration is empty, returns an empty slice
//   - For very large enumerations, consider processing elements incrementally.
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(n) - allocates memory for all elements
//   - The enumeration stops only when exhausted or if upstream operations stop it
//   - Returned slice preserves the order of elements from the enumeration
//   - Use with caution for infinite or very large enumerations
func (q Enumerator[T]) ToSlice() []T {
	return toSliceInternal(q)
}

// ToSlice converts an enumeration to a slice containing all elements.
// This operation is useful for materializing lazy enumerations into concrete collections.
//
// The to slice operation will:
//   - Iterate through all elements in the enumeration
//   - Collect all elements into a new slice
//   - Return the slice containing all elements in order
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A slice containing all elements from the enumeration, or empty slice if enumerator is nil
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to collect all elements. For large
// enumerations, this may be expensive in both time and memory.
//
// ⚠️ Memory note: This operation buffers all elements in memory.
//
// Notes:
//   - If the enumerator is nil, returns an empty slice (not nil)
//   - If the enumeration is empty, returns an empty slice
//   - For very large enumerations, consider processing elements incrementally.
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(n) - allocates memory for all elements
//   - The enumeration stops only when exhausted or if upstream operations stop it
//   - Returned slice preserves the order of elements from the enumeration
//   - Use with caution for infinite or very large enumerations
func (q EnumeratorAny[T]) ToSlice() []T {
	return toSliceInternal(q)
}

// ToSlice converts an ordered enumeration to a slice containing all elements in sorted order.
// This operation is useful for materializing lazy ordered enumerations into concrete sorted collections.
//
// The to slice operation will:
//   - Execute deferred sorting rules and iterate through all elements in sorted order
//   - Collect all elements into a new slice preserving the sorted order
//   - Return the slice containing all elements in their sorted positions
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A slice containing all elements from the enumeration in sorted order,
//	or empty slice if enumerator is nil
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// The operation must iterate through the entire enumeration and perform sorting.
// Time complexity: O(n log n) for sorting + O(n) for collection = O(n log n).
// For large enumerations, this may be expensive in both time and memory.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting and collection.
// Space complexity: O(n) - allocates memory for all elements plus temporary sorting storage.
//
// Notes:
//   - If the enumerator is nil, returns an empty slice (not nil)
//   - If the enumeration is empty, returns an empty slice
//   - Sorting is performed according to all accumulated sorting rules (OrderBy + ThenBy levels)
//   - The returned slice preserves the fully sorted order of elements
//   - Actual sorting computation occurs only during this first materialization
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Use with caution for infinite or very large enumerations
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) ToSlice() []T {
	return toSliceInternal(o.getSortedEnumerator())
}

// ToSlice converts an ordered enumeration to a slice containing all elements in sorted order.
// This operation is useful for materializing lazy ordered enumerations into concrete sorted collections.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The to slice operation will:
//   - Execute deferred sorting rules and iterate through all elements in sorted order
//   - Collect all elements into a new slice preserving the sorted order
//   - Return the slice containing all elements in their sorted positions
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A slice containing all elements from the enumeration in sorted order,
//	or empty slice if enumerator is nil
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// The operation must iterate through the entire enumeration and perform sorting.
// Time complexity: O(n log n) for sorting + O(n) for collection = O(n log n).
// For large enumerations, this may be expensive in both time and memory.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting and collection.
// Space complexity: O(n) - allocates memory for all elements plus temporary sorting storage.
//
// Notes:
//   - If the enumerator is nil, returns an empty slice (not nil)
//   - If the enumeration is empty, returns an empty slice
//   - Sorting is performed according to all accumulated sorting rules (OrderBy + ThenBy levels)
//   - Custom comparer functions are used for element comparison during sorting
//   - The returned slice preserves the fully sorted order of elements
//   - Actual sorting computation occurs only during this first materialization
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Use with caution for infinite or very large enumerations
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) ToSlice() []T {
	return toSliceInternal(o.getSortedEnumerator())
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
