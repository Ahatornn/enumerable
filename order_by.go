package enumerable

import "github.com/ahatornn/enumerable/comparer"

// OrderBy sorts the elements of a sequence in ascending order according to a comparer function.
// This method is the primary sorting operation that establishes the first sorting level.
//
// The OrderBy operation:
//   - Accumulates sorting rules without immediate execution (lazy evaluation)
//   - Returns an OrderEnumerator that supports fluent chaining with ThenBy operations
//   - Uses the provided comparer function to determine element ordering
//   - Maintains stable sorting (equal elements preserve relative order)
//   - Works with comparable types using direct comparison
//
// Parameters:
//
//	comparer - a ComparerFunc that defines the sort order by comparing two elements of type T
//	           The function should return:
//	             < 0 if first element is less than second
//	             = 0 if elements are equal
//	             > 0 if first element is greater than second
//
// Returns:
//
//	An OrderEnumerator[T] that contains the sorted elements and supports further sorting operations
//
// ⚠️ Performance note: This is a deferred execution operation that accumulates sorting rules.
// Actual sorting computation occurs during first enumeration. Time complexity: O(1) for rule
// accumulation, O(n log n) for actual sorting during enumeration.
//
// Notes:
//   - This is a lazy operation - no sorting occurs until enumeration begins
//   - Multiple OrderBy calls overwrite previous sorting rules (use ThenBy for multiple levels)
//   - The comparer function must be deterministic and consistent
//   - For subsequent sorting levels, use ThenBy or ThenByDescending
//   - Works with both comparable and non-comparable types through appropriate comparer functions
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
func (e Enumerator[T]) OrderBy(comparer comparer.ComparerFunc[T]) OrderEnumerator[T] {
	return newOrderEnumerator(e, comparer, false)
}

// OrderByDescending sorts the elements of a sequence in descending order according to a comparer function.
// This method is the primary sorting operation that establishes the first sorting level in reverse order.
//
// The OrderByDescending operation:
//   - Accumulates sorting rules without immediate execution (lazy evaluation)
//   - Returns an OrderEnumerator that supports fluent chaining with ThenBy operations
//   - Uses the provided comparer function to determine element ordering (reversed)
//   - Maintains stable sorting (equal elements preserve relative order)
//   - Works with comparable types using direct comparison
//
// Parameters:
//
//	comparer - a ComparerFunc that defines the sort order by comparing two elements of type T
//	           The function should return:
//	             < 0 if first element is less than second
//	             = 0 if elements are equal
//	             > 0 if first element is greater than second
//	           Note: The final order will be reversed (descending)
//
// Returns:
//
//	An OrderEnumerator[T] that contains the sorted elements and supports further sorting operations
//
// ⚠️ Performance note: This is a deferred execution operation that accumulates sorting rules.
// Actual sorting computation occurs during first enumeration. Time complexity: O(1) for rule
// accumulation, O(n log n) for actual sorting during enumeration.
//
// Notes:
//   - This is a lazy operation - no sorting occurs until enumeration begins
//   - Multiple OrderByDescending calls overwrite previous sorting rules (use ThenByDescending for multiple levels)
//   - The comparer function must be deterministic and consistent
//   - For subsequent sorting levels, use ThenBy or ThenByDescending
//   - Works with both comparable and non-comparable types through appropriate comparer functions
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
func (e Enumerator[T]) OrderByDescending(comparer comparer.ComparerFunc[T]) OrderEnumerator[T] {
	return newOrderEnumerator(e, comparer, true)
}

// OrderBy sorts the elements of a sequence in ascending order according to a comparer function.
// This method is the primary sorting operation for any type T, including non-comparable types.
//
// The OrderBy operation:
//   - Accumulates sorting rules without immediate execution (lazy evaluation)
//   - Returns an OrderEnumeratorAny that supports fluent chaining with ThenBy operations
//   - Uses the provided comparer function to determine element ordering
//   - Maintains stable sorting (equal elements preserve relative order)
//   - Works with any type T, including non-comparable types
//
// Parameters:
//
//	comparer - a ComparerFunc that defines the sort order by comparing two elements of type T
//	           The function should return:
//	             < 0 if first element is less than second
//	             = 0 if elements are equal
//	             > 0 if first element is greater than second
//
// Returns:
//
//	An OrderEnumeratorAny[T] that contains the sorted elements and supports further sorting operations
//
// ⚠️ Performance note: This is a deferred execution operation that accumulates sorting rules.
// Actual sorting computation occurs during first enumeration. Time complexity: O(1) for rule
// accumulation, O(n log n) for actual sorting during enumeration.
//
// Notes:
//   - This is a lazy operation - no sorting occurs until enumeration begins
//   - Works with any type T, including complex structs with non-comparable fields
//   - Custom comparer functions must handle all possible input values, including nil
//   - The comparer function must be deterministic and consistent
//   - For subsequent sorting levels, use ThenBy or ThenByDescending
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
func (e EnumeratorAny[T]) OrderBy(comparer comparer.ComparerFunc[T]) OrderEnumeratorAny[T] {
	return newOrderEnumeratorAny(e, comparer, false)
}

// OrderByDescending sorts the elements of a sequence in descending order according to a comparer function.
// This method is the primary sorting operation for any type T, including non-comparable types.
//
// The OrderByDescending operation:
//   - Accumulates sorting rules without immediate execution (lazy evaluation)
//   - Returns an OrderEnumeratorAny that supports fluent chaining with ThenBy operations
//   - Uses the provided comparer function to determine element ordering (reversed)
//   - Maintains stable sorting (equal elements preserve relative order)
//   - Works with any type T, including non-comparable types
//
// Parameters:
//
//	comparer - a ComparerFunc that defines the sort order by comparing two elements of type T
//	           The function should return:
//	             < 0 if first element is less than second
//	             = 0 if elements are equal
//	             > 0 if first element is greater than second
//	           Note: The final order will be reversed (descending)
//
// Returns:
//
//	An OrderEnumeratorAny[T] that contains the sorted elements and supports further sorting operations
//
// ⚠️ Performance note: This is a deferred execution operation that accumulates sorting rules.
// Actual sorting computation occurs during first enumeration. Time complexity: O(1) for rule
// accumulation, O(n log n) for actual sorting during enumeration.
//
// Notes:
//   - This is a lazy operation - no sorting occurs until enumeration begins
//   - Works with any type T, including complex structs with non-comparable fields
//   - Custom comparer functions must handle all possible input values, including nil
//   - The comparer function must be deterministic and consistent
//   - For subsequent sorting levels, use ThenBy or ThenByDescending
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
func (e EnumeratorAny[T]) OrderByDescending(comparer comparer.ComparerFunc[T]) OrderEnumeratorAny[T] {
	return newOrderEnumeratorAny(e, comparer, true)
}
