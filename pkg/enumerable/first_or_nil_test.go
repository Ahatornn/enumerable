package enumerable

import (
	"testing"
)

func TestFirstOrNil(t *testing.T) {
	t.Run("first element from non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != 1 {
			t.Errorf("Expected value 1, got %d", *result)
		}
	})

	t.Run("first element from non-empty for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if len(*result) != 2 {
			t.Errorf("Expected first element length 2, got %d", len(*result))
		}

		if (*result)[0] != 1 || (*result)[1] != 2 {
			t.Errorf("Expected first element [1,2], got %v", *result)
		}
	})

	t.Run("first element from single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != 42 {
			t.Errorf("Expected value 42, got %d", *result)
		}
	})

	t.Run("first element from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil for empty slice, got pointer to %d", *result)
		}
	})

	t.Run("first element from nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil for nil enumerator, got pointer to %d", *result)
		}
	})

	t.Run("first string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != "hello" {
			t.Errorf("Expected value 'hello', got '%s'", *result)
		}
	})

	t.Run("first empty string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "world", "go"})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != "" {
			t.Errorf("Expected empty string, got '%s'", *result)
		}
	})
}

func TestFirstOrNilStruct(t *testing.T) {
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

		enumerator := FromSlice(people)
		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		expected := Person{Name: "Alice", Age: 30}
		if *result != expected {
			t.Errorf("Expected %+v, got %+v", expected, *result)
		}
	})

	t.Run("first struct from empty slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}

		enumerator := FromSlice(people)
		result := enumerator.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil for empty struct slice, got %+v", *result)
		}
	})
}

func TestFirstOrNilBoolean(t *testing.T) {
	t.Run("first true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != true {
			t.Errorf("Expected true, got %t", *result)
		}
	})

	t.Run("first false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, true, false})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != false {
			t.Errorf("Expected false, got %t", *result)
		}
	})
}

func TestFirstOrNilEarlyTermination(t *testing.T) {
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
		result := enum.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != 1 {
			t.Errorf("Expected value 1, got %d", *result)
		}

		if callCount != 1 {
			t.Errorf("Expected exactly 1 call, got %d", callCount)
		}
	})
}

func TestFirstOrNilEdgeCases(t *testing.T) {
	t.Run("first zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 1, 2, 3})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != 0 {
			t.Errorf("Expected value 0, got %d", *result)
		}
	})

	t.Run("distinguishing nil from zero value", func(t *testing.T) {
		t.Parallel()
		empty := FromSlice([]int{})
		withZero := FromSlice([]int{0})

		emptyResult := empty.FirstOrNil()
		zeroResult := withZero.FirstOrNil()

		// emptyResult должен быть nil
		if emptyResult != nil {
			t.Errorf("Expected nil for empty slice, got pointer")
		}

		if zeroResult == nil {
			t.Errorf("Expected pointer for slice with zero value, got nil")
		} else if *zeroResult != 0 {
			t.Errorf("Expected 0, got %d", *zeroResult)
		}
	})

	t.Run("first with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != "test" {
			t.Errorf("Expected 'test', got '%s'", *result)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkFirstOrNil(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrNil()
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrNil()
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrNil()
		}
	})
}
