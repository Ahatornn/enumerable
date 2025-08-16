package enumerable

// ToChannel converts an enumeration to a channel that yields all elements.
// This operation is useful for converting lazy enumerations into channel-based
// processing pipelines or for interoperability with goroutines.
//
// The to channel operation will:
// - Create a new channel with specified buffer size
// - Start a goroutine that iterates through the enumeration
// - Send each element to the channel
// - Close the channel when enumeration is complete
// - Handle nil enumerators gracefully
//
// Parameters:
//   bufferSize - the buffer size for the returned channel (0 for unbuffered)
//
// Returns:
//   A read-only channel containing all elements from the enumeration
//
// ⚠️ Resource management: This operation starts a goroutine that runs
// until the enumeration is complete. Always consume all elements or
// the goroutine may block indefinitely.
//
// ⚠️ Blocking behavior: The goroutine will block when trying to send
// to a full channel if there's no reader. Use appropriate buffer size.
//
// ⚠️ Warning: If the returned channel is not fully consumed, the goroutine
// may leak. Always range over the entire channel or ensure proper cleanup.
//
// Notes:
// - If the enumerator is nil, returns a closed channel
// - If the enumeration is empty, returns an empty but closed channel
// - The goroutine automatically closes the channel when done
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(bufferSize) plus any upstream buffering
// - The enumeration runs in a separate goroutine, enabling concurrent processing
func (q Enumerator[T]) ToChannel(bufferSize int) <-chan T {
	return toChannelInternal(q, bufferSize)
}

// ToChannel converts an enumeration to a channel that yields all elements.
// This operation is useful for converting lazy enumerations into channel-based
// processing pipelines or for interoperability with goroutines.
//
// The to channel operation will:
// - Create a new channel with specified buffer size
// - Start a goroutine that iterates through the enumeration
// - Send each element to the channel
// - Close the channel when enumeration is complete
// - Handle nil enumerators gracefully
//
// Parameters:
//   bufferSize - the buffer size for the returned channel (0 for unbuffered)
//
// Returns:
//   A read-only channel containing all elements from the enumeration
//
// ⚠️ Resource management: This operation starts a goroutine that runs
// until the enumeration is complete. Always consume all elements or
// the goroutine may block indefinitely.
//
// ⚠️ Blocking behavior: The goroutine will block when trying to send
// to a full channel if there's no reader. Use appropriate buffer size.
//
// ⚠️ Warning: If the returned channel is not fully consumed, the goroutine
// may leak. Always range over the entire channel or ensure proper cleanup.
//
// Notes:
// - If the enumerator is nil, returns a closed channel
// - If the enumeration is empty, returns an empty but closed channel
// - The goroutine automatically closes the channel when done
// - Time complexity: O(n) where n is the number of elements
// - Space complexity: O(bufferSize) plus any upstream buffering
// - The enumeration runs in a separate goroutine, enabling concurrent processing
func (q AnyEnumerator[T]) ToChannel(bufferSize int) <-chan T {
	return toChannelInternal(q, bufferSize)
}

func toChannelInternal[T any](enumerator func(func(T) bool), bufferSize int) <-chan T {
	if enumerator == nil {
		ch := make(chan T)
		close(ch)
		return ch
	}

	ch := make(chan T, bufferSize)

	go func() {
		defer close(ch)
		enumerator(func(item T) bool {
			ch <- item
			return true
		})
	}()

	return ch
}
