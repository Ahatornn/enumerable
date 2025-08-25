package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MaxBy returns the largest element in the enumeration according to a custom comparison function.
// This operation is useful for finding the maximum element based on custom ordering criteria.
//
// The MaxBy operation will:
//   - Return the largest element and true if the enumeration is non-empty
//   - Return the zero value of T and false if the enumeration is empty or nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators gracefully
//   - Use the provided comparer function to determine element ordering
//
// Parameters:
//
//	cmp - a ComparerFunc that defines the ordering of elements by returning:
//	      -1 if x < y, 0 if x == y, +1 if x > y
//
// Returns:
//
//	The maximum element according to the comparer function and true if found,
//	zero value of T and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns zero value and false
//   - If the enumeration is empty, returns zero value and false
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The comparer function is called for each element pair during processing
//   - If multiple elements are equally maximal, the first one encountered is returned
//   - This is a terminal operation that materializes the enumeration
//   - The comparer function must be consistent and deterministic for correct results
//   - For large enumerations, consider the performance cost of the comparison operations
//   - The comparer should satisfy mathematical comparison properties (consistency, antisymmetry, transitivity)
//   - To find the minimum element, use MinBy instead
func (e Enumerator[T]) MaxBy(cmp comparer.ComparerFunc[T]) (T, bool) {
	return maxByInternal(e, cmp)
}

// MaxBy returns the largest element in the enumeration according to a custom comparison function.
// This operation is useful for finding the maximum element based on custom ordering criteria.
//
// The MaxBy operation will:
//   - Return the largest element and true if the enumeration is non-empty
//   - Return the zero value of T and false if the enumeration is empty or nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators gracefully
//   - Use the provided comparer function to determine element ordering
//
// Parameters:
//
//	cmp - a ComparerFunc that defines the ordering of elements by returning:
//	      -1 if x < y, 0 if x == y, +1 if x > y
//
// Returns:
//
//	The maximum element according to the comparer function and true if found,
//	zero value of T and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns zero value and false
//   - If the enumeration is empty, returns zero value and false
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The comparer function is called for each element pair during processing
//   - If multiple elements are equally maximal, the first one encountered is returned
//   - This is a terminal operation that materializes the enumeration
//   - The comparer function must be consistent and deterministic for correct results
//   - For large enumerations, consider the performance cost of the comparison operations
//   - The comparer should satisfy mathematical comparison properties (consistency, antisymmetry, transitivity)
//   - To find the minimum element, use MinBy instead
//   - Thread safety depends on the underlying enumerator implementation
func (e EnumeratorAny[T]) MaxBy(cmp comparer.ComparerFunc[T]) (T, bool) {
	return maxByInternal(e, cmp)
}

func maxByInternal[T any](enumerator func(yield func(T) bool), cmp comparer.ComparerFunc[T]) (T, bool) {
	var max T
	if enumerator == nil || cmp == nil {
		return max, false
	}
	found := false
	enumerator(func(item T) bool {
		if !found {
			max = item
			found = true
		} else if cmp(item, max) > 0 {
			max = item
		}
		return true
	})
	return max, found
}
