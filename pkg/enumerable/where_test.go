package enumerable

import (
	"testing"
)

func TestWhere(t *testing.T) {
	t.Run("filter even numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})

		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		expected := []int{2, 4, 6, 8}
		actual := []int{}

		filtered(func(item int) bool {
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

	t.Run("filter even numbers for non-comparable", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1},
			{2},
			{3},
			{4},
			{5},
			{6},
			{7},
			{8},
		})

		filtered := enumerator.Where(func(slice []int) bool {
			return len(slice) > 0 && slice[0]%2 == 0
		})

		expected := [][]int{
			{2},
			{4},
			{6},
			{8},
		}
		actual := [][]int{}

		filtered(func(item []int) bool {
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

	t.Run("filter odd numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})

		filtered := enumerator.Where(func(n int) bool { return n%2 == 1 })

		expected := []int{1, 3, 5, 7}
		actual := []int{}

		filtered(func(item int) bool {
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

	t.Run("filter with no matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 3, 5, 7})

		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		count := 0
		filtered(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when no matches, got %d", count)
		}
	})

	t.Run("filter with all matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 4, 6, 8})

		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		expected := []int{2, 4, 6, 8}
		actual := []int{}

		filtered(func(item int) bool {
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

		filtered := enumerator.Where(func(n int) bool { return n > 0 })

		count := 0
		filtered(func(item int) bool {
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
		filtered := enumerator.Where(func(n int) bool { return n > 0 })
		count := 0
		filtered(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 elements from Where on nil enumerator, got %d", count)
		}
	})
}

func TestWhereString(t *testing.T) {
	t.Run("filter strings by length", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "hello", "hi", "world", "golang"})

		filtered := enumerator.Where(func(s string) bool { return len(s) > 3 })

		expected := []string{"hello", "world", "golang"}
		actual := []string{}

		filtered(func(item string) bool {
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

	t.Run("filter strings by prefix", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"apple", "banana", "avocado", "cherry", "apricot"})

		filtered := enumerator.Where(func(s string) bool { return len(s) > 0 && s[0] == 'a' })

		expected := []string{"apple", "avocado", "apricot"}
		actual := []string{}

		filtered(func(item string) bool {
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

func TestWhereEarlyTermination(t *testing.T) {
	t.Run("early termination by consumer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 4, 6, 8, 10, 12, 14, 16})

		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		actual := []int{}
		filtered(func(item int) bool {
			if len(actual) >= 3 {
				return false
			}
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

	t.Run("early termination skips remaining elements", func(t *testing.T) {
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

		filtered := enum.Where(func(n int) bool { return n%2 == 0 })

		actual := []int{}
		filtered(func(item int) bool {
			if len(actual) >= 2 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{2, 4}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		if callCount >= 100 {
			t.Errorf("Expected early termination, but processed %d items", callCount)
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})
}

func TestWhereStruct(t *testing.T) {
	t.Run("filter structs by field", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 20},
			{Name: "Eve", Age: 40},
		}

		enumerator := FromSlice(people)
		filtered := enumerator.Where(func(p Person) bool { return p.Age >= 30 })

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Charlie", Age: 35},
			{Name: "Eve", Age: 40},
		}

		actual := []Person{}
		filtered(func(item Person) bool {
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

func TestWhereEdgeCases(t *testing.T) {
	t.Run("single element matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		filtered := enumerator.Where(func(n int) bool { return n == 42 })

		expected := []int{42}
		actual := []int{}

		filtered(func(item int) bool {
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

	t.Run("single element doesn't match", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		filtered := enumerator.Where(func(n int) bool { return n != 42 })

		count := 0
		filtered(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when single element doesn't match, got %d", count)
		}
	})

	t.Run("predicate always true", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		filtered := enumerator.Where(func(n int) bool { return true })

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		filtered(func(item int) bool {
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

	t.Run("predicate always false", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		filtered := enumerator.Where(func(n int) bool { return false })

		count := 0
		filtered(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when predicate always false, got %d", count)
		}
	})
}

func TestWhereBoolean(t *testing.T) {
	t.Run("filter booleans", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		filtered := enumerator.Where(func(b bool) bool { return b })

		expected := []bool{true, true, true}
		actual := []bool{}

		filtered(func(item bool) bool {
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
func BenchmarkWhere(b *testing.B) {
	b.Run("filter half elements", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })
			filtered(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("filter all elements", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			filtered := enumerator.Where(func(n int) bool { return true })
			filtered(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("filter no elements", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			filtered := enumerator.Where(func(n int) bool { return false })
			filtered(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
