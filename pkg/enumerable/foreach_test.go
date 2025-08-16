package enumerable

import (
	"testing"
)

func TestForEach(t *testing.T) {
	t.Run("for each element in non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		var results []int
		enumerator.ForEach(func(n int) {
			results = append(results, n)
		})

		expected := []int{1, 2, 3, 4, 5}
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		for i, v := range expected {
			if results[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, results[i])
			}
		}
	})

	t.Run("for each element in non-empty slice for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		// Используем slices of slices (не comparable)
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		})

		var results [][]int
		enumerator.ForEach(func(slice []int) {
			results = append(results, slice)
		})

		expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		for i, v := range expected {
			if len(results[i]) != len(v) {
				t.Errorf("Expected slice length %d at index %d, got %d", len(v), i, len(results[i]))
				continue
			}
			for j, val := range v {
				if results[i][j] != val {
					t.Errorf("Expected %d at index [%d][%d], got %d", val, i, j, results[i][j])
				}
			}
		}
	})

	t.Run("for each element in single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		var results []int
		enumerator.ForEach(func(n int) {
			results = append(results, n)
		})

		expected := []int{42}
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		if results[0] != expected[0] {
			t.Errorf("Expected %d, got %d", expected[0], results[0])
		}
	})

	t.Run("for each element in empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		callCount := 0
		enumerator.ForEach(func(n int) {
			callCount++
		})

		if callCount != 0 {
			t.Errorf("Expected 0 calls for empty slice, got %d", callCount)
		}
	})

	t.Run("for each element in nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		callCount := 0
		enumerator.ForEach(func(n int) {
			callCount++
		})

		if callCount != 0 {
			t.Errorf("Expected 0 calls for nil enumerator, got %d", callCount)
		}
	})

	t.Run("for each string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		var results []string
		enumerator.ForEach(func(s string) {
			results = append(results, s)
		})

		expected := []string{"hello", "world", "go"}
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		for i, v := range expected {
			if results[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, results[i])
			}
		}
	})
}

func TestForEachStruct(t *testing.T) {
	t.Run("for each struct element", func(t *testing.T) {
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

		var results []Person
		enumerator.ForEach(func(p Person) {
			results = append(results, p)
		})

		expected := people
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		for i, v := range expected {
			if results[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, results[i])
			}
		}
	})
}

func TestForEachBoolean(t *testing.T) {
	t.Run("for each boolean element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false})

		var results []bool
		enumerator.ForEach(func(b bool) {
			results = append(results, b)
		})

		expected := []bool{true, false, true, false}
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		for i, v := range expected {
			if results[i] != v {
				t.Errorf("Expected %t at index %d, got %t", v, i, results[i])
			}
		}
	})
}

func TestForEachSideEffects(t *testing.T) {
	t.Run("side effects with external variable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		sum := 0
		count := 0

		enumerator.ForEach(func(n int) {
			sum += n
			count++
		})

		if sum != 15 {
			t.Errorf("Expected sum 15, got %d", sum)
		}

		if count != 5 {
			t.Errorf("Expected count 5, got %d", count)
		}
	})

	t.Run("side effects with string concatenation", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "b", "c"})

		result := ""
		enumerator.ForEach(func(s string) {
			result += s
		})

		if result != "abc" {
			t.Errorf("Expected 'abc', got '%s'", result)
		}
	})
}

func TestForEachWithOperations(t *testing.T) {
	t.Run("for each after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		var results []int
		filtered.ForEach(func(n int) {
			results = append(results, n)
		})

		expected := []int{2, 4, 6}
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		for i, v := range expected {
			if results[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, results[i])
			}
		}
	})

	t.Run("for each after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		var results []int
		distinct.ForEach(func(n int) {
			results = append(results, n)
		})

		expected := []int{1, 2, 3, 4}
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		for i, v := range expected {
			if results[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, results[i])
			}
		}
	})
}

func TestForEachEdgeCases(t *testing.T) {
	t.Run("for each with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0})

		count := 0
		sum := 0
		enumerator.ForEach(func(n int) {
			count++
			sum += n
		})

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}

		if sum != 0 {
			t.Errorf("Expected sum 0, got %d", sum)
		}
	})

	t.Run("for each with empty strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "", ""})

		count := 0
		enumerator.ForEach(func(s string) {
			count++
		})

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}
	})

	t.Run("for each with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 3)

		var results []string
		enumerator.ForEach(func(s string) {
			results = append(results, s)
		})

		expected := []string{"test", "test", "test"}
		if len(results) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(results))
		}

		for i, v := range expected {
			if results[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, results[i])
			}
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkForEach(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			enumerator.ForEach(func(n int) {
				_ = n
			})
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			enumerator.ForEach(func(n int) {
				_ = n
			})
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			enumerator.ForEach(func(n int) {
				_ = n
			})
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			enumerator.ForEach(func(n int) {
				_ = n
			})
		}
	})
}
