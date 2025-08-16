package enumerable

// Enumerator represents a sequence of values that can be iterated over
// using Go's range loop syntax (Go 1.22+ range-over-func feature).
// The iteration can be stopped early by the consumer returning false
// from the yield function.
//
// Type Parameters:
//   T - the type of values to enumerate (must be comparable)
//
// Notes:
// - Thread safety depends on the implementation
// - Designed to work with Go 1.22+ range-over-func feature
type Enumerator[T comparable] func(yield func(T) bool)

// AnyEnumerator represents a sequence of values that can be iterated over
// using Go's range loop syntax (Go 1.22+ range-over-func feature).
// The iteration can be stopped early by the consumer returning false from
//  the yield function.
//
// Type Parameters:
//   T - the type of values to enumerate (no constraints)
//
// Notes:
// - Thread safety depends on the implementation
// - Designed to work with Go 1.22+ range-over-func feature
type AnyEnumerator[T any] func(yield func(T) bool)
