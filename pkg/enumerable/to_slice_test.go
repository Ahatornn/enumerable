package enumerable

import (
	"testing"
)

func TestToSlice(t *testing.T) {
	t.Run("convert non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

	t.Run("convert non-empty slice for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

	t.Run("convert single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		if len(result) != 1 {
			t.Fatalf("Expected length 1, got %d", len(result))
		}

		if result[0] != 42 {
			t.Errorf("Expected 42, got %d", result[0])
		}
	})

	t.Run("convert empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected empty slice, got nil")
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("convert nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.ToSlice()

		if result == nil {
			t.Error("Expected empty slice for nil enumerator, got nil")
		}
		if len(result) != 0 {
			t.Errorf("Expected empty slice for nil enumerator, got slice with length %d", len(result))
		}
	})

	t.Run("convert string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

func TestToSliceStruct(t *testing.T) {
	t.Run("convert struct slice", func(t *testing.T) {
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
		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

	t.Run("convert empty struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}

		enumerator := FromSlice(people)
		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected empty slice, got nil")
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})
}

func TestToSliceBoolean(t *testing.T) {
	t.Run("convert boolean slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []bool{true, false, true, false}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %t at index %d, got %t", v, i, result[i])
			}
		}
	})
}

func TestToSliceWithOperations(t *testing.T) {
	t.Run("to slice after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		result := filtered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

	t.Run("to slice after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		result := distinct.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

	t.Run("to slice after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(4)

		result := taken.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

func TestToSliceEdgeCases(t *testing.T) {
	t.Run("to slice with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

	t.Run("to slice with empty strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "", ""})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []string{"", "", ""}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected '%s' at index %d, got '%s'", v, i, result[i])
			}
		}
	})

	t.Run("to slice with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 3)

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

	t.Run("to slice with range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 5) // 1, 2, 3, 4, 5

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
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

// Benchmark для проверки производительности
func BenchmarkToSlice(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToSlice()
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToSlice()
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToSlice()
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToSlice()
		}
	})
}
