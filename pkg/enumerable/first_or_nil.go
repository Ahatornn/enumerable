package enumerable

// FirstOrNil returns a pointer to the first element of an enumeration.
// This operation is useful for getting the first element when it exists,
// with the ability to distinguish between "no elements" and "zero value" cases.
//
// The first operation will:
// - Return a pointer to the first element if the enumeration contains elements
// - Return nil if the enumeration is empty or nil
// - Stop enumeration immediately after finding the first element
// - Handle nil enumerators gracefully
//
// Returns:
//   A pointer to the first element, or nil if enumeration is empty or nil
//
// Notes:
// - If the enumerator is nil, returns nil
// - If the enumeration is empty, returns nil
// - Lazy evaluation stops immediately after finding the first element
// - This is a terminal operation that materializes the enumeration
// - Very efficient - O(1) time complexity for non-empty enumerations
// - No elements are buffered - memory efficient
// - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
// - Safe for all types including those with zero values like 0, "", false, etc.
func (q Enumerator[T]) FirstOrNil() *T {
	return firstOrNilInternal(q)
}

// FirstOrNil returns a pointer to the first element of an enumeration.
// This operation is useful for getting the first element when it exists,
// with the ability to distinguish between "no elements" and "zero value" cases.
//
// The first operation will:
// - Return a pointer to the first element if the enumeration contains elements
// - Return nil if the enumeration is empty or nil
// - Stop enumeration immediately after finding the first element
// - Handle nil enumerators gracefully
//
// Returns:
//   A pointer to the first element, or nil if enumeration is empty or nil
//
// Notes:
// - If the enumerator is nil, returns nil
// - If the enumeration is empty, returns nil
// - Lazy evaluation stops immediately after finding the first element
// - This is a terminal operation that materializes the enumeration
// - Very efficient - O(1) time complexity for non-empty enumerations
// - No elements are buffered - memory efficient
// - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
// - Safe for all types including those with zero values like 0, "", false, etc.
func (q AnyEnumerator[T]) FirstOrNil() *T {
	return firstOrNilInternal(q)
}

func firstOrNilInternal[T any](enumerator func(func(T) bool)) *T {
	if enumerator == nil {
		return nil
	}
	var result *T
	enumerator(func(item T) bool {
		temp := item
		result = &temp
		return false
	})
	return result
}
