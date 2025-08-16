package enumerable

import (
	"testing"
)

func TestLastOrNil(t *testing.T) {
	t.Run("last element from non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 5 {
			t.Errorf("Expected value 5, got %d", *result)
		}
	})

	t.Run("last element from non-empty slice for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if len(*result) != 2 {
			t.Errorf("Expected last element length 2, got %d", len(*result))
		}

		if (*result)[0] != 5 || (*result)[1] != 6 {
			t.Errorf("Expected last element [5,6], got %v", *result)
		}
	})

	t.Run("last element from single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 42 {
			t.Errorf("Expected value 42, got %d", *result)
		}
	})

	t.Run("last element from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil for empty slice, got pointer to %d", *result)
		}
	})

	t.Run("last element from nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil for nil enumerator, got pointer to %d", *result)
		}
	})

	t.Run("last string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != "go" {
			t.Errorf("Expected value 'go', got '%s'", *result)
		}
	})

	t.Run("last empty string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", ""})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != "" {
			t.Errorf("Expected empty string, got '%s'", *result)
		}
	})
}

func TestLastOrNilStruct(t *testing.T) {
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

		enumerator := FromSlice(people)
		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		expected := Person{Name: "Charlie", Age: 35}
		if *result != expected {
			t.Errorf("Expected %+v, got %+v", expected, *result)
		}
	})

	t.Run("last struct from empty slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}

		enumerator := FromSlice(people)
		result := enumerator.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil for empty struct slice, got %+v", *result)
		}
	})
}

func TestLastOrNilBoolean(t *testing.T) {
	t.Run("last true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, true, false})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != false {
			t.Errorf("Expected false, got %t", *result)
		}
	})

	t.Run("last false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != true {
			t.Errorf("Expected true, got %t", *result)
		}
	})
}

func TestLastOrNilWithOperations(t *testing.T) {
	t.Run("last after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		result := filtered.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 6 {
			t.Errorf("Expected value 6, got %d", *result)
		}
	})

	t.Run("last after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(5)

		result := taken.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 5 {
			t.Errorf("Expected value 5, got %d", *result)
		}
	})
}

func TestLastOrNilEdgeCases(t *testing.T) {
	t.Run("last zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 0})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 0 {
			t.Errorf("Expected value 0, got %d", *result)
		}
	})

	t.Run("distinguishing nil from zero value", func(t *testing.T) {
		t.Parallel()
		empty := FromSlice([]int{})
		withZero := FromSlice([]int{0})

		emptyResult := empty.LastOrNil()
		zeroResult := withZero.LastOrNil()

		if emptyResult != nil {
			t.Errorf("Expected nil for empty slice, got pointer")
		}

		if zeroResult == nil {
			t.Errorf("Expected pointer for slice with zero value, got nil")
		} else if *zeroResult != 0 {
			t.Errorf("Expected 0, got %d", *zeroResult)
		}
	})

	t.Run("last with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != "test" {
			t.Errorf("Expected 'test', got '%s'", *result)
		}
	})

	t.Run("last with range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 5) // 1, 2, 3, 4, 5

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 5 {
			t.Errorf("Expected value 5, got %d", *result)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkLastOrNil(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrNil()
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrNil()
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrNil()
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrNil()
		}
	})
}
