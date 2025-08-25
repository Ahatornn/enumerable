package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MinBool returns the smallest boolean value extracted from elements of the enumeration
// using a key selector function and natural boolean ordering (false < true).
// This operation is useful for finding the minimum boolean key derived from complex elements.
//
// The MinBool operation will:
//   - Apply the keySelector function to each element to extract a bool value
//   - Compare extracted keys using natural boolean ordering (false < true)
//   - Return the smallest bool key and true if the enumeration is non-empty
//   - Return false and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when false is found (since false is the minimum value)
//
// Parameters:
//
//	keySelector - a function that extracts a bool key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (false, false).
//
// Returns:
//
//	The minimum bool key value extracted from elements and true if found,
//	false and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (false, false)
//   - If keySelector is nil, returns (false, false)
//   - If the enumeration is empty, returns (false, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when false is found (since false < true)
//   - For large enumerations without false values, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinBool(keySelector func(T) bool) (bool, bool) {
	return minBoolInternal(e, keySelector, comparer.ComparerBool)
}

// MinBool returns the smallest boolean value extracted from elements of the enumeration
// using a key selector function and natural boolean ordering (false < true).
// This operation is useful for finding the minimum boolean key derived from complex elements.
//
// Same as Enumerator[T].MinBool, but operates on EnumeratorAny[T].
//
// Parameters:
//
//	keySelector - a function that extracts a bool key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (false, false).
//
// Returns:
//
//	The minimum bool key value extracted from elements and true if found,
//	false and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (false, false)
//   - If keySelector is nil, returns (false, false)
//   - If the enumeration is empty, returns (false, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when false is found (since false < true)
//   - For large enumerations without false values, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinBool(keySelector func(T) bool) (bool, bool) {
	return minBoolInternal(e, keySelector, comparer.ComparerBool)
}

func minBoolInternal[T any](
	enumerator func(yield func(T) bool),
	keySelector func(T) bool,
	cmp comparer.ComparerFunc[bool],
) (bool, bool) {
	if enumerator == nil || keySelector == nil || cmp == nil {
		return false, false
	}

	found := false
	minKey := false

	enumerator(func(item T) bool {
		key := keySelector(item)

		if !key {
			minKey = false
			found = true
			return false
		}

		if !found {
			minKey = true
			found = true
		}

		return true
	})

	return minKey, found
}
