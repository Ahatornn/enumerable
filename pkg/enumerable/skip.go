package enumerable

// Skip bypasses a specified number of elements in an enumeration and then yields the remaining elements.
// This operation is useful for pagination, skipping headers, or bypassing initial elements.
//
// The skip operation will:
// - Bypass the first n elements from the enumeration
// - Yield all remaining elements in order
// - Handle edge cases gracefully (n <= 0, n >= count)
// - Support early termination when consumer returns false
//
// Parameters:
//   n - the number of elements to skip (must be non-negative)
//
// Returns:
//   An Enumerator[T] that yields elements after skipping the first n elements
//
// Notes:
// - If n <= 0, returns the original enumerator unchanged
// - If n >= total number of elements, returns an empty enumerator
// - If the original enumerator is nil, returns an empty enumerator (not nil)
// - Lazy evaluation - elements are processed and skipped during iteration
// - No elements are buffered - memory efficient
// - Negative values of n are treated as 0
// - The enumeration is consumed sequentially, so skipped elements are still processed
func (q Enumerator[T]) Skip(n int) Enumerator[T] {
	if q == nil {
		return Empty[T]()
	}
	return skipInternal(q, n)
}

// Skip bypasses a specified number of elements in an enumeration and then yields the remaining elements.
// This operation is useful for pagination, skipping headers, or bypassing initial elements.
//
// The skip operation will:
// - Bypass the first n elements from the enumeration
// - Yield all remaining elements in order
// - Handle edge cases gracefully (n <= 0, n >= count)
// - Support early termination when consumer returns false
//
// Parameters:
//   n - the number of elements to skip (must be non-negative)
//
// Returns:
//   An AnyEnumerator[T] that yields elements after skipping the first n elements
//
// Notes:
// - If n <= 0, returns the original enumerator unchanged
// - If n >= total number of elements, returns an empty enumerator
// - If the original enumerator is nil, returns an empty enumerator (not nil)
// - Lazy evaluation - elements are processed and skipped during iteration
// - No elements are buffered - memory efficient
// - Negative values of n are treated as 0
// - The enumeration is consumed sequentially, so skipped elements are still processed
func (q AnyEnumerator[T]) Skip(n int) AnyEnumerator[T] {
	if q == nil {
		return EmptyAny[T]()
	}
	return skipInternal(q, n)
}

func skipInternal[T any](enumerator func(func(T) bool), n int) func(func(T) bool) {
	if n <= 0 {
		return enumerator
	}
	return func(yield func(T) bool) {
		var skipped int
		enumerator(func(item T) bool {
			if skipped < n {
				skipped++
				return true
			}
			return yield(item)
		})
	}
}
