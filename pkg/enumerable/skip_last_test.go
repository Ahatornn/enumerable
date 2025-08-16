package enumerable

import (
	"testing"
)

func TestSkipLast(t *testing.T) {
	t.Run("basic skip last", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		skipped := enumerator.SkipLast(2)

		expected := []int{1, 2, 3}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("basic skip last for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		})

		skipped := enumerator.SkipLast(2)

		expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
		actual := [][]int{}

		skipped(func(item []int) bool {
			copy := make([]int, len(item))
			for i, v := range item {
				copy[i] = v
			}
			actual = append(actual, copy)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if len(actual[i]) != len(v) {
				t.Errorf("Expected slice length %d at index %d, got %d", len(v), i, len(actual[i]))
				continue
			}
			for j, val := range v {
				if actual[i][j] != val {
					t.Errorf("Expected %d at index [%d][%d], got %d", val, i, j, actual[i][j])
				}
			}
		}
	})

	t.Run("skip last zero elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		skipped := enumerator.SkipLast(0)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("skip more than available", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		skipped := enumerator.SkipLast(5)

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when skipping more than available, got %d", count)
		}
	})

	t.Run("skip exactly all elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		skipped := enumerator.SkipLast(5)

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when skipping exactly all elements, got %d", count)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		skipped := enumerator.SkipLast(3)

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty slice, got %d", count)
		}
	})

	t.Run("nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		skipped := enumerator.SkipLast(3)

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator, got %d", count)
		}
	})
}

func TestSkipLastString(t *testing.T) {
	t.Run("skip last strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "b", "c", "d", "e"})

		skipped := enumerator.SkipLast(2)

		expected := []string{"a", "b", "c"}
		actual := []string{}

		skipped(func(item string) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
			}
		}
	})
}

func TestSkipLastEarlyTermination(t *testing.T) {
	t.Run("early termination after skip last", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})

		skipped := enumerator.SkipLast(2)

		actual := []int{}
		skipped(func(item int) bool {
			if len(actual) >= 3 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{1, 2, 3}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})
}

func TestSkipLastStruct(t *testing.T) {
	t.Run("skip last structs", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
			{Name: "Eve", Age: 32},
		}

		enumerator := FromSlice(people)
		skipped := enumerator.SkipLast(2)

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}

		actual := []Person{}
		skipped(func(item Person) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, actual[i])
			}
		}
	})
}

func TestSkipLastEdgeCases(t *testing.T) {
	t.Run("single element skip zero", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		skipped := enumerator.SkipLast(0)

		expected := []int{42}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("single element skip one", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		skipped := enumerator.SkipLast(1)

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when skipping single element, got %d", count)
		}
	})

	t.Run("two elements skip one", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{10, 20})

		skipped := enumerator.SkipLast(1)

		expected := []int{10}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("negative skip count", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		skipped := enumerator.SkipLast(-1)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})
}

func TestSkipLastBoolean(t *testing.T) {
	t.Run("skip last booleans", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		skipped := enumerator.SkipLast(2)

		expected := []bool{true, false, true}
		actual := []bool{}

		skipped(func(item bool) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %t at index %d, got %t", v, i, actual[i])
			}
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkSkipLast(b *testing.B) {
	b.Run("small skip last", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			skipped := enumerator.SkipLast(10)
			skipped(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large skip last", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			skipped := enumerator.SkipLast(1000)
			skipped(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
