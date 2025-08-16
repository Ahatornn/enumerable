package enumerable

import (
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("basic range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 5)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		enumerator(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("basic range for non-comparable slice", func(t *testing.T) {
		t.Parallel()
		enumerator := RangeAny(1, 5)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		enumerator(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("zero start", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(0, 3)

		expected := []int{0, 1, 2}
		actual := []int{}

		enumerator(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("negative start", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(-2, 4)

		expected := []int{-2, -1, 0, 1}
		actual := []int{}

		enumerator(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("zero count", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(5, 0)

		count := 0
		enumerator(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items, got %d", count)
		}
	})

	t.Run("one count", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(10, 1)

		var result int
		found := false

		enumerator(func(item int) bool {
			result = item
			found = true
			return true
		})

		if !found {
			t.Error("Expected to find one element, but found none")
		}

		if result != 10 {
			t.Errorf("Expected 10, got %d", result)
		}
	})

	t.Run("early termination", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(100, 10)

		actual := []int{}
		enumerator(func(item int) bool {
			actual = append(actual, item)
			return len(actual) < 3
		})

		expected := []int{100, 101, 102}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("large range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1000, 100)

		first := -1
		last := -1
		count := 0

		enumerator(func(item int) bool {
			if first == -1 {
				first = item
			}
			last = item
			count++
			return true
		})

		if first != 1000 {
			t.Errorf("Expected first element to be 1000, got %d", first)
		}

		if last != 1099 {
			t.Errorf("Expected last element to be 1099, got %d", last)
		}

		if count != 100 {
			t.Errorf("Expected count to be 100, got %d", count)
		}
	})
}

func TestRangeEdgeCases(t *testing.T) {
	t.Run("negative count", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(5, -1)

		count := 0
		enumerator(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items with negative count, got %d", count)
		}
	})

	t.Run("start at boundary", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(0, 0)

		count := 0
		enumerator(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items, got %d", count)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkRange(b *testing.B) {
	b.Run("small range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			enumerator := Range(0, 10)
			enumerator(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			enumerator := Range(0, 1000)
			enumerator(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
