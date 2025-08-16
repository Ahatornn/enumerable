package enumerable

// FromSlice creates an Enumerator[T] that yields all elements from the input slice in order.
//
// The enumerator will produce exactly len(items) values, one for each element in the original
// slice, preserving their original order. The iteration can be stopped early by the consumer.
//
// Parameters:
//   items - slice of elements to enumerate (elements must be comparable)
//
// Returns:
//   An Enumerator[T] that iterates over the slice elements
//
// Notes:
// - The slice is captured by reference (modifications will affect iteration)
// - For empty slices, produces no values (like Empty())
// - Safe for nil slices (treated as empty)
// - Preserves the original element order
func FromSlice[T comparable](items []T) Enumerator[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

// FromSliceAny creates an AnyEnumerator[T] that yields all elements from the input slice in order.
//
// The enumerator will produce exactly len(items) values, one for each element in the original
// slice, preserving their original order. The iteration can be stopped early by the consumer.
//
// Parameters:
//   items - slice of elements to enumerate (no type constraints)
//
// Returns:
//   An AnyEnumerator[T] that iterates over the slice elements
//
// Notes:
// - The slice is captured by reference (modifications will affect iteration)
// - For empty slices, produces no values
// - Safe for nil slices (treated as empty)
// - Preserves the original element order
func FromSliceAny[T any](items []T) AnyEnumerator[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}
