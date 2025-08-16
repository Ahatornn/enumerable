package enumerable

// ForEach executes the specified action for each element in the enumeration.
// This operation is useful for performing side effects like printing, logging,
// or modifying external state for each element.
//
// The for each operation will:
// - Execute the action function for each element in the enumeration
// - Process all elements in the enumeration
// - Handle nil enumerators gracefully
// - Not return any value (void operation)
//
// Parameters:
//   action - the action to execute for each element
//
// ⚠️ Performance note: This operation must iterate through the entire
// enumeration, which may be expensive for large enumerations.
//
// ⚠️ Side effects warning: The action function may have side effects.
// Use with caution in functional programming contexts.
//
// Notes:
// - If the enumerator is nil, no action is executed
// - This is a terminal operation that materializes the enumeration
// - All elements are processed regardless of action behavior
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - The enumeration stops only when exhausted or if upstream operations stop it
// - Action function should handle all possible values including zero values
func (q Enumerator[T]) ForEach(action func(T)) {
	forEachInternal(q, action)
}

// ForEach executes the specified action for each element in the enumeration.
// This operation is useful for performing side effects like printing, logging,
// or modifying external state for each element.
//
// The for each operation will:
// - Execute the action function for each element in the enumeration
// - Process all elements in the enumeration
// - Handle nil enumerators gracefully
// - Not return any value (void operation)
//
// Parameters:
//   action - the action to execute for each element
//
// ⚠️ Performance note: This operation must iterate through the entire
// enumeration, which may be expensive for large enumerations.
//
// ⚠️ Side effects warning: The action function may have side effects.
// Use with caution in functional programming contexts.
//
// Notes:
// - If the enumerator is nil, no action is executed
// - This is a terminal operation that materializes the enumeration
// - All elements are processed regardless of action behavior
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - The enumeration stops only when exhausted or if upstream operations stop it
// - Action function should handle all possible values including zero values
func (q AnyEnumerator[T]) ForEach(action func(T)) {
	forEachInternal(q, action)
}

func forEachInternal[T any](enumerator func(func(T) bool), action func(T)) {
	if enumerator == nil {
		return
	}
	enumerator(func(item T) bool {
		action(item)
		return true
	})
}
