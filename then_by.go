package enumerable

import "github.com/ahatornn/enumerable/comparer"

// ThenBy performs a subsequent ordering of the elements in ascending order according to a comparer function.
// This method adds a secondary (or further) sorting level to an existing ordered sequence.
//
// The ThenBy operation:
//   - Accumulates additional sorting rules without immediate execution (lazy evaluation)
//   - Returns an OrderEnumerator that supports further chaining with additional ThenBy operations
//   - Uses the provided comparer function to determine secondary element ordering
//   - Maintains stable sorting (equal elements preserve relative order)
//   - Works with comparable types using direct comparison
//
// Parameters:
//
//	comparer - a ComparerFunc that defines the secondary sort order by comparing two elements of type T
//	           The function should return:
//	             < 0 if first element is less than second
//	             = 0 if elements are equal
//	             > 0 if first element is greater than second
//
// Returns:
//
//	An OrderEnumerator[T] that contains elements sorted by both primary and secondary criteria
//
// ⚠️ Performance note: This is a deferred execution operation that accumulates sorting rules.
// Actual sorting computation occurs during first enumeration. Time complexity: O(1) for rule
// accumulation, O(n log n) for actual sorting during enumeration with all accumulated rules.
//
// Notes:
//   - This is a lazy operation - no sorting occurs until enumeration begins
//   - Can be chained multiple times to create multiple sorting levels
//   - Each ThenBy adds a new sorting level with lower priority than previous levels
//   - The comparer function must be deterministic and consistent
//   - Works with both comparable and non-comparable types through appropriate comparer functions
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
func (o OrderEnumerator[T]) ThenBy(comparer comparer.ComparerFunc[T]) OrderEnumerator[T] {
	return o.addSortLevel(comparer, false)
}

// ThenByDescending performs a subsequent ordering of the elements in descending order according to a comparer function.
// This method adds a secondary (or further) sorting level to an existing ordered sequence in reverse order.
//
// The ThenByDescending operation:
//   - Accumulates additional sorting rules without immediate execution (lazy evaluation)
//   - Returns an OrderEnumerator that supports further chaining with additional ThenBy operations
//   - Uses the provided comparer function to determine secondary element ordering (reversed)
//   - Maintains stable sorting (equal elements preserve relative order)
//   - Works with comparable types using direct comparison
//
// Parameters:
//
//	comparer - a ComparerFunc that defines the secondary sort order by comparing two elements of type T
//	           The function should return:
//	             < 0 if first element is less than second
//	             = 0 if elements are equal
//	             > 0 if first element is greater than second
//	           Note: The final order for this level will be reversed (descending)
//
// Returns:
//
//	An OrderEnumerator[T] that contains elements sorted by both primary and secondary criteria
//
// ⚠️ Performance note: This is a deferred execution operation that accumulates sorting rules.
// Actual sorting computation occurs during first enumeration. Time complexity: O(1) for rule
// accumulation, O(n log n) for actual sorting during enumeration with all accumulated rules.
//
// Notes:
//   - This is a lazy operation - no sorting occurs until enumeration begins
//   - Can be chained multiple times to create multiple sorting levels
//   - Each ThenByDescending adds a new sorting level with lower priority than previous levels
//   - The comparer function must be deterministic and consistent
//   - Works with both comparable and non-comparable types through appropriate comparer functions
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
func (o OrderEnumerator[T]) ThenByDescending(comparer comparer.ComparerFunc[T]) OrderEnumerator[T] {
	return o.addSortLevel(comparer, true)
}

// ThenBy performs a subsequent ordering of the elements in ascending order according to a comparer function.
// This method adds a secondary (or further) sorting level to an existing ordered sequence for any type T.
//
// The ThenBy operation:
//   - Accumulates additional sorting rules without immediate execution (lazy evaluation)
//   - Returns an OrderEnumeratorAny that supports further chaining with additional ThenBy operations
//   - Uses the provided comparer function to determine secondary element ordering
//   - Maintains stable sorting (equal elements preserve relative order)
//   - Works with any type T, including non-comparable types
//
// Parameters:
//
//	comparer - a ComparerFunc that defines the secondary sort order by comparing two elements of type T
//	           The function should return:
//	             < 0 if first element is less than second
//	             = 0 if elements are equal
//	             > 0 if first element is greater than second
//
// Returns:
//
//	An OrderEnumeratorAny[T] that contains elements sorted by both primary and secondary criteria
//
// ⚠️ Performance note: This is a deferred execution operation that accumulates sorting rules.
// Actual sorting computation occurs during first enumeration. Time complexity: O(1) for rule
// accumulation, O(n log n) for actual sorting during enumeration with all accumulated rules.
//
// Notes:
//   - This is a lazy operation - no sorting occurs until enumeration begins
//   - Can be chained multiple times to create multiple sorting levels
//   - Each ThenBy adds a new sorting level with lower priority than previous levels
//   - Custom comparer functions must handle all possible input values, including nil
//   - The comparer function must be deterministic and consistent
//   - Works with any type T, including complex structs with non-comparable fields
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
func (o OrderEnumeratorAny[T]) ThenBy(comparer comparer.ComparerFunc[T]) OrderEnumeratorAny[T] {
	return o.addSortLevel(comparer, false)
}

// ThenByDescending performs a subsequent ordering of the elements in descending order according to a comparer function.
// This method adds a secondary (or further) sorting level to an existing ordered sequence for any type T.
//
// The ThenByDescending operation:
//   - Accumulates additional sorting rules without immediate execution (lazy evaluation)
//   - Returns an OrderEnumeratorAny that supports further chaining with additional ThenBy operations
//   - Uses the provided comparer function to determine secondary element ordering (reversed)
//   - Maintains stable sorting (equal elements preserve relative order)
//   - Works with any type T, including non-comparable types
//
// Parameters:
//
//	comparer - a ComparerFunc that defines the secondary sort order by comparing two elements of type T
//	           The function should return:
//	             < 0 if first element is less than second
//	             = 0 if elements are equal
//	             > 0 if first element is greater than second
//	           Note: The final order for this level will be reversed (descending)
//
// Returns:
//
//	An OrderEnumeratorAny[T] that contains elements sorted by both primary and secondary criteria
//
// ⚠️ Performance note: This is a deferred execution operation that accumulates sorting rules.
// Actual sorting computation occurs during first enumeration. Time complexity: O(1) for rule
// accumulation, O(n log n) for actual sorting during enumeration with all accumulated rules.
//
// Notes:
//   - This is a lazy operation - no sorting occurs until enumeration begins
//   - Can be chained multiple times to create multiple sorting levels
//   - Each ThenByDescending adds a new sorting level with lower priority than previous levels
//   - Custom comparer functions must handle all possible input values, including nil
//   - The comparer function must be deterministic and consistent
//   - Works with any type T, including complex structs with non-comparable fields
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
func (o OrderEnumeratorAny[T]) ThenByDescending(comparer comparer.ComparerFunc[T]) OrderEnumeratorAny[T] {
	return o.addSortLevel(comparer, true)
}
