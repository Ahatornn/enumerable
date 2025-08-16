package enumerable

import (
	"testing"
)

func TestLongCount(t *testing.T) {
	t.Run("long count non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		count := enumerator.LongCount()

		if count != 5 {
			t.Errorf("Expected count 5, got %d", count)
		}
	})

	t.Run("long count non-empty slice for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		})

		count := enumerator.LongCount()

		if count != 5 {
			t.Errorf("Expected count 5, got %d", count)
		}
	})

	t.Run("long count single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		count := enumerator.LongCount()

		if count != 1 {
			t.Errorf("Expected count 1, got %d", count)
		}
	})

	t.Run("long count empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		count := enumerator.LongCount()

		if count != 0 {
			t.Errorf("Expected count 0, got %d", count)
		}
	})

	t.Run("long count nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		count := enumerator.LongCount()

		if count != 0 {
			t.Errorf("Expected count 0 for nil enumerator, got %d", count)
		}
	})

	t.Run("long count string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		count := enumerator.LongCount()

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}
	})
}

func TestLongCountStruct(t *testing.T) {
	t.Run("long count struct slice", func(t *testing.T) {
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
		count := enumerator.LongCount()

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}
	})
}

func TestLongCountWithOperations(t *testing.T) {
	t.Run("long count after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		count := filtered.LongCount()

		if count != 3 {
			t.Errorf("Expected count 3 (even numbers), got %d", count)
		}
	})

	t.Run("long count after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		count := distinct.LongCount()

		if count != 4 {
			t.Errorf("Expected count 4 (unique elements), got %d", count)
		}
	})
}

func TestLongCountEdgeCases(t *testing.T) {
	t.Run("long count with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0, 0})

		count := enumerator.LongCount()

		if count != 4 {
			t.Errorf("Expected count 4, got %d", count)
		}
	})

	t.Run("long count repeat large number", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 1000)

		count := enumerator.LongCount()

		if count != 1000 {
			t.Errorf("Expected count 1000, got %d", count)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkLongCount(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LongCount()
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LongCount()
		}
	})
}
