package enumerable

// FirstOrDefault returns the first element of an enumeration, or a default value
// if the enumeration is empty or nil.
//
// The first or default operation will:
// - Return the first element from the enumeration if it exists
// - Return the provided default value if the enumeration is empty or nil
// - Stop enumeration immediately after finding the first element
// - Handle nil enumerators gracefully
//
// Parameters:
//   defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//   The first element of the enumeration, or the default value if enumeration is empty
//
// Notes:
// - If the enumerator is nil, returns the defaultValue
// - If the enumeration is empty, returns the defaultValue
// - Lazy evaluation stops immediately after finding the first element
// - This is a terminal operation that materializes the enumeration
// - Very efficient - O(1) time complexity for non-empty enumerations
// - No elements are buffered - memory efficient
// - Unlike FirstOrNil(), this method returns the value directly, not a pointer
// - Safe for all types including those with zero values like 0, "", false, etc.
// - When using zero value as default, consider using FirstOrNil() for distinction
func (q Enumerator[T]) FirstOrDefault(defaultValue T) T {
	if first := q.FirstOrNil(); first != nil {
		return *first
	}
	return defaultValue
}

// FirstOrDefault returns the first element of an enumeration, or a default value
// if the enumeration is empty or nil.
//
// The first or default operation will:
// - Return the first element from the enumeration if it exists
// - Return the provided default value if the enumeration is empty or nil
// - Stop enumeration immediately after finding the first element
// - Handle nil enumerators gracefully
//
// Parameters:
//   defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//   The first element of the enumeration, or the default value if enumeration is empty
//
// Notes:
// - If the enumerator is nil, returns the defaultValue
// - If the enumeration is empty, returns the defaultValue
// - Lazy evaluation stops immediately after finding the first element
// - This is a terminal operation that materializes the enumeration
// - Very efficient - O(1) time complexity for non-empty enumerations
// - No elements are buffered - memory efficient
// - Unlike FirstOrNil(), this method returns the value directly, not a pointer
// - Safe for all types including those with zero values like 0, "", false, etc.
// - When using zero value as default, consider using FirstOrNil() for distinction
func (q AnyEnumerator[T]) FirstOrDefault(defaultValue T) T {
	if first := q.FirstOrNil(); first != nil {
		return *first
	}
	return defaultValue
}
