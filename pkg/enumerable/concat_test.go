package enumerable

import (
	"testing"
)

func TestConcat(t *testing.T) {
	t.Run("basic concatenation", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{4, 5, 6})

		concatenated := first.Concat(second)

		expected := []int{1, 2, 3, 4, 5, 6}
		actual := []int{}

		concatenated(func(item int) bool {
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

	t.Run("basic concatenation for non-comparable", func(t *testing.T) {
		t.Parallel()

		// Используем slices of slices (не comparable)
		first := FromSliceAny([][]int{{1, 2}, {3, 4}})
		second := FromSliceAny([][]int{{5, 6}, {7, 8}})

		concatenated := first.Concat(second)

		// Собираем результаты
		actual := [][]int{}
		concatenated(func(item []int) bool {
			actual = append(actual, item)
			return true
		})

		// Проверяем количество элементов
		if len(actual) != 4 {
			t.Fatalf("Expected 4 items, got %d", len(actual))
		}

		// Проверяем содержимое
		expected := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
		for i, expectedSlice := range expected {
			if len(actual[i]) != len(expectedSlice) {
				t.Errorf("Expected slice length %d at index %d, got %d", len(expectedSlice), i, len(actual[i]))
				continue
			}
			for j, v := range expectedSlice {
				if actual[i][j] != v {
					t.Errorf("Expected %d at index [%d][%d], got %d", v, i, j, actual[i][j])
				}
			}
		}
	})

	t.Run("concat with empty first", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{})
		second := FromSlice([]int{1, 2, 3})

		concatenated := first.Concat(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		concatenated(func(item int) bool {
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

	t.Run("concat with empty second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{})

		concatenated := first.Concat(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		concatenated(func(item int) bool {
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

	t.Run("concat two empty", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{})
		second := FromSlice([]int{})

		concatenated := first.Concat(second)

		count := 0
		concatenated(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from concatenating two empty enumerations, got %d", count)
		}
	})

	t.Run("concat with nil first", func(t *testing.T) {
		t.Parallel()
		var first Enumerator[int] = nil
		second := FromSlice([]int{1, 2, 3})

		concatenated := first.Concat(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		concatenated(func(item int) bool {
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

	t.Run("concat with nil second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		var second Enumerator[int] = nil

		concatenated := first.Concat(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		concatenated(func(item int) bool {
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

	t.Run("concat two nil", func(t *testing.T) {
		t.Parallel()
		var first Enumerator[int] = nil
		var second Enumerator[int] = nil

		concatenated := first.Concat(second)

		count := 0
		concatenated(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from concatenating two nil enumerations, got %d", count)
		}
	})
}

func TestConcatString(t *testing.T) {
	t.Run("string concatenation", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]string{"hello", "world"})
		second := FromSlice([]string{"foo", "bar"})

		concatenated := first.Concat(second)

		expected := []string{"hello", "world", "foo", "bar"}
		actual := []string{}

		concatenated(func(item string) bool {
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

func TestConcatEarlyTermination(t *testing.T) {
	t.Run("early termination in first enumeration", func(t *testing.T) {
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{6, 7, 8, 9, 10})

		concatenated := first.Concat(second)

		actual := []int{}
		concatenated(func(item int) bool {
			if len(actual) >= 3 {
				return false // Останавливаемся до добавления 4-го элемента
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

	t.Run("early termination in second enumeration", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2})
		second := FromSlice([]int{3, 4, 5, 6, 7})

		concatenated := first.Concat(second)

		actual := []int{}
		stopAt := 4
		concatenated(func(item int) bool {
			actual = append(actual, item)
			return len(actual) < stopAt
		})

		expected := []int{1, 2, 3, 4}
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

func TestConcatStruct(t *testing.T) {
	t.Run("struct concatenation", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		first := FromSlice([]Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		})

		second := FromSlice([]Person{
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
		})

		concatenated := first.Concat(second)

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
		}

		actual := []Person{}
		concatenated(func(item Person) bool {
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

func TestConcatEdgeCases(t *testing.T) {
	t.Run("single element concatenation", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{42})
		second := FromSlice([]int{24})

		concatenated := first.Concat(second)

		expected := []int{42, 24}
		actual := []int{}

		concatenated(func(item int) bool {
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

	t.Run("concat with repeat", func(t *testing.T) {
		t.Parallel()
		first := Repeat(1, 2)
		second := Repeat(2, 3)

		concatenated := first.Concat(second)

		expected := []int{1, 1, 2, 2, 2}
		actual := []int{}

		concatenated(func(item int) bool {
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

// Benchmark для проверки производительности
func BenchmarkConcat(b *testing.B) {
	b.Run("small concatenation", func(b *testing.B) {
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{4, 5, 6})

		for i := 0; i < b.N; i++ {
			concatenated := first.Concat(second)
			concatenated(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large concatenation", func(b *testing.B) {
		first := make([]int, 1000)
		second := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			first[i] = i
			second[i] = i + 1000
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			firstEnum := FromSlice(first)
			secondEnum := FromSlice(second)

			concatenated := firstEnum.Concat(secondEnum)

			var count int
			concatenated(func(item int) bool {
				count++
				return true
			})

			// Убеждаемся, что результат используется
			if count != 2000 {
				b.Fatalf("Expected 2000 items, got %d", count)
			}
		}
	})
}
