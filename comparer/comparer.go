package comparer

// ComparerFunc defines a function type that compares two values of type T and returns their relative ordering.
// It is used throughout the enumerable library for sorting, ordering, and comparison operations.
//
// Type Parameters:
//
//	T - the type of values to compare (can be any type)
//
// The function should return:
//
//	-1 if x is less than y
//	 0 if x is equal to y
//	+1 if x is greater than y
//
// Implementations should ensure the following mathematical properties:
//   - Consistency: f(x, y) should always return the same result for identical inputs
//   - Antisymmetry: if f(x, y) < 0 then f(y, x) > 0
//   - Transitivity: if f(x, y) < 0 and f(y, z) < 0 then f(x, z) < 0
//   - Equality: f(x, y) == 0 if and only if x and y are considered equal
//
// Notes:
//   - For natural ordering of built-in types, use predefined comparer functions
//   - For custom ordering logic, create ComparerFunc instances
//   - Thread safety depends on the function implementation
//   - Nil handling should be consistent within the function
type ComparerFunc[T any] func(x, y T) int
