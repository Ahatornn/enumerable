package enumerable

// DefaultIfEmpty returns an enumeration that contains the elements of the current enumeration,
// or a single default value if the current enumeration is empty.
//
// The resulting enumeration will:
//   - Yield all elements from the current enumeration if it contains any elements
//   - Yield exactly one element (the specified default value) if the current enumeration is empty
//   - Handle nil enumerators gracefully (treated as empty, yields default value)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	defaultValue - the value to yield if the current enumeration contains no elements
//
// Returns:
//
//	A new Enumerator[T] that yields either the original elements or the default value
//
// Notes:
//   - Nil enumerators are treated as empty enumerations
//   - If the current enumeration yields at least one element, the default value is never yielded
//   - Even if enumeration is terminated early (yield returns false), default value is not added
//     if at least one element was processed
//   - Lazy evaluation - elements are produced on-demand during iteration
//   - No elements are buffered - memory efficient
//   - Safe for use with nil enumerators
func (e Enumerator[T]) DefaultIfEmpty(defaultValue T) Enumerator[T] {
	return defaultIfEmptyInternal(e, defaultValue)
}

// DefaultIfEmpty returns an enumeration that contains the elements of the current enumeration,
// or a single default value if the current enumeration is empty.
//
// The resulting enumeration will:
//   - Yield all elements from the current enumeration if it contains any elements
//   - Yield exactly one element (the specified default value) if the current enumeration is empty
//   - Handle nil enumerators gracefully (treated as empty, yields default value)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	defaultValue - the value to yield if the current enumeration contains no elements
//
// Returns:
//
//	A new EnumeratorAny[T] that yields either the original elements or the default value
//
// Notes:
//   - Nil enumerators are treated as empty enumerations
//   - If the current enumeration yields at least one element, the default value is never yielded
//   - Even if enumeration is terminated early (yield returns false), default value is not added
//     if at least one element was processed
//   - Lazy evaluation - elements are produced on-demand during iteration
//   - No elements are buffered - memory efficient
//   - Safe for use with nil enumerators
func (e EnumeratorAny[T]) DefaultIfEmpty(defaultValue T) EnumeratorAny[T] {
	return defaultIfEmptyInternal(e, defaultValue)
}

func defaultIfEmptyInternal[T any](enumerator func(func(T) bool), defaultValue T) func(func(T) bool) {
	return func(yield func(T) bool) {
		if enumerator == nil {
			yield(defaultValue)
			return
		}

		hasElements := false
		enumerator(func(item T) bool {
			hasElements = true
			return yield(item)
		})

		if !hasElements {
			yield(defaultValue)
		}
	}
}
