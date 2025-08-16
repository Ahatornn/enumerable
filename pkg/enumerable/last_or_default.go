package enumerable

// LastOrDefault returns the last element of an enumeration, or a default value
// if the enumeration is empty or nil.
//
// The last or default operation will:
// - Iterate through all elements in the enumeration
// - Keep track of the most recent element encountered
// - Return the last element if enumeration contains elements
// - Return the provided default value if the enumeration is empty or nil
// - Handle nil enumerators gracefully
//
// Parameters:
//   defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//   The last element of the enumeration, or the default value if enumeration is empty
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the last element. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
// - If the enumerator is nil, returns the defaultValue
// - If the enumeration is empty, returns the defaultValue
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - Unlike LastOrNil(), this method returns the value directly, not a pointer
// - Safe for all types including those with zero values like 0, "", false, etc.
// - When using zero value as default, consider using LastOrNil() for distinction
func (q Enumerator[T]) LastOrDefault(defaultValue T) T {
	if last := q.LastOrNil(); last != nil {
		return *last
	}
	return defaultValue
}

// LastOrDefault returns the last element of an enumeration, or a default value
// if the enumeration is empty or nil.
//
// The last or default operation will:
// - Iterate through all elements in the enumeration
// - Keep track of the most recent element encountered
// - Return the last element if enumeration contains elements
// - Return the provided default value if the enumeration is empty or nil
// - Handle nil enumerators gracefully
//
// Parameters:
//   defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//   The last element of the enumeration, or the default value if enumeration is empty
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the last element. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
// - If the enumerator is nil, returns the defaultValue
// - If the enumeration is empty, returns the defaultValue
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - Unlike LastOrNil(), this method returns the value directly, not a pointer
// - Safe for all types including those with zero values like 0, "", false, etc.
// - When using zero value as default, consider using LastOrNil() for distinction
func (q AnyEnumerator[T]) LastOrDefault(defaultValue T) T {
	if last := q.LastOrNil(); last != nil {
		return *last
	}
	return defaultValue
}
