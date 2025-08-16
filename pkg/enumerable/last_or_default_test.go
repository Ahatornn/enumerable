package enumerable

import (
	"testing"
)

func TestLastOrDefault(t *testing.T) {
	t.Run("last element from non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.LastOrDefault(-1)

		if result != 5 {
			t.Errorf("Expected value 5, got %d", result)
		}
	})

	t.Run("last element from non-empty slice for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		defaultSlice := []int{-1, -2}

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		})

		result := enumerator.LastOrDefault(defaultSlice)

		if len(result) != 2 {
			t.Errorf("Expected last element length 2, got %d", len(result))
		}

		if result[0] != 5 || result[1] != 6 {
			t.Errorf("Expected last element [5,6], got %v", result)
		}
	})

	t.Run("last element from single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.LastOrDefault(-1)

		if result != 42 {
			t.Errorf("Expected value 42, got %d", result)
		}
	})

	t.Run("last element from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.LastOrDefault(-1)

		if result != -1 {
			t.Errorf("Expected default value -1, got %d", result)
		}
	})

	t.Run("last element from nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.LastOrDefault(-1)

		if result != -1 {
			t.Errorf("Expected default value -1 for nil enumerator, got %d", result)
		}
	})

	t.Run("last string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.LastOrDefault("default")

		if result != "go" {
			t.Errorf("Expected value 'go', got '%s'", result)
		}
	})

	t.Run("last empty string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", ""})

		result := enumerator.LastOrDefault("default")

		if result != "" {
			t.Errorf("Expected empty string, got '%s'", result)
		}
	})

	t.Run("last string from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{})

		result := enumerator.LastOrDefault("default")

		if result != "default" {
			t.Errorf("Expected default value 'default', got '%s'", result)
		}
	})
}

func TestLastOrDefaultStruct(t *testing.T) {
	t.Run("last struct element", func(t *testing.T) {
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
		result := enumerator.LastOrDefault(defaultPerson)

		expected := Person{Name: "Charlie", Age: 35}
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	t.Run("last struct from empty slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}
		defaultPerson := Person{Name: "Unknown", Age: 0}

		enumerator := FromSlice(people)
		result := enumerator.LastOrDefault(defaultPerson)

		if result != defaultPerson {
			t.Errorf("Expected default person %+v, got %+v", defaultPerson, result)
		}
	})
}

func TestLastOrDefaultBoolean(t *testing.T) {
	t.Run("last false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		result := enumerator.LastOrDefault(false)

		if result != true {
			t.Errorf("Expected true, got %t", result)
		}
	})

	t.Run("last true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, true, false})

		result := enumerator.LastOrDefault(true)

		if result != false {
			t.Errorf("Expected false, got %t", result)
		}
	})

	t.Run("last boolean from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{})

		result := enumerator.LastOrDefault(true)

		if result != true {
			t.Errorf("Expected default value true, got %t", result)
		}
	})
}

func TestLastOrDefaultWithOperations(t *testing.T) {
	t.Run("last after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		result := filtered.LastOrDefault(-1)

		if result != 6 {
			t.Errorf("Expected value 6, got %d", result)
		}
	})

	t.Run("last after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(5)

		result := taken.LastOrDefault(-1)

		if result != 5 {
			t.Errorf("Expected value 5, got %d", result)
		}
	})
}

func TestLastOrDefaultEdgeCases(t *testing.T) {
	t.Run("last zero value with non-zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 0})

		result := enumerator.LastOrDefault(-1)

		if result != 0 {
			t.Errorf("Expected value 0, got %d", result)
		}
	})

	t.Run("last zero value with zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 0})

		result := enumerator.LastOrDefault(0)

		if result != 0 {
			t.Errorf("Expected value 0, got %d", result)
		}
	})

	t.Run("empty slice with zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.LastOrDefault(0)

		if result != 0 {
			t.Errorf("Expected default value 0, got %d", result)
		}
	})

	t.Run("distinguishing zero value from default", func(t *testing.T) {
		t.Parallel()
		withZero := FromSlice([]int{1, 2, 0})
		empty := FromSlice([]int{})

		zeroResult := withZero.LastOrDefault(-1)
		emptyResult := empty.LastOrDefault(-1)

		if zeroResult != 0 {
			t.Errorf("Expected 0 from slice with zero, got %d", zeroResult)
		}

		if emptyResult != -1 {
			t.Errorf("Expected -1 from empty slice, got %d", emptyResult)
		}
	})

	t.Run("last with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		result := enumerator.LastOrDefault("default")

		if result != "test" {
			t.Errorf("Expected 'test', got '%s'", result)
		}
	})

	t.Run("last with range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 5) // 1, 2, 3, 4, 5

		result := enumerator.LastOrDefault(-1)

		if result != 5 {
			t.Errorf("Expected value 5, got %d", result)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkLastOrDefault(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrDefault(-1)
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrDefault(-1)
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrDefault(-1)
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrDefault(-1)
		}
	})
}
