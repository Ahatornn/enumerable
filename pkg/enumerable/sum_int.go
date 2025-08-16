package enumerable

// SumInt computes the sum of integers obtained by applying a selector function
// to each element in the enumeration.
// This operation is useful for calculating totals, aggregates, or numeric summaries.
//
// The sum int operation will:
// - Apply the selector function to each element to extract an integer value
// - Sum all the extracted integer values
// - Return the total sum
// - Handle nil enumerators gracefully
//
// Parameters:
//   selector - a function that extracts an integer value from each element
//
// Returns:
//   The sum of all integer values extracted from the elements
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to sum all values. For large
// enumerations, this may be expensive.
//
// ⚠️ Overflow warning: Integer overflow may occur with very large sums.
// Consider using SumInt64 for larger ranges.
//
// Notes:
// - If the enumerator is nil, returns 0
// - If the enumeration is empty, returns 0
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - The enumeration stops only when exhausted or if upstream operations stop it
// - Selector function should handle all possible input values safely
func (q Enumerator[T]) SumInt(selector func(T) int) int {
	return sumIntInternal(q, selector)
}

// SumInt computes the sum of integers obtained by applying a selector function
// to each element in the enumeration.
// This operation is useful for calculating totals, aggregates, or numeric summaries.
//
// The sum int operation will:
// - Apply the selector function to each element to extract an integer value
// - Sum all the extracted integer values
// - Return the total sum
// - Handle nil enumerators gracefully
//
// Parameters:
//   selector - a function that extracts an integer value from each element
//
// Returns:
//   The sum of all integer values extracted from the elements
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to sum all values. For large
// enumerations, this may be expensive.
//
// ⚠️ Overflow warning: Integer overflow may occur with very large sums.
// Consider using SumInt64 for larger ranges.
//
// Notes:
// - If the enumerator is nil, returns 0
// - If the enumeration is empty, returns 0
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - The enumeration stops only when exhausted or if upstream operations stop it
// - Selector function should handle all possible input values safely
func (q AnyEnumerator[T]) SumInt(selector func(T) int) int {
	return sumIntInternal(q, selector)
}

func sumIntInternal[T any](enumerator func(func(T) bool), selector func(T) int) int {
	if enumerator == nil {
		return 0
	}
	var sum int
	enumerator(func(item T) bool {
		sum += selector(item)
		return true
	})
	return sum
}
