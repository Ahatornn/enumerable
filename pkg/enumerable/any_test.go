package enumerable

import (
	"testing"
)

func TestAny(t *testing.T) {
	t.Run("non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.Any()

		if !result {
			t.Error("Expected true for non-empty slice, got false")
		}
	})

	t.Run("non-empty slice for non-comparable", func(t *testing.T) {
		t.Parallel()

		enumerator1 := FromSliceAny([][]int{{1, 2}, {3, 4}})
		if !enumerator1.Any() {
			t.Error("Expected true for non-empty slice of slices, got false")
		}

		enumerator2 := FromSliceAny([]map[string]int{{"key": 1}, {"key": 2}})
		if !enumerator2.Any() {
			t.Error("Expected true for non-empty slice of maps, got false")
		}

		enumerator3 := FromSliceAny([]func(){func() {}, func() {}})
		if !enumerator3.Any() {
			t.Error("Expected true for non-empty slice of functions, got false")
		}

		type Container struct {
			Data []int
			Fn   func()
		}
		enumerator4 := FromSliceAny([]Container{
			{Data: []int{1, 2}, Fn: func() {}},
		})
		if !enumerator4.Any() {
			t.Error("Expected true for non-empty slice of structs with non-comparable fields, got false")
		}
	})

	t.Run("single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.Any()

		if !result {
			t.Error("Expected true for single element slice, got false")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.Any()

		if result {
			t.Error("Expected false for empty slice, got true")
		}
	})

	t.Run("nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.Any()

		if result {
			t.Error("Expected false for nil enumerator, got true")
		}
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world"})

		result := enumerator.Any()

		if !result {
			t.Error("Expected true for non-empty string slice, got false")
		}
	})

	t.Run("empty string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{})

		result := enumerator.Any()

		if result {
			t.Error("Expected false for empty string slice, got true")
		}
	})
}

func TestAnyStruct(t *testing.T) {
	t.Run("struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		}

		enumerator := FromSlice(people)
		result := enumerator.Any()

		if !result {
			t.Error("Expected true for non-empty struct slice, got false")
		}
	})

	t.Run("empty struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}

		enumerator := FromSlice(people)
		result := enumerator.Any()

		if result {
			t.Error("Expected false for empty struct slice, got true")
		}
	})
}

func TestAnyBoolean(t *testing.T) {
	t.Run("boolean slice with elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false})

		result := enumerator.Any()

		if !result {
			t.Error("Expected true for non-empty boolean slice, got false")
		}
	})

	t.Run("single false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false})

		result := enumerator.Any()

		if !result {
			t.Error("Expected true for single false element (element exists), got false")
		}
	})
}

func TestAnyEarlyTermination(t *testing.T) {
	t.Run("stops after first element", func(t *testing.T) {
		t.Parallel()
		callCount := 0

		enumerator := func(yield func(int) bool) {
			for i := 1; i <= 100; i++ {
				callCount++
				if !yield(i) {
					return
				}
			}
		}

		var enum Enumerator[int] = enumerator
		result := enum.Any()

		if !result {
			t.Error("Expected true, got false")
		}

		if callCount != 1 {
			t.Errorf("Expected exactly 1 call, got %d", callCount)
		}
	})

	t.Run("large enumeration efficiency", func(t *testing.T) {
		t.Parallel()
		largeSlice := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			largeSlice[i] = i
		}

		enumerator := FromSlice(largeSlice)
		result := enumerator.Any()

		if !result {
			t.Error("Expected true for large enumeration, got false")
		}
	})
}

func TestAnyEdgeCases(t *testing.T) {
	t.Run("zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0})

		result := enumerator.Any()

		if !result {
			t.Error("Expected true for slice with zero values, got false")
		}
	})

	t.Run("empty strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", ""})

		result := enumerator.Any()

		if !result {
			t.Error("Expected true for slice with empty strings, got false")
		}
	})

	t.Run("repeat with zero count", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(42, 0)

		result := enumerator.Any()

		if result {
			t.Error("Expected false for Repeat with zero count, got true")
		}
	})

	t.Run("repeat with one count", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(42, 1)

		result := enumerator.Any()

		if !result {
			t.Error("Expected true for Repeat with one count, got false")
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkAny(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.Any()
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.Any()
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.Any()
		}
	})
}
