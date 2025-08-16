package enumerable

import (
	"testing"
)

func TestTakeWhile(t *testing.T) {
	t.Run("basic take while", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 1, 2})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 4 })

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

	t.Run("basic take while for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{2, 3},
			{3, 4, 5},
			{4, 5},
			{5, 6},
			{1, 2},
			{2, 3},
		})

		taken := enumerator.TakeWhile(func(slice []int) bool {
			return len(slice) < 3
		})

		expected := [][]int{
			{1, 2},
			{2, 3},
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

	t.Run("take while even numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 4, 6, 7, 8, 10, 11, 12})

		taken := enumerator.TakeWhile(func(n int) bool { return n%2 == 0 })

		expected := []int{2, 4, 6}
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

	t.Run("predicate never matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 4, 5, 6})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when predicate never matches, got %d", count)
		}
	})

	t.Run("predicate always matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 10 })

		expected := []int{1, 2, 3, 4}
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

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

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

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

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

func TestTakeWhileString(t *testing.T) {
	t.Run("take while strings start with letter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"apple", "banana", "123start", "cherry", "456end"})

		taken := enumerator.TakeWhile(func(s string) bool {
			return len(s) > 0 && s[0] >= 'a' && s[0] <= 'z'
		})

		expected := []string{"apple", "banana"}
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

func TestTakeWhileEarlyTermination(t *testing.T) {
	t.Run("early termination by consumer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 8 })

		actual := []int{}
		taken(func(item int) bool {
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

	t.Run("early termination by predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 4, 6, 7, 8, 10})

		taken := enumerator.TakeWhile(func(n int) bool { return n%2 == 0 })

		actual := []int{}
		taken(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		expected := []int{2, 4, 6}
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

func TestTakeWhileStruct(t *testing.T) {
	t.Run("take while struct field condition", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 20},
			{Name: "Eve", Age: 40},
		}

		enumerator := FromSlice(people)
		taken := enumerator.TakeWhile(func(p Person) bool { return p.Age < 35 })

		expected := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
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

func TestTakeWhileEdgeCases(t *testing.T) {
	t.Run("single element matches predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		expected := []int{2}
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

	t.Run("single element doesn't match predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when single element doesn't match, got %d", count)
		}
	})

	t.Run("predicate returns true then false", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 1, 3, 1, 4})

		taken := enumerator.TakeWhile(func(n int) bool { return n > 1 })

		expected := []int{2}
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

func TestTakeWhileBoolean(t *testing.T) {
	t.Run("take while true values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, false, true, false})

		taken := enumerator.TakeWhile(func(b bool) bool { return b })

		expected := []bool{true, true}
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
func BenchmarkTakeWhile(b *testing.B) {
	b.Run("small take while", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			taken := enumerator.TakeWhile(func(n int) bool { return n < 500 })
			taken(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("no take while", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i + 1000
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			taken := enumerator.TakeWhile(func(n int) bool { return n < 500 })
			taken(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
