package enumerable

// Any determines whether an enumeration contains any elements.
// This operation is useful for checking if a sequence is non-empty.
//
// The any operation will:
// - Return true if the enumeration contains at least one element
// - Return false if the enumeration is empty or nil
// - Stop enumeration immediately after finding the first element
// - Handle nil enumerators gracefully
//
// Returns:
//   true if the enumeration contains any elements, false otherwise
//
// Notes:
// - If the enumerator is nil, returns false
// - If the enumeration is empty, returns false
// - Lazy evaluation stops immediately after finding the first element
// - This is a terminal operation that materializes the enumeration
// - Very efficient - O(1) time complexity for non-empty enumerations
// - No elements are buffered - memory efficient
func (q Enumerator[T]) Any() bool {
	return anyInternal(q)
}

// Any determines whether an enumeration contains any elements.
// This operation is useful for checking if a sequence is non-empty.
//
// The any operation will:
// - Return true if the enumeration contains at least one element
// - Return false if the enumeration is empty or nil
// - Stop enumeration immediately after finding the first element
// - Handle nil enumerators gracefully
//
// Returns:
//   true if the enumeration contains any elements, false otherwise
//
// Notes:
// - If the enumerator is nil, returns false
// - If the enumeration is empty, returns false
// - Lazy evaluation stops immediately after finding the first element
// - This is a terminal operation that materializes the enumeration
// - Very efficient - O(1) time complexity for non-empty enumerations
// - No elements are buffered - memory efficient
func (q AnyEnumerator[T]) Any() bool {
	return anyInternal(q)
}

func anyInternal[T any](enumerator func(func(T) bool)) bool {
	if enumerator == nil {
		return false
	}
	var has bool
	enumerator(func(T) bool {
		has = true
		return false
	})
	return has
}
