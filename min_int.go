package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MinInt returns the smallest integer value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric key (such as ID, age, price, etc.)
// derived from complex elements in the sequence.
//
// The MinInt operation will:
//   - Apply the keySelector function to each element to extract an int value
//   - Compare extracted keys using natural numeric ordering (ascending)
//   - Return the smallest int key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts an int key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum int key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, consider the performance cost of key extraction
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinInt(keySelector func(T) int) (int, bool) {
	return minIntInternal(e, keySelector, comparer.ComparerInt)
}

// MinInt returns the smallest integer value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric key (such as ID, age, price, etc.)
// derived from complex elements in the sequence.
//
// The MinInt operation will:
//   - Apply the keySelector function to each element to extract an int value
//   - Compare extracted keys using natural numeric ordering (ascending)
//   - Return the smallest int key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts an int key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum int key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, consider the performance cost of key extraction
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinInt(keySelector func(T) int) (int, bool) {
	return minIntInternal(e, keySelector, comparer.ComparerInt)
}

// MinInt64 returns the smallest int64 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric key (such as ID, timestamp, size, etc.)
// derived from complex elements in the sequence.
//
// The MinInt64 operation will:
//   - Apply the keySelector function to each element to extract an int64 value
//   - Compare extracted keys using natural numeric ordering (ascending)
//   - Return the smallest int64 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts an int64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum int64 key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, consider the performance cost of key extraction
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinInt64(keySelector func(T) int64) (int64, bool) {
	return minInt64Internal(e, keySelector, comparer.ComparerInt64)
}

// MinInt64 returns the smallest int64 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric key (such as ID, timestamp, size, etc.)
// derived from complex elements in the sequence.
//
// The MinInt64 operation will:
//   - Apply the keySelector function to each element to extract an int64 value
//   - Compare extracted keys using natural numeric ordering (ascending)
//   - Return the smallest int64 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts an int64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum int64 key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, consider the performance cost of key extraction
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinInt64(keySelector func(T) int64) (int64, bool) {
	return minInt64Internal(e, keySelector, comparer.ComparerInt64)
}

func minIntInternal[T any](enumerator func(yield func(T) bool), keySelector func(T) int, cmp comparer.ComparerFunc[int]) (int, bool) {
	var minKey int
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return 0, false
	}

	enumerator(func(item T) bool {
		key := keySelector(item)

		if !found {
			minKey = key
			found = true
		} else if cmp(key, minKey) < 0 {
			minKey = key
		}

		return true
	})

	return minKey, found
}

func minInt64Internal[T any](enumerator func(yield func(T) bool), keySelector func(T) int64, cmp comparer.ComparerFunc[int64]) (int64, bool) {
	var minKey int64
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return 0, false
	}

	enumerator(func(item T) bool {
		key := keySelector(item)

		if !found {
			minKey = key
			found = true
		} else if cmp(key, minKey) < 0 {
			minKey = key
		}

		return true
	})

	return minKey, found
}
