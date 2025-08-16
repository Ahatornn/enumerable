package enumerable

// Repeat generates a sequence containing the same item repeated 'count' times.
//
// Parameters:
//   item  - value to repeat (any comparable type)
//   count - number of repetitions (must be non-negative)
//
// Returns:
//   An Enumerator[T] that can be used in range loops.
//
// Notes:
// - For count = 0, produces an empty sequence (no iterations)
// - For count < 0, behavior is undefined (should be avoided)
// - Works with any comparable type T (int, string, structs etc.)
func Repeat[T comparable](item T, count int) Enumerator[T] {
	return func(yield func(T) bool) {
		for i := 0; i < count; i++ {
			if !yield(item) {
				return
			}
		}
	}
}

// RepeatAny generates a sequence containing the same item repeated 'count' times.
//
// Parameters:
//   item  - value to repeat (any type)
//   count - number of repetitions (must be non-negative)
//
// Returns:
//   An AnyEnumerator[T] that can be used in range loops.
//
// Notes:
// - For count = 0, produces an empty sequence (no iterations)
// - For count < 0, behavior is undefined (should be avoided)
// - Works with any type T (comparable and non-comparable types)
func RepeatAny[T any](item T, count int) AnyEnumerator[T] {
	return func(yield func(T) bool) {
		for i := 0; i < count; i++ {
			if !yield(item) {
				return
			}
		}
	}
}
