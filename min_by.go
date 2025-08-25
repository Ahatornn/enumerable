package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MinBy returns the smallest element in the enumeration according to a custom comparison function.
// This operation is useful for finding the minimum element based on custom ordering criteria.
//
// The minby operation will:
//   - Return the smallest element and true if the enumeration is non-empty
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
//	The minimum element according to the comparer function and true if found,
//	zero value of T and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns zero value and false
//   - If the enumeration is empty, returns zero value and false
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The comparer function is called for each element pair during processing
//   - If multiple elements are equally minimal, the first one encountered is returned
//   - This is a terminal operation that materializes the enumeration
//   - The comparer function must be consistent and deterministic for correct results
//   - For large enumerations, consider the performance cost of the comparison operations
//   - The comparer should satisfy mathematical comparison properties (consistency, antisymmetry, transitivity)
func (e Enumerator[T]) MinBy(cmp comparer.ComparerFunc[T]) (T, bool) {
	return minByInternal(e, cmp)
}

// MinBy returns the smallest element in the enumeration according to a custom comparison function.
// This operation is useful for finding the minimum element based on custom ordering criteria.
//
// The minby operation will:
//   - Return the smallest element and true if the enumeration is non-empty
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
//	The minimum element according to the comparer function and true if found,
//	zero value of T and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns zero value and false
//   - If the enumeration is empty, returns zero value and false
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The comparer function is called for each element pair during processing
//   - If multiple elements are equally minimal, the first one encountered is returned
//   - This is a terminal operation that materializes the enumeration
//   - The comparer function must be consistent and deterministic for correct results
//   - For large enumerations, consider the performance cost of the comparison operations
//   - The comparer should satisfy mathematical comparison properties (consistency, antisymmetry, transitivity)
func (e EnumeratorAny[T]) MinBy(cmp comparer.ComparerFunc[T]) (T, bool) {
	return minByInternal(e, cmp)
}

func minByInternal[T any](enumerator func(yield func(T) bool), cmp comparer.ComparerFunc[T]) (T, bool) {
	var min T
	if enumerator == nil || cmp == nil {
		return min, false
	}
	found := false
	enumerator(func(item T) bool {
		if !found {
			min = item
			found = true
		} else if cmp(item, min) < 0 {
			min = item
		}
		return true
	})
	return min, found
}
