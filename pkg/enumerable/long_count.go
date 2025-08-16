package enumerable

// LongCount returns the number of elements in an enumeration as an int64.
// This operation is useful for determining the size of large sequences
// where the count might exceed the range of int.
//
// The long count operation will:
// - Iterate through all elements in the enumeration
// - Count each element encountered using int64 to prevent overflow
// - Return the total count as int64
// - Handle nil enumerators gracefully
//
// Returns:
//   The number of elements in the enumeration as int64 (0 for empty or nil enumerations)
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to count all elements. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
// - If the enumerator is nil, returns 0
// - If the enumeration is empty, returns 0
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - In Go, int is platform-dependent (32-bit on 32-bit systems, 64-bit on 64-bit systems)
// - Use LongCount when you expect very large collections that might overflow int
func (q Enumerator[T]) LongCount() int64 {
	return longCountInternal(q)
}

// LongCount returns the number of elements in an enumeration as an int64.
// This operation is useful for determining the size of large sequences
// where the count might exceed the range of int.
//
// The long count operation will:
// - Iterate through all elements in the enumeration
// - Count each element encountered using int64 to prevent overflow
// - Return the total count as int64
// - Handle nil enumerators gracefully
//
// Returns:
//   The number of elements in the enumeration as int64 (0 for empty or nil enumerations)
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to count all elements. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
// - If the enumerator is nil, returns 0
// - If the enumeration is empty, returns 0
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(1) - constant space usage
// - In Go, int is platform-dependent (32-bit on 32-bit systems, 64-bit on 64-bit systems)
// - Use LongCount when you expect very large collections that might overflow int
func (q AnyEnumerator[T]) LongCount() int64 {
	return longCountInternal(q)
}

func longCountInternal[T any](enumerator func(func(T) bool)) int64 {
	if enumerator == nil {
		return 0
	}
	var cnt int64
	enumerator(func(T) bool {
		cnt++
		return true
	})
	return cnt
}
