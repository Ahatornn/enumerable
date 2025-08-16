package enumerable

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	t.Run("empty integer enumerator", func(t *testing.T) {
		enumerator := Empty[int]()

		count := 0
		enumerator(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty enumerator, got %d", count)
		}
	})

	t.Run("empty slice enumerator for non-comparable", func(t *testing.T) {
		enumerator := EmptyAny[[]int]()

		count := 0
		enumerator(func(item []int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty enumerator, got %d", count)
		}
	})

	t.Run("empty string enumerator", func(t *testing.T) {
		enumerator := Empty[string]()

		items := []string{}
		enumerator(func(item string) bool {
			items = append(items, item)
			return true
		})

		if len(items) != 0 {
			t.Errorf("Expected empty slice, got %d items: %v", len(items), items)
		}
	})

	t.Run("empty struct enumerator", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		enumerator := Empty[Person]()

		found := false
		enumerator(func(item Person) bool {
			found = true
			return true
		})

		if found {
			t.Error("Expected no items from empty enumerator, but found some")
		}
	})

	t.Run("empty boolean enumerator", func(t *testing.T) {
		enumerator := Empty[bool]()

		called := false
		enumerator(func(item bool) bool {
			called = true
			return true
		})

		if called {
			t.Error("Expected yield function to never be called")
		}
	})

	t.Run("empty float64 enumerator", func(t *testing.T) {
		enumerator := Empty[float64]()

		sum := 0.0
		enumerator(func(item float64) bool {
			sum += item
			return true
		})

		if sum != 0.0 {
			t.Errorf("Expected sum to be 0.0, got %f", sum)
		}
	})

	t.Run("multiple iterations on same empty enumerator", func(t *testing.T) {
		enumerator := Empty[int]()

		// First iteration
		count1 := 0
		enumerator(func(item int) bool {
			count1++
			return true
		})

		// Second iteration
		count2 := 0
		enumerator(func(item int) bool {
			count2++
			return true
		})

		if count1 != 0 || count2 != 0 {
			t.Errorf("Expected both iterations to yield 0 items, got %d and %d", count1, count2)
		}
	})
}

func TestEmptyTypeSafety(t *testing.T) {
	t.Run("type inference works correctly", func(t *testing.T) {
		// Проверяем, что типы выводятся правильно
		intEnum := Empty[int]()
		stringEnum := Empty[string]()
		boolEnum := Empty[bool]()

		// Просто проверяем, что компиляция проходит
		_ = intEnum
		_ = stringEnum
		_ = boolEnum
	})

	t.Run("complex type empty enumerator", func(t *testing.T) {
		type ComplexStruct struct {
			ID     int
			Name   string
			Values *[]int
		}

		enumerator := Empty[ComplexStruct]()

		called := false
		enumerator(func(item ComplexStruct) bool {
			called = true
			return true
		})

		if called {
			t.Error("Expected no calls to yield function")
		}
	})
}

func TestEmptyPerformance(t *testing.T) {
	t.Run("empty enumerator is fast", func(t *testing.T) {
		enumerator := Empty[int]()

		// Должно выполниться мгновенно, так как нет итераций
		enumerator(func(item int) bool {
			t.Error("This should never be called")
			return true
		})

		// Если мы дошли до этой точки - тест прошел
	})
}

// Benchmark для проверки производительности
func BenchmarkEmpty(b *testing.B) {
	b.Run("benchmark empty int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			enumerator := Empty[int]()
			enumerator(func(item int) bool {
				// Этот код никогда не выполнится
				b.Fatal("This should never be called")
				return true
			})
		}
	})

	b.Run("benchmark empty string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			enumerator := Empty[string]()
			enumerator(func(item string) bool {
				// Этот код никогда не выполнится
				b.Fatal("This should never be called")
				return true
			})
		}
	})
}
