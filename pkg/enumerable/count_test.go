package enumerable

import (
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("count non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		count := enumerator.Count()

		if count != 5 {
			t.Errorf("Expected count 5, got %d", count)
		}
	})

	t.Run("count non-empty slice for non-comparable", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		})

		count := enumerator.Count()

		if count != 5 {
			t.Errorf("Expected count 5, got %d", count)
		}
	})

	t.Run("count single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		count := enumerator.Count()

		if count != 1 {
			t.Errorf("Expected count 1, got %d", count)
		}
	})

	t.Run("count empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		count := enumerator.Count()

		if count != 0 {
			t.Errorf("Expected count 0, got %d", count)
		}
	})

	t.Run("count nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		count := enumerator.Count()

		if count != 0 {
			t.Errorf("Expected count 0 for nil enumerator, got %d", count)
		}
	})

	t.Run("count string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		count := enumerator.Count()

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}
	})

	t.Run("count empty string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{})

		count := enumerator.Count()

		if count != 0 {
			t.Errorf("Expected count 0, got %d", count)
		}
	})
}

func TestCountStruct(t *testing.T) {
	t.Run("count struct slice", func(t *testing.T) {
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
		count := enumerator.Count()

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}
	})

	t.Run("count empty struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}

		enumerator := FromSlice(people)
		count := enumerator.Count()

		if count != 0 {
			t.Errorf("Expected count 0, got %d", count)
		}
	})
}

func TestCountBoolean(t *testing.T) {
	t.Run("count boolean slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false})

		count := enumerator.Count()

		if count != 4 {
			t.Errorf("Expected count 4, got %d", count)
		}
	})

	t.Run("count single boolean", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false})

		count := enumerator.Count()

		if count != 1 {
			t.Errorf("Expected count 1, got %d", count)
		}
	})
}

func TestCountWithOperations(t *testing.T) {
	t.Run("count after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		count := filtered.Count()

		if count != 3 {
			t.Errorf("Expected count 3 (even numbers), got %d", count)
		}
	})

	t.Run("count after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		count := distinct.Count()

		if count != 4 {
			t.Errorf("Expected count 4 (unique elements), got %d", count)
		}
	})

	t.Run("count after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(3)

		count := taken.Count()

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}
	})
}

func TestCountEdgeCases(t *testing.T) {
	t.Run("count with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0, 0})

		count := enumerator.Count()

		if count != 4 {
			t.Errorf("Expected count 4, got %d", count)
		}
	})

	t.Run("count with empty strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "", ""})

		count := enumerator.Count()

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}
	})

	t.Run("count repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		count := enumerator.Count()

		if count != 5 {
			t.Errorf("Expected count 5, got %d", count)
		}
	})

	t.Run("count repeat zero", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(42, 0)

		count := enumerator.Count()

		if count != 0 {
			t.Errorf("Expected count 0, got %d", count)
		}
	})

	t.Run("count range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 10)

		count := enumerator.Count()

		if count != 10 {
			t.Errorf("Expected count 10, got %d", count)
		}
	})
}

func TestCountLargeEnumeration(t *testing.T) {
	t.Run("count large slice", func(t *testing.T) {
		t.Parallel()
		largeSlice := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			largeSlice[i] = i
		}

		enumerator := FromSlice(largeSlice)
		count := enumerator.Count()

		if count != 10000 {
			t.Errorf("Expected count 10000, got %d", count)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkCount(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items) // Создаем каждый раз заново
			result := enumerator.Count()
			if result != 5 {
				b.Fatalf("Expected 5, got %d", result)
			}
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result := enumerator.Count()
			if result != 1000 {
				b.Fatalf("Expected 1000, got %d", result)
			}
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result := enumerator.Count()
			if result != 10000 {
				b.Fatalf("Expected 10000, got %d", result)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]int{})
			result := enumerator.Count()
			if result != 0 {
				b.Fatalf("Expected 0, got %d", result)
			}
		}
	})
}
