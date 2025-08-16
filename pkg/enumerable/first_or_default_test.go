package enumerable

import (
	"testing"
)

func TestFirstOrDefault(t *testing.T) {
	t.Run("first element from non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.FirstOrDefault(-1)

		if result != 1 {
			t.Errorf("Expected value 1, got %d", result)
		}
	})

	t.Run("first element from non-empty for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		defaultSlice := []int{-1, -2}

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		})

		result := enumerator.FirstOrDefault(defaultSlice)

		if len(result) != 2 {
			t.Errorf("Expected first element length 2, got %d", len(result))
		}

		if result[0] != 1 || result[1] != 2 {
			t.Errorf("Expected first element [1,2], got %v", result)
		}
	})

	t.Run("first element from single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.FirstOrDefault(-1)

		if result != 42 {
			t.Errorf("Expected value 42, got %d", result)
		}
	})

	t.Run("first element from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.FirstOrDefault(-1)

		if result != -1 {
			t.Errorf("Expected default value -1, got %d", result)
		}
	})

	t.Run("first element from empty non-comparable slice", func(t *testing.T) {
		t.Parallel()

		defaultSlice := []int{-1, -2}
		enumerator := FromSliceAny([][]int{})

		result := enumerator.FirstOrDefault(defaultSlice)

		if len(result) != 2 {
			t.Errorf("Expected default length 2, got %d", len(result))
		}

		if result[0] != -1 || result[1] != -2 {
			t.Errorf("Expected default [-1,-2], got %v", result)
		}
	})

	t.Run("first element from nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.FirstOrDefault(-1)

		if result != -1 {
			t.Errorf("Expected default value -1 for nil enumerator, got %d", result)
		}
	})

	t.Run("first string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.FirstOrDefault("default")

		if result != "hello" {
			t.Errorf("Expected value 'hello', got '%s'", result)
		}
	})

	t.Run("first empty string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "world", "go"})

		result := enumerator.FirstOrDefault("default")

		if result != "" {
			t.Errorf("Expected empty string, got '%s'", result)
		}
	})

	t.Run("first string from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{})

		result := enumerator.FirstOrDefault("default")

		if result != "default" {
			t.Errorf("Expected default value 'default', got '%s'", result)
		}
	})
}

func TestFirstOrDefaultStruct(t *testing.T) {
	t.Run("first struct element", func(t *testing.T) {
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

		defaultPerson := Person{Name: "Unknown", Age: 0}
		enumerator := FromSlice(people)
		result := enumerator.FirstOrDefault(defaultPerson)

		expected := Person{Name: "Alice", Age: 30}
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	t.Run("first struct from empty slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}
		defaultPerson := Person{Name: "Unknown", Age: 0}

		enumerator := FromSlice(people)
		result := enumerator.FirstOrDefault(defaultPerson)

		if result != defaultPerson {
			t.Errorf("Expected default person %+v, got %+v", defaultPerson, result)
		}
	})
}

func TestFirstOrDefaultBoolean(t *testing.T) {
	t.Run("first true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		result := enumerator.FirstOrDefault(false)

		if result != true {
			t.Errorf("Expected true, got %t", result)
		}
	})

	t.Run("first false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, true, false})

		result := enumerator.FirstOrDefault(true)

		if result != false {
			t.Errorf("Expected false, got %t", result)
		}
	})

	t.Run("first boolean from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{})

		result := enumerator.FirstOrDefault(true)

		if result != true {
			t.Errorf("Expected default value true, got %t", result)
		}
	})
}

func TestFirstOrDefaultEarlyTermination(t *testing.T) {
	t.Run("stops after first element", func(t *testing.T) {
		t.Parallel()
		callCount := 0

		// Создаем enumerator, который подсчитывает вызовы
		enumerator := func(yield func(int) bool) {
			for i := 1; i <= 100; i++ {
				callCount++
				if !yield(i) {
					return
				}
			}
		}

		var enum Enumerator[int] = enumerator
		result := enum.FirstOrDefault(-1)

		if result != 1 {
			t.Errorf("Expected value 1, got %d", result)
		}

		if callCount != 1 {
			t.Errorf("Expected exactly 1 call, got %d", callCount)
		}
	})
}

func TestFirstOrDefaultEdgeCases(t *testing.T) {
	t.Run("first zero value with non-zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 1, 2, 3})

		result := enumerator.FirstOrDefault(-1)

		if result != 0 {
			t.Errorf("Expected value 0, got %d", result)
		}
	})

	t.Run("first zero value with zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 1, 2, 3})

		result := enumerator.FirstOrDefault(0)

		if result != 0 {
			t.Errorf("Expected value 0, got %d", result)
		}
	})

	t.Run("empty slice with zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.FirstOrDefault(0)

		if result != 0 {
			t.Errorf("Expected default value 0, got %d", result)
		}
	})

	t.Run("distinguishing zero value from default", func(t *testing.T) {
		t.Parallel()
		withZero := FromSlice([]int{0})
		empty := FromSlice([]int{})

		zeroResult := withZero.FirstOrDefault(-1)
		emptyResult := empty.FirstOrDefault(-1)

		if zeroResult != 0 {
			t.Errorf("Expected 0 from slice with zero, got %d", zeroResult)
		}

		if emptyResult != -1 {
			t.Errorf("Expected -1 from empty slice, got %d", emptyResult)
		}
	})

	t.Run("first with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		result := enumerator.FirstOrDefault("default")

		if result != "test" {
			t.Errorf("Expected 'test', got '%s'", result)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkFirstOrDefault(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrDefault(-1)
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrDefault(-1)
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrDefault(-1)
		}
	})
}
