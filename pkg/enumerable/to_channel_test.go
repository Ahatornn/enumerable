package enumerable

import (
	"testing"
	"time"
)

func TestToChannel(t *testing.T) {
	t.Run("convert non-empty slice to channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 2, 3, 4, 5}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("convert non-empty slice to channel for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		})

		ch := enumerator.ToChannel(2)

		var result [][]int
		for item := range ch {
			copy := make([]int, len(item))
			for i, v := range item {
				copy[i] = v
			}
			result = append(result, copy)
		}

		expected := [][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if len(result[i]) != len(v) {
				t.Errorf("Expected slice length %d at index %d, got %d", len(v), i, len(result[i]))
				continue
			}
			for j, val := range v {
				if result[i][j] != val {
					t.Errorf("Expected %d at index [%d][%d], got %d", val, i, j, result[i][j])
				}
			}
		}
	})

	t.Run("convert single element to channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 1 {
			t.Fatalf("Expected length 1, got %d", len(result))
		}

		if result[0] != 42 {
			t.Errorf("Expected 42, got %d", result[0])
		}
	})

	t.Run("convert empty slice to channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("convert nil enumerator to channel", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice from nil enumerator, got length %d", len(result))
		}
	})

	t.Run("convert string slice to channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		ch := enumerator.ToChannel(0)

		var result []string
		for item := range ch {
			result = append(result, item)
		}

		expected := []string{"hello", "world", "go"}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
			}
		}
	})
}

func TestToChannelBuffered(t *testing.T) {
	t.Run("buffered channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		ch := enumerator.ToChannel(3) // Buffer size 3

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 2, 3, 4, 5}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})
}

func TestToChannelStruct(t *testing.T) {
	t.Run("convert struct slice to channel", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}

		enumerator := FromSlice(people)
		ch := enumerator.ToChannel(0)

		var result []Person
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != len(people) {
			t.Fatalf("Expected length %d, got %d", len(people), len(result))
		}

		for i, v := range people {
			if result[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, result[i])
			}
		}
	})
}

func TestToChannelWithOperations(t *testing.T) {
	t.Run("to channel after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		ch := filtered.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{2, 4, 6}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("to channel after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		ch := distinct.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 2, 3, 4}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})
}

func TestToChannelConcurrent(t *testing.T) {
	t.Run("concurrent channel consumption", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 10)

		ch := enumerator.ToChannel(5)

		resultChan := make(chan []int)
		go func() {
			var result []int
			for item := range ch {
				result = append(result, item)
			}
			resultChan <- result
		}()

		select {
		case result := <-resultChan:
			expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			if len(result) != len(expected) {
				t.Fatalf("Expected length %d, got %d", len(expected), len(result))
			}

			for i, v := range expected {
				if result[i] != v {
					t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
				}
			}
		case <-time.After(1 * time.Second):
			t.Fatal("Timeout waiting for channel consumption")
		}
	})
}

func TestToChannelEdgeCases(t *testing.T) {
	t.Run("to channel with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0})

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{0, 0, 0}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("to channel with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 3)

		ch := enumerator.ToChannel(0)

		var result []string
		for item := range ch {
			result = append(result, item)
		}

		expected := []string{"test", "test", "test"}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
			}
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkToChannel(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ch := enumerator.ToChannel(0)
			for range ch {
				// Consume all items
			}
		}
	})

	b.Run("medium enumeration buffered", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ch := enumerator.ToChannel(100)
			for range ch {
				// Consume all items
			}
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ch := enumerator.ToChannel(1000)
			for range ch {
				// Consume all items
			}
		}
	})
}
