package enumerable

import (
	"testing"
)

func TestTakeLast(t *testing.T) {
	t.Run("basic take last", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})

		taken := enumerator.TakeLast(3)

		expected := []int{6, 7, 8}
		actual := []int{}

		taken(func(item int) bool {
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

	t.Run("basic take last for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
			{11, 12},
			{13, 14},
			{15, 16},
		})

		taken := enumerator.TakeLast(3)

		expected := [][]int{
			{11, 12},
			{13, 14},
			{15, 16},
		}
		actual := [][]int{}

		taken(func(item []int) bool {
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

	t.Run("take last zero elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		taken := enumerator.TakeLast(0)

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when taking zero elements, got %d", count)
		}
	})

	t.Run("take negative number", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		taken := enumerator.TakeLast(-1)

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when taking negative number, got %d", count)
		}
	})

	t.Run("take more than available", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		taken := enumerator.TakeLast(5)

		expected := []int{1, 2, 3}
		actual := []int{}

		taken(func(item int) bool {
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

	t.Run("take exactly all elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		taken := enumerator.TakeLast(5)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		taken(func(item int) bool {
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

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		taken := enumerator.TakeLast(3)

		count := 0
		taken(func(item int) bool {
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

		taken := enumerator.TakeLast(3)

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator, got %d", count)
		}
	})
}

func TestTakeLastString(t *testing.T) {
	t.Run("take last strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "b", "c", "d", "e"})

		taken := enumerator.TakeLast(3)

		expected := []string{"c", "d", "e"}
		actual := []string{}

		taken(func(item string) bool {
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

func TestTakeLastEarlyTermination(t *testing.T) {
	t.Run("early termination by consumer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		taken := enumerator.TakeLast(3)

		actual := []int{}
		taken(func(item int) bool {
			if len(actual) >= 2 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{3, 4}
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

func TestTakeLastStruct(t *testing.T) {
	t.Run("take last structs", func(t *testing.T) {
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
		taken := enumerator.TakeLast(3)

		expected := []Person{
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
			{Name: "Eve", Age: 32},
		}

		actual := []Person{}
		taken(func(item Person) bool {
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

func TestTakeLastEdgeCases(t *testing.T) {
	t.Run("single element take one", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		taken := enumerator.TakeLast(1)

		expected := []int{42}
		actual := []int{}

		taken(func(item int) bool {
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

	t.Run("single element take zero", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		taken := enumerator.TakeLast(0)

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when taking zero elements, got %d", count)
		}
	})

	t.Run("large take from small collection", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		taken := enumerator.TakeLast(1000)

		expected := []int{1, 2, 3}
		actual := []int{}

		taken(func(item int) bool {
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

	t.Run("take last with ring buffer wraparound", func(t *testing.T) {
		t.Parallel()
		// Test case where ring buffer wraps around
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		taken := enumerator.TakeLast(5)

		expected := []int{95, 96, 97, 98, 99}
		actual := []int{}

		taken(func(item int) bool {
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

func TestTakeLastBoolean(t *testing.T) {
	t.Run("take last booleans", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		taken := enumerator.TakeLast(3)

		expected := []bool{true, false, true}
		actual := []bool{}

		taken(func(item bool) bool {
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
func BenchmarkTakeLast(b *testing.B) {
	b.Run("small take last", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			taken := enumerator.TakeLast(10)
			taken(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large take last", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			taken := enumerator.TakeLast(1000)
			taken(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
