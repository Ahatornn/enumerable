package enumerable

// LastOrNil returns a pointer to the last element of an enumeration,
// or nil if the enumeration is empty or nil.
//
// The last or nil operation will:
// - Iterate through all elements in the enumeration
// - Keep track of the most recent element encountered
// - Return a pointer to the last element if enumeration contains elements
// - Return nil if the enumeration is empty or nil
// - Handle nil enumerators gracefully
//
// Returns:
//   A pointer to the last element, or nil if enumeration is empty or nil
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the last element. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
// - If the enumerator is nil, returns nil
// - If the enumeration is empty, returns nil
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
// - Safe for all types including those with zero values like 0, "", false, etc.
func (q Enumerator[T]) LastOrNil() *T {
	return lastOrNilInternal(q)
}

// LastOrNil returns a pointer to the last element of an enumeration,
// or nil if the enumeration is empty or nil.
//
// The last or nil operation will:
// - Iterate through all elements in the enumeration
// - Keep track of the most recent element encountered
// - Return a pointer to the last element if enumeration contains elements
// - Return nil if the enumeration is empty or nil
// - Handle nil enumerators gracefully
//
// Returns:
//   A pointer to the last element, or nil if enumeration is empty or nil
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the last element. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
// - If the enumerator is nil, returns nil
// - If the enumeration is empty, returns nil
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
// - Safe for all types including those with zero values like 0, "", false, etc.
func (q AnyEnumerator[T]) LastOrNil() *T {
	return lastOrNilInternal(q)
}

func lastOrNilInternal[T any](enumerator func(func(T) bool)) *T {
	if enumerator == nil {
		return nil
	}

	var lastItem T
	var found bool

	enumerator(func(item T) bool {
		lastItem = item
		found = true
		return true
	})

	if !found {
		return nil
	}

	return &lastItem
}
