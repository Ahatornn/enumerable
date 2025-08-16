package enumerable

import (
	"testing"
)

func TestSkipWhile(t *testing.T) {
	t.Run("basic skip while", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 1, 2})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		expected := []int{3, 4, 5, 1, 2}
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

	t.Run("basic skip while for non-comparable", func(t *testing.T) {
		t.Parallel()

		type Person struct {
			Name string
			Age  int
		}

		enumerator := FromSliceAny([]Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 35},
			{Name: "David", Age: 40},
			{Name: "Eve", Age: 45},
			{Name: "Frank", Age: 25},
			{Name: "Grace", Age: 30},
		})

		skipped := enumerator.SkipWhile(func(p Person) bool {
			return p.Age < 35
		})

		expected := []Person{
			{Name: "Charlie", Age: 35},
			{Name: "David", Age: 40},
			{Name: "Eve", Age: 45},
			{Name: "Frank", Age: 25},
			{Name: "Grace", Age: 30},
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
			if actual[i].Name != v.Name || actual[i].Age != v.Age {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, actual[i])
			}
		}
	})

	t.Run("skip while even numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 4, 6, 7, 8, 10, 11, 12})

		skipped := enumerator.SkipWhile(func(n int) bool { return n%2 == 0 })

		expected := []int{7, 8, 10, 11, 12}
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

	t.Run("predicate never matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 4, 5, 6})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		expected := []int{3, 4, 5, 6}
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

	t.Run("predicate always matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 10 })

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when predicate always matches, got %d", count)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

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

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

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

func TestSkipWhileString(t *testing.T) {
	t.Run("skip while strings start with letter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"apple", "banana", "123start", "cherry", "456end"})

		skipped := enumerator.SkipWhile(func(s string) bool {
			return len(s) > 0 && s[0] >= 'a' && s[0] <= 'z'
		})

		expected := []string{"123start", "cherry", "456end"}
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

func TestSkipWhileEarlyTermination(t *testing.T) {
	t.Run("early termination after skip while", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 5 })

		actual := []int{}
		skipped(func(item int) bool {
			if len(actual) >= 3 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{5, 6, 7}
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

func TestSkipWhileStruct(t *testing.T) {
	t.Run("skip while struct field condition", func(t *testing.T) {
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
		skipped := enumerator.SkipWhile(func(p Person) bool { return p.Age < 35 })

		expected := []Person{
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 20},
			{Name: "Eve", Age: 40},
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

func TestSkipWhileEdgeCases(t *testing.T) {
	t.Run("single element matches predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when single element matches, got %d", count)
		}
	})

	t.Run("single element doesn't match predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		expected := []int{5}
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

	t.Run("predicate returns false then true", func(t *testing.T) {
		t.Parallel()
		// Important: once we stop skipping, we yield everything after
		enumerator := FromSlice([]int{2, 1, 3, 1, 4}) // 2 matches, 1 doesn't, so yield 1 and all after

		skipped := enumerator.SkipWhile(func(n int) bool { return n > 1 })

		expected := []int{1, 3, 1, 4}
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

func TestSkipWhileBoolean(t *testing.T) {
	t.Run("skip while true values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, false, true, false})

		skipped := enumerator.SkipWhile(func(b bool) bool { return b })

		expected := []bool{false, true, false}
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
func BenchmarkSkipWhile(b *testing.B) {
	b.Run("small skip while", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			skipped := enumerator.SkipWhile(func(n int) bool { return n < 500 })
			skipped(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("no skip while", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i + 1000 // All items >= 1000
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			skipped := enumerator.SkipWhile(func(n int) bool { return n < 500 })
			skipped(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
