package enumerable

import (
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("all elements satisfy predicate", func(t *testing.T) {
		enumerator := FromSlice([]int{2, 4, 6, 8})

		result := enumerator.All(func(n int) bool {
			return n%2 == 0
		})

		if !result {
			t.Error("Expected true when all elements satisfy predicate")
		}
	})

	t.Run("all elements satisfy predicate for non-comparable", func(t *testing.T) {

		enumerator := FromSliceAny([][]int{
			{3, 2},
			{3, 4},
			{3, 6},
		})

		result := enumerator.All(func(slice []int) bool {
			return len(slice) > 0 && slice[0] == 3
		})

		if !result {
			t.Error("Expected true when all elements satisfy predicate")
		}
	})

	t.Run("not all elements satisfy predicate", func(t *testing.T) {
		enumerator := FromSlice([]int{2, 4, 5, 8})

		result := enumerator.All(func(n int) bool {
			return n%2 == 0
		})

		if result {
			t.Error("Expected false when not all elements satisfy predicate")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		enumerator := FromSlice([]int{})

		result := enumerator.All(func(n int) bool {
			return n > 0
		})

		if !result {
			t.Error("Expected true for empty enumeration (vacuous truth)")
		}
	})

	t.Run("nil enumerator", func(t *testing.T) {
		var enumerator Enumerator[int] = nil

		result := enumerator.All(func(n int) bool {
			return n > 0
		})

		if !result {
			t.Error("Expected true for nil enumerator")
		}
	})

	t.Run("single element satisfies", func(t *testing.T) {
		enumerator := FromSlice([]int{42})

		result := enumerator.All(func(n int) bool {
			return n > 0
		})

		if !result {
			t.Error("Expected true when single element satisfies predicate")
		}
	})

	t.Run("single element does not satisfy", func(t *testing.T) {
		enumerator := FromSlice([]int{-1})

		result := enumerator.All(func(n int) bool {
			return n > 0
		})

		if result {
			t.Error("Expected false when single element does not satisfy predicate")
		}
	})

	t.Run("string all non-empty", func(t *testing.T) {
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.All(func(s string) bool {
			return s != ""
		})

		if !result {
			t.Error("Expected true when all strings are non-empty")
		}
	})

	t.Run("string with empty element", func(t *testing.T) {
		enumerator := FromSlice([]string{"hello", "", "go"})

		result := enumerator.All(func(s string) bool {
			return s != ""
		})

		if result {
			t.Error("Expected false when contains empty string")
		}
	})

	t.Run("early termination optimization", func(t *testing.T) {
		items := make([]int, 1000)
		items[0] = -1
		for i := 1; i < 1000; i++ {
			items[i] = 2
		}

		enumerator := FromSlice(items)
		callCount := 0

		result := enumerator.All(func(n int) bool {
			callCount++
			return n > 0
		})

		if result {
			t.Error("Expected false when first element does not satisfy")
		}

		if callCount != 1 {
			t.Errorf("Expected predicate to be called only once, got %d calls", callCount)
		}
	})
}

func TestAllWithStructs(t *testing.T) {
	t.Run("struct all valid", func(t *testing.T) {
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

		result := enumerator.All(func(p Person) bool {
			return p.Age > 0
		})

		if !result {
			t.Error("Expected true when all persons have positive age")
		}
	})

	t.Run("struct with invalid element", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: -5},
			{Name: "Charlie", Age: 35},
		}

		enumerator := FromSlice(people)

		result := enumerator.All(func(p Person) bool {
			return p.Age > 0
		})

		if result {
			t.Error("Expected false when one person has invalid age")
		}
	})
}

func TestAllEdgeCases(t *testing.T) {
	t.Run("all with zero values", func(t *testing.T) {
		enumerator := FromSlice([]int{0, 0, 0})

		result := enumerator.All(func(n int) bool {
			return n == 0
		})

		if !result {
			t.Error("Expected true when all elements are zero")
		}
	})

	t.Run("all with boolean values", func(t *testing.T) {
		enumerator := FromSlice([]bool{true, true, true})

		result := enumerator.All(func(b bool) bool {
			return b
		})

		if !result {
			t.Error("Expected true when all boolean values are true")
		}
	})

	t.Run("all with mixed boolean values", func(t *testing.T) {
		enumerator := FromSlice([]bool{true, false, true})

		result := enumerator.All(func(b bool) bool {
			return b
		})

		if result {
			t.Error("Expected false when contains false value")
		}
	})

	t.Run("predicate always returns true", func(t *testing.T) {
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.All(func(n int) bool {
			return true
		})

		if !result {
			t.Error("Expected true when predicate always returns true")
		}
	})

	t.Run("predicate always returns false", func(t *testing.T) {
		enumerator := FromSlice([]int{1, 2, 3})

		result := enumerator.All(func(n int) bool {
			return false
		})

		if result {
			t.Error("Expected false when predicate always returns false")
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkAll(b *testing.B) {
	b.Run("small slice all true", func(b *testing.B) {
		items := []int{2, 4, 6, 8, 10}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			enumerator.All(func(n int) bool {
				return n%2 == 0
			})
		}
	})

	b.Run("large slice early termination", func(b *testing.B) {
		items := make([]int, 10000)
		items[0] = 1
		for i := 1; i < 10000; i++ {
			items[i] = 2
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			enumerator.All(func(n int) bool {
				return n%2 == 0
			})
		}
	})

	b.Run("large slice all true", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = 2
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			enumerator.All(func(n int) bool {
				return n%2 == 0
			})
		}
	})
}
