package enumerable

import (
	"testing"
)

func TestDefaultIfEmpty(t *testing.T) {
	t.Run("non-empty enumeration", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 2, 3})

		result := source.DefaultIfEmpty(-1)

		expected := []int{1, 2, 3}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("empty enumeration", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{})

		result := source.DefaultIfEmpty(-1)

		expected := []int{-1}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("nil enumeration", func(t *testing.T) {
		t.Parallel()
		var source Enumerator[int] = nil

		result := source.DefaultIfEmpty(42)

		expected := []int{42}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("single element enumeration", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{42})

		result := source.DefaultIfEmpty(-1)

		expected := []int{42}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("string enumeration with default", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]string{})

		result := source.DefaultIfEmpty("default")

		expected := []string{"default"}
		actual := []string{}

		result(func(item string) bool {
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

	t.Run("struct enumeration with default", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		source := FromSlice([]Person{})

		defaultPerson := Person{Name: "Unknown", Age: 0}
		result := source.DefaultIfEmpty(defaultPerson)

		expected := []Person{defaultPerson}
		actual := []Person{}

		result(func(item Person) bool {
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

func TestDefaultIfEmptyAny(t *testing.T) {
	t.Run("non-comparable slice enumeration", func(t *testing.T) {
		t.Parallel()
		source := FromSliceAny([][]int{{1, 2}, {3, 4}})

		defaultValue := []int{-1, -1}
		result := source.DefaultIfEmpty(defaultValue)

		expected := [][]int{{1, 2}, {3, 4}}
		actual := [][]int{}

		result(func(item []int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

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

	t.Run("empty non-comparable slice enumeration", func(t *testing.T) {
		t.Parallel()
		source := FromSliceAny([][]int{})

		defaultValue := []int{-1, -1}
		result := source.DefaultIfEmpty(defaultValue)

		expected := [][]int{{-1, -1}}
		actual := [][]int{}

		result(func(item []int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

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

	t.Run("nil non-comparable enumeration", func(t *testing.T) {
		t.Parallel()
		var source EnumeratorAny[[]int] = nil

		defaultValue := []int{0, 0}
		result := source.DefaultIfEmpty(defaultValue)

		expected := [][]int{{0, 0}}
		actual := [][]int{}

		result(func(item []int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

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
}

func TestDefaultIfEmptyEarlyTermination(t *testing.T) {
	t.Run("early termination with non-empty source", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 2, 3, 4, 5})

		result := source.DefaultIfEmpty(-1)

		actual := []int{}
		result(func(item int) bool {
			if len(actual) >= 2 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{1, 2}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("early termination with empty source", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{})

		result := source.DefaultIfEmpty(-1)

		count := 0
		result(func(item int) bool {
			count++
			return false
		})

		if count != 1 {
			t.Errorf("Expected exactly 1 item (default value), got %d", count)
		}
	})

	t.Run("early termination with nil source", func(t *testing.T) {
		t.Parallel()
		var source Enumerator[int] = nil

		result := source.DefaultIfEmpty(42)

		count := 0
		result(func(item int) bool {
			count++
			return false
		})

		if count != 1 {
			t.Errorf("Expected exactly 1 item (default value), got %d", count)
		}
	})
}

func TestDefaultIfEmptyEdgeCases(t *testing.T) {
	t.Run("zero as default value", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{})

		result := source.DefaultIfEmpty(0)

		expected := []int{0}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("multiple default if empty calls", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{})

		result := source.DefaultIfEmpty(-1).DefaultIfEmpty(-2)

		expected := []int{-1}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("default if empty after elements added", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1}).DefaultIfEmpty(-1)

		expected := []int{1}
		actual := []int{}

		source(func(item int) bool {
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
func BenchmarkDefaultIfEmpty(b *testing.B) {
	b.Run("non-empty enumeration", func(b *testing.B) {
		source := FromSlice([]int{1, 2, 3, 4, 5})

		for i := 0; i < b.N; i++ {
			result := source.DefaultIfEmpty(-1)
			result(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		source := FromSlice([]int{})

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			result := source.DefaultIfEmpty(-1)

			count := 0
			result(func(item int) bool {
				count++
				return true
			})

			if count != 1 {
				b.Fatalf("Expected 1 item, got %d", count)
			}
		}
	})

	b.Run("large non-empty enumeration", func(b *testing.B) {
		data := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			data[i] = i
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			source := FromSlice(data)
			result := source.DefaultIfEmpty(-1)

			var count int
			result(func(item int) bool {
				count++
				return true
			})

			// Убеждаемся, что результат используется
			if count != 10000 {
				b.Fatalf("Expected 10000 items, got %d", count)
			}
		}
	})
}
