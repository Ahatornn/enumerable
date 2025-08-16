package enumerable

// TakeWhile returns elements from an enumeration as long as a specified condition is true.
// This operation is useful for taking elements until a certain condition is met,
// such as taking elements while they are valid or within a range.
//
// The take while operation will:
// - Yield elements from the start while the predicate returns true
// - Stop enumeration as soon as the predicate returns false for an element
// - Handle edge cases gracefully (nil enumerator, always false predicate)
// - Support early termination when consumer returns false
//
// Parameters:
//   predicate - a function that determines whether to take an element
//
// Returns:
//   An Enumerator[T] that yields elements while the condition is true
//
// Notes:
// - If the predicate immediately returns false for the first element, returns empty enumerator
// - If the predicate never returns false, returns all elements from the enumeration
// - If the original enumerator is nil, returns an empty enumerator
// - Lazy evaluation - elements are processed and evaluated during iteration
// - No elements are buffered - memory efficient
// - The enumeration stops as soon as the predicate returns false or consumer returns false
// - Once the predicate returns false for any element, no subsequent elements are evaluated
func (q Enumerator[T]) TakeWhile(predicate func(T) bool) Enumerator[T] {
	if q == nil {
		return Empty[T]()
	}
	return takeWhileInternal(q, predicate)
}

// TakeWhile returns elements from an enumeration as long as a specified condition is true.
// This operation is useful for taking elements until a certain condition is met,
// such as taking elements while they are valid or within a range.
//
// The take while operation will:
// - Yield elements from the start while the predicate returns true
// - Stop enumeration as soon as the predicate returns false for an element
// - Handle edge cases gracefully (nil enumerator, always false predicate)
// - Support early termination when consumer returns false
//
// Parameters:
//   predicate - a function that determines whether to take an element
//
// Returns:
//   An AnyEnumerator[T] that yields elements while the condition is true
//
// Notes:
// - If the predicate immediately returns false for the first element, returns empty enumerator
// - If the predicate never returns false, returns all elements from the enumeration
// - If the original enumerator is nil, returns an empty enumerator
// - Lazy evaluation - elements are processed and evaluated during iteration
// - No elements are buffered - memory efficient
// - The enumeration stops as soon as the predicate returns false or consumer returns false
// - Once the predicate returns false for any element, no subsequent elements are evaluated
func (q AnyEnumerator[T]) TakeWhile(predicate func(T) bool) AnyEnumerator[T] {
	if q == nil {
		return EmptyAny[T]()
	}
	return takeWhileInternal(q, predicate)
}

func takeWhileInternal[T any](enumerator func(func(T) bool), predicate func(T) bool) func(func(T) bool) {
	return func(yield func(T) bool) {
		enumerator(func(item T) bool {
			if !predicate(item) {
				return false
			}
			return yield(item)
		})
	}
}
