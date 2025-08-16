package enumerable

// TakeLast returns a specified number of contiguous elements from the end of an enumeration.
// This operation is useful for getting the final elements, such as last N records,
// trailing averages, or end-of-sequence markers.
//
// The take last operation will:
// - Buffer elements to track the last n elements seen so far
// - Yield the final n elements when enumeration completes
// - Handle edge cases gracefully (n <= 0, n >= count)
// - Support early termination when consumer returns false
//
// Parameters:
//   n - the number of elements to take from the end (must be non-negative)
//
// Returns:
//   An Enumerator[T] that yields the last n elements
//
// ⚠️ Performance note: This operation buffers up to n elements in memory
// to track which elements should be yielded. For large values of n,
// this may consume significant memory. All elements must be processed
// before any are yielded.
//
// ⚠️ Evaluation note: This operation is not lazy in the traditional sense -
// the entire source enumeration must be consumed before yielding begins.
// Elements are yielded in order of their appearance in the original enumeration.
//
// Notes:
// - If n <= 0, returns an empty enumerator
// - If n >= total number of elements, returns all available elements
// - If the original enumerator is nil, returns an empty enumerator
// - Negative values of n are treated as 0
// - The enumeration stops as soon as the consumer returns false
func (q Enumerator[T]) TakeLast(n int) Enumerator[T] {
	if q == nil || n <= 0 {
		return Empty[T]()
	}

	return takeLastInternal(q, n)
}

// TakeLast returns a specified number of contiguous elements from the end of an enumeration.
// This operation is useful for getting the final elements, such as last N records,
// trailing averages, or end-of-sequence markers.
//
// The take last operation will:
// - Buffer elements to track the last n elements seen so far
// - Yield the final n elements when enumeration completes
// - Handle edge cases gracefully (n <= 0, n >= count)
// - Support early termination when consumer returns false
//
// Parameters:
//   n - the number of elements to take from the end (must be non-negative)
//
// Returns:
//   An AnyEnumerator[T] that yields the last n elements
//
// ⚠️ Performance note: This operation buffers up to n elements in memory
// to track which elements should be yielded. For large values of n,
// this may consume significant memory. All elements must be processed
// before any are yielded.
//
// ⚠️ Evaluation note: This operation is not lazy in the traditional sense -
// the entire source enumeration must be consumed before yielding begins.
// Elements are yielded in order of their appearance in the original enumeration.
//
// Notes:
// - If n <= 0, returns an empty enumerator
// - If n >= total number of elements, returns all available elements
// - If the original enumerator is nil, returns an empty enumerator
// - Negative values of n are treated as 0
// - The enumeration stops as soon as the consumer returns false
func (q AnyEnumerator[T]) TakeLast(n int) AnyEnumerator[T] {
	if q == nil || n <= 0 {
		return EmptyAny[T]()
	}

	return takeLastInternal(q, n)
}

func takeLastInternal[T any](enumerator func(func(T) bool), n int) func(func(T) bool) {
	return func(yield func(T) bool) {
		buffer := make([]T, n)
		count := 0
		index := 0

		enumerator(func(item T) bool {
			buffer[index] = item
			index = (index + 1) % n
			if count < n {
				count++
			}
			return true
		})

		if count == 0 {
			return
		}

		startIndex := (index - count + n) % n
		for i := 0; i < count; i++ {
			itemIndex := (startIndex + i) % n
			if !yield(buffer[itemIndex]) {
				return
			}
		}
	}
}
