package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MaxInt returns the largest integer value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric key (such as ID, age, price, etc.)
// derived from complex elements in the sequence.
//
// The MaxInt operation will:
//   - Apply the keySelector function to each element to extract an int value
//   - Compare extracted keys using natural numeric ordering (descending)
//   - Return the largest int key and true if the enumeration is non-empty
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
//	The maximum int key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, consider the performance cost of key extraction
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxInt(keySelector func(T) int) (int, bool) {
	return minIntInternal(e, keySelector, reverseIntComparer)
}

// MaxInt returns the largest integer value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric key (such as ID, age, price, etc.)
// derived from complex elements in the sequence.
//
// The MaxInt operation will:
//   - Apply the keySelector function to each element to extract an int value
//   - Compare extracted keys using natural numeric ordering (descending)
//   - Return the largest int key and true if the enumeration is non-empty
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
//	The maximum int key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, consider the performance cost of key extraction
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxInt(keySelector func(T) int) (int, bool) {
	return minIntInternal(e, keySelector, reverseIntComparer)
}

// MaxInt64 returns the largest int64 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric key (such as ID, timestamp, size, etc.)
// derived from complex elements in the sequence.
//
// The MaxInt64 operation will:
//   - Apply the keySelector function to each element to extract an int64 value
//   - Compare extracted keys using natural numeric ordering (descending)
//   - Return the largest int64 key and true if the enumeration is non-empty
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
//	The maximum int64 key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, consider the performance cost of key extraction
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxInt64(keySelector func(T) int64) (int64, bool) {
	return minInt64Internal(e, keySelector, reverseInt64Comparer)
}

// MaxInt64 returns the largest int64 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric key (such as ID, timestamp, size, etc.)
// derived from complex elements in the sequence.
//
// The MaxInt64 operation will:
//   - Apply the keySelector function to each element to extract an int64 value
//   - Compare extracted keys using natural numeric ordering (descending)
//   - Return the largest int64 key and true if the enumeration is non-empty
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
//	The maximum int64 key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, consider the performance cost of key extraction
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxInt64(keySelector func(T) int64) (int64, bool) {
	return minInt64Internal(e, keySelector, reverseInt64Comparer)
}

var reverseIntComparer = comparer.ComparerFunc[int](func(a, b int) int {
	return comparer.ComparerInt(b, a)
})

var reverseInt64Comparer = comparer.ComparerFunc[int64](func(a, b int64) int {
	return comparer.ComparerInt64(b, a)
})
