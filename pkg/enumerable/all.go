package enumerable

// All determines whether all elements in the enumeration satisfy a predicate.
// Returns true if every element matches the predicate, or if the enumeration is empty.
//
// The method will:
// - Apply the predicate to each element in the enumeration
// - Return false immediately when the first non-matching element is found
// - Return true if all elements match or if there are no elements
// - Short-circuit evaluation (stops at first false result)
//
// Parameters:
//   predicate - a function that takes an element and returns true/false
//
// Returns:
//   true if all elements satisfy the predicate or enumeration is empty
//   false if at least one element does not satisfy the predicate
//
// Notes:
// - For empty enumerations, returns true (vacuous truth)
// - For nil enumerators, returns true (consistent with empty behavior)
// - Uses short-circuit evaluation for performance
// - The predicate function should be pure (no side effects)
// - Stops enumeration as soon as a non-matching element is found
func (q Enumerator[T]) All(predicate func(T) bool) bool {
	return allInternal(q, predicate)
}

// All determines whether all elements in the enumeration satisfy a predicate.
// Returns true if every element matches the predicate, or if the enumeration is empty.
//
// The method will:
// - Apply the predicate to each element in the enumeration
// - Return false immediately when the first non-matching element is found
// - Return true if all elements match or if there are no elements
// - Short-circuit evaluation (stops at first false result)
//
// Parameters:
//   predicate - a function that takes an element and returns true/false
//
// Returns:
//   true if all elements satisfy the predicate or enumeration is empty
//   false if at least one element does not satisfy the predicate
//
// Notes:
// - For empty enumerations, returns true (vacuous truth)
// - For nil enumerators, returns true (consistent with empty behavior)
// - Uses short-circuit evaluation for performance
// - The predicate function should be pure (no side effects)
// - Stops enumeration as soon as a non-matching element is found
func (q AnyEnumerator[T]) All(predicate func(T) bool) bool {
	return allInternal(q, predicate)
}

func allInternal[T any](enumerator func(func(T) bool), predicate func(T) bool) bool {
	if enumerator == nil {
		return true
	}
	var anyItem bool
	var all = true
	enumerator(func(item T) bool {
		anyItem = true
		if !predicate(item) {
			all = false
			return false
		}
		return true
	})
	return !anyItem || all
}
