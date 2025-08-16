package enumerable

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat integer", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(42, 3)

		expected := []int{42, 42, 42}
		actual := []int{}

		enumerator(func(item int) bool {
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

	t.Run("repeat map", func(t *testing.T) {
		t.Parallel()
		item := map[string]int{"key": 42}
		enumerator := RepeatAny(item, 2)

		actual := []map[string]int{}
		enumerator(func(item map[string]int) bool {
			copy := make(map[string]int)
			for k, v := range item {
				copy[k] = v
			}
			actual = append(actual, copy)
			return true
		})

		if len(actual) != 2 {
			t.Fatalf("Expected length 2, got %d", len(actual))
		}

		for i, m := range actual {
			if len(m) != 1 {
				t.Errorf("Expected map length 1 at index %d, got %d", i, len(m))
				continue
			}
			if val, ok := m["key"]; !ok || val != 42 {
				t.Errorf("Expected map[%s] = %d at index %d, got %v", "key", 42, i, m)
			}
		}
	})

	t.Run("repeat string", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("hello", 4)

		expected := []string{"hello", "hello", "hello", "hello"}
		actual := []string{}

		enumerator(func(item string) bool {
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

	t.Run("zero count", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 0)

		count := 0
		enumerator(func(item string) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items, got %d", count)
		}
	})

	t.Run("one count", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(3.14, 1)

		var result float64
		found := false

		enumerator(func(item float64) bool {
			result = item
			found = true
			return true
		})

		if !found {
			t.Error("Expected to find one element, but found none")
		}

		if result != 3.14 {
			t.Errorf("Expected 3.14, got %f", result)
		}
	})

	t.Run("early termination", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("stop", 10)

		actual := []string{}
		enumerator(func(item string) bool {
			actual = append(actual, item)
			return len(actual) < 3 // Останавливаемся после 3 элементов
		})

		expected := []string{"stop", "stop", "stop"}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
			}
		}
	})

	t.Run("repeat struct", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		person := Person{Name: "Alice", Age: 30}
		enumerator := Repeat(person, 2)

		actual := []Person{}
		enumerator(func(item Person) bool {
			actual = append(actual, item)
			return true
		})

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Alice", Age: 30},
		}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, actual[i])
			}
		}
	})

	t.Run("repeat boolean", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(true, 5)

		actual := []bool{}
		enumerator(func(item bool) bool {
			actual = append(actual, item)
			return true
		})

		expected := []bool{true, true, true, true, true}
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

func TestRepeatEdgeCases(t *testing.T) {
	t.Run("negative count", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(1, -1)

		count := 0
		enumerator(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items with negative count, got %d", count)
		}
	})

	t.Run("repeat zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(0, 3)

		actual := []int{}
		enumerator(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		expected := []int{0, 0, 0}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("repeat empty string", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("", 2)

		actual := []string{}
		enumerator(func(item string) bool {
			actual = append(actual, item)
			return true
		})

		expected := []string{"", ""}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected empty string at index %d, got %s", i, actual[i])
			}
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkRepeat(b *testing.B) {
	b.Run("small repeat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			enumerator := Repeat("test", 10)
			enumerator(func(item string) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large repeat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			enumerator := Repeat(1, 1000)
			enumerator(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("repeat struct", func(b *testing.B) {
		type SimpleStruct struct {
			A, B int
		}
		item := SimpleStruct{A: 1, B: 2}

		for i := 0; i < b.N; i++ {
			enumerator := Repeat(item, 100)
			enumerator(func(item SimpleStruct) bool {
				_ = item
				return true
			})
		}
	})
}
