package enumerable

// FromChannel creates an Enumerator[T] that yields values received from a channel.
// The enumeration continues until the channel is closed or the consumer stops iteration.
//
// The enumerator will:
// - Yield each value received from the channel in order
// - Terminate when the channel is closed
// - Support early termination when the consumer returns false
//
// Parameters:
//   ch - the source channel to enumerate (read-only)
//
// Returns:
//   An Enumerator[T] that iterates over channel values
//
// Notes:
// - The enumerator will block waiting for new values when channel is empty
// - If the channel is never closed, iteration may hang indefinitely
// - Safe for nil channels (will act like closed channels, producing no values)
// - Channel receive operations occur during enumeration (not beforehand)
// - The channel should only be read through the enumerator
func FromChannel[T comparable](ch <-chan T) Enumerator[T] {
	return func(yield func(T) bool) {
		if ch == nil {
			return
		}
		for item := range ch {
			if !yield(item) {
				return
			}
		}
	}
}

// FromChannelAny creates an AnyEnumerator[T] that yields values received from a channel.
// The enumeration continues until the channel is closed or the consumer stops iteration.
//
// The enumerator will:
// - Yield each value received from the channel in order
// - Terminate when the channel is closed
// - Support early termination when the consumer returns false
//
// Parameters:
//   ch - the source channel to enumerate (read-only)
//
// Returns:
//   An AnyEnumerator[T] that iterates over channel values
//
// Notes:
// - The enumerator will block waiting for new values when channel is empty
// - If the channel is never closed, iteration may hang indefinitely
// - Safe for nil channels (will act like closed channels, producing no values)
// - Channel receive operations occur during enumeration (not beforehand)
// - The channel should only be read through the enumerator
func FromChannelAny[T any](ch <-chan T) AnyEnumerator[T] {
	return func(yield func(T) bool) {
		if ch == nil {
			return
		}
		for item := range ch {
			if !yield(item) {
				return
			}
		}
	}
}
