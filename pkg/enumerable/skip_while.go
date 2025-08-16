package enumerable

// SkipWhile bypasses elements in an enumeration as long as a specified condition is true
// and then yields the remaining elements.
// This operation is useful for skipping elements until a certain condition is met.
//
// The skip while operation will:
// - Bypass elements from the beginning while the predicate returns true
// - Once the predicate returns false for an element, yield that element and all subsequent elements
// - Support early termination when consumer returns false
//
// Parameters:
//   predicate - a function that determines whether to skip an element
//
// Returns:
//   An Enumerator[T] that yields elements after skipping initial elements that match the condition
//
// Notes:
// - If the predicate never returns false, returns an empty enumerator
// - If the predicate immediately returns false for the first element, returns the original enumerator
// - If the original enumerator is nil, returns an empty enumerator
// - Lazy evaluation - elements are processed and evaluated during iteration
// - No elements are buffered - memory efficient
// - The enumeration stops skipping as soon as the first non-matching element is found
// - Once skipping stops, all remaining elements are yielded (even if they would match the predicate)
func (q Enumerator[T]) SkipWhile(predicate func(T) bool) Enumerator[T] {
	if q == nil {
		return Empty[T]()
	}
	return skipWhileInternal(q, predicate)
}

// SkipWhile bypasses elements in an enumeration as long as a specified condition is true
// and then yields the remaining elements.
// This operation is useful for skipping elements until a certain condition is met.
//
// The skip while operation will:
// - Bypass elements from the beginning while the predicate returns true
// - Once the predicate returns false for an element, yield that element and all subsequent elements
// - Support early termination when consumer returns false
//
// Parameters:
//   predicate - a function that determines whether to skip an element
//
// Returns:
//   An AnyEnumerator[T] that yields elements after skipping initial elements that match the condition
//
// Notes:
// - If the predicate never returns false, returns an empty enumerator
// - If the predicate immediately returns false for the first element, returns the original enumerator
// - If the original enumerator is nil, returns an empty enumerator
// - Lazy evaluation - elements are processed and evaluated during iteration
// - No elements are buffered - memory efficient
// - The enumeration stops skipping as soon as the first non-matching element is found
// - Once skipping stops, all remaining elements are yielded (even if they would match the predicate)
func (q AnyEnumerator[T]) SkipWhile(predicate func(T) bool) AnyEnumerator[T] {
	if q == nil {
		return EmptyAny[T]()
	}
	return skipWhileInternal(q, predicate)
}

func skipWhileInternal[T any](enumerator func(func(T) bool), predicate func(T) bool) func(func(T) bool) {
	return func(yield func(T) bool) {
		skipping := true
		enumerator(func(item T) bool {
			if skipping {
				if predicate(item) {
					return true
				}
				skipping = false
			}
			return yield(item)
		})
	}
}
