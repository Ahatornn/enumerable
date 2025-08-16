package enumerable

// Take returns a specified number of contiguous elements from the start of an enumeration.
// This operation is useful for pagination, limiting results, or taking samples from sequences.
//
// The take operation will:
// - Yield the first n elements from the enumeration
// - Stop enumeration once n elements have been yielded
// - Handle edge cases gracefully (n <= 0, n >= count)
// - Support early termination when consumer returns false
//
// Parameters:
//   n - the number of elements to take (must be non-negative)
//
// Returns:
//   An Enumerator[T] that yields at most n elements from the start
//
// Notes:
// - If n <= 0, returns an empty enumerator
// - If n >= total number of elements, returns all available elements
// - If the original enumerator is nil, returns an empty enumerator
// - Lazy evaluation - elements are processed and yielded during iteration
// - No elements are buffered - memory efficient
// - Negative values of n are treated as 0
// - Early termination by the consumer stops further enumeration
// - The enumeration stops as soon as n elements are yielded or the source is exhausted
func (q Enumerator[T]) Take(n int) Enumerator[T] {
	if q == nil || n <= 0 {
		return Empty[T]()
	}
	return takeInternal(q, n)
}

// Take returns a specified number of contiguous elements from the start of an enumeration.
// This operation is useful for pagination, limiting results, or taking samples from sequences.
//
// The take operation will:
// - Yield the first n elements from the enumeration
// - Stop enumeration once n elements have been yielded
// - Handle edge cases gracefully (n <= 0, n >= count)
// - Support early termination when consumer returns false
//
// Parameters:
//   n - the number of elements to take (must be non-negative)
//
// Returns:
//   An AnyEnumerator[T] that yields at most n elements from the start
//
// Notes:
// - If n <= 0, returns an empty enumerator
// - If n >= total number of elements, returns all available elements
// - If the original enumerator is nil, returns an empty enumerator
// - Lazy evaluation - elements are processed and yielded during iteration
// - No elements are buffered - memory efficient
// - Negative values of n are treated as 0
// - Early termination by the consumer stops further enumeration
// - The enumeration stops as soon as n elements are yielded or the source is exhausted
func (q AnyEnumerator[T]) Take(n int) AnyEnumerator[T] {
	if q == nil || n <= 0 {
		return EmptyAny[T]()
	}
	return takeInternal(q, n)
}

func takeInternal[T any](enumerator func(func(T) bool), n int) func(func(T) bool) {
	return func(yield func(T) bool) {
		var taken int
		enumerator(func(item T) bool {
			if taken >= n {
				return false
			}
			if !yield(item) {
				return false
			}
			taken++
			return true
		})
	}
}
