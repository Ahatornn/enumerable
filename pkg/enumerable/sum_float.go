package enumerable

// SumFloat computes the sum of float32 values obtained by applying a selector function
// to each element in the enumeration.
// This operation is useful for calculating totals, aggregates, or numeric summaries
// with floating-point precision.
//
// The sum float operation will:
// - Apply the selector function to each element to extract a float32 value
// - Sum all the extracted float32 values
// - Return the total sum
// - Handle nil enumerators gracefully
//
// Parameters:
//   selector - a function that extracts a float32 value from each element
//
// Returns:
//   The sum of all float32 values extracted from the elements
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to sum all values. For large
// enumerations, this may be expensive.
//
// ⚠️ Precision warning: Floating-point arithmetic may introduce small
// rounding errors. Consider using appropriate rounding for display.
//
// Notes:
// - If the enumerator is nil, returns 0
// - If the enumeration is empty, returns 0
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - The enumeration stops only when exhausted or if upstream operations stop it
// - Selector function should handle all possible input values safely
// - For double precision, consider implementing or using SumFloat64
func (q Enumerator[T]) SumFloat(selector func(T) float32) float32 {
	return sumFloatInternal(q, selector)
}

// SumFloat computes the sum of float32 values obtained by applying a selector function
// to each element in the enumeration.
// This operation is useful for calculating totals, aggregates, or numeric summaries
// with floating-point precision.
//
// The sum float operation will:
// - Apply the selector function to each element to extract a float32 value
// - Sum all the extracted float32 values
// - Return the total sum
// - Handle nil enumerators gracefully
//
// Parameters:
//   selector - a function that extracts a float32 value from each element
//
// Returns:
//   The sum of all float32 values extracted from the elements
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to sum all values. For large
// enumerations, this may be expensive.
//
// ⚠️ Precision warning: Floating-point arithmetic may introduce small
// rounding errors. Consider using appropriate rounding for display.
//
// Notes:
// - If the enumerator is nil, returns 0
// - If the enumeration is empty, returns 0
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - The enumeration stops only when exhausted or if upstream operations stop it
// - Selector function should handle all possible input values safely
// - For double precision, consider implementing or using SumFloat64
func (q AnyEnumerator[T]) SumFloat(selector func(T) float32) float32 {
	return sumFloatInternal(q, selector)
}

func sumFloatInternal[T any](enumerator func(func(T) bool), selector func(T) float32) float32 {
	if enumerator == nil {
		return 0
	}
	var sum float32
	enumerator(func(item T) bool {
		sum += selector(item)
		return true
	})
	return sum
}
