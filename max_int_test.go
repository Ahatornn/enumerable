package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMaxInt(t *testing.T) {
	t.Run("max int from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})

	t.Run("max int from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, -2, -8, -1, -9})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != -1 {
			t.Errorf("Expected max -1, got %d", max)
		}
	})

	t.Run("max int from mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, 2, -8, 1, 0})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 2 {
			t.Errorf("Expected max 2, got %d", max)
		}
	})

	t.Run("max int single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 42 {
			t.Errorf("Expected max 42, got %d", max)
		}
	})

	t.Run("max int empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		max, ok := enumerator.MaxInt(selector.Int)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for empty slice, got %d", max)
		}
	})

	t.Run("max int nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		max, ok := enumerator.MaxInt(selector.Int)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for nil enumerator, got %d", max)
		}
	})

	t.Run("max int with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		max, ok := enumerator.MaxInt(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for nil keySelector, got %d", max)
		}
	})

	t.Run("max int with custom key selector", func(t *testing.T) {
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
		max, ok := enumerator.MaxInt(func(p Person) int { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 35 {
			t.Errorf("Expected max age 35, got %d", max)
		}
	})

	t.Run("max int with string length key selector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hi", "world", "go"})

		max, ok := enumerator.MaxInt(func(s string) int { return len(s) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 5 {
			t.Errorf("Expected max length 5, got %d", max)
		}
	})

	t.Run("max int with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 9, 8, 1, 9})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})

	t.Run("max int with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 0, 8, 0, 9})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})
}

func TestMax64Int(t *testing.T) {
	t.Run("max int64 from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{5, 2, 8, 1, 9})

		max, ok := enumerator.MaxInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})

	t.Run("max int64 from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{-5, -2, -8, -1, -9})

		max, ok := enumerator.MaxInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != -1 {
			t.Errorf("Expected max -1, got %d", max)
		}
	})

	t.Run("max int64 from mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{-5, 2, -8, 1, 0})

		max, ok := enumerator.MaxInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 2 {
			t.Errorf("Expected max 2, got %d", max)
		}
	})

	t.Run("max int64 single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{42})

		max, ok := enumerator.MaxInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 42 {
			t.Errorf("Expected max 42, got %d", max)
		}
	})

	t.Run("max int64 empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{})

		max, ok := enumerator.MaxInt64(selector.Int64)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for empty slice, got %d", max)
		}
	})

	t.Run("max int64 nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int64] = nil

		max, ok := enumerator.MaxInt64(selector.Int64)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for nil enumerator, got %d", max)
		}
	})

	t.Run("max int64 with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{1, 2, 3})

		max, ok := enumerator.MaxInt64(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for nil keySelector, got %d", max)
		}
	})

	t.Run("max int with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int64
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}

		enumerator := FromSlice(people)
		max, ok := enumerator.MaxInt64(func(p Person) int64 { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 35 {
			t.Errorf("Expected max age 35, got %d", max)
		}
	})

	t.Run("max int64 with string length key selector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hi", "world", "go"})

		max, ok := enumerator.MaxInt64(func(s string) int64 { return int64(len(s)) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 5 {
			t.Errorf("Expected max length 5, got %d", max)
		}
	})

	t.Run("max int64 with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{5, 9, 8, 1, 9})

		max, ok := enumerator.MaxInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})

	t.Run("max int64 with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{5, 0, 8, 0, 9})

		max, ok := enumerator.MaxInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})
}

func TestMaxIntStruct(t *testing.T) {
	t.Run("max int from struct field", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price int
		}

		products := []Product{
			{Name: "Laptop", Price: 1000},
			{Name: "Mouse", Price: 50},
			{Name: "Keyboard", Price: 80},
		}

		enumerator := FromSlice(products)
		max, ok := enumerator.MaxInt(func(p Product) int { return p.Price })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 1000 {
			t.Errorf("Expected max price 1000, got %d", max)
		}
	})

	t.Run("max int from struct with negative values", func(t *testing.T) {
		t.Parallel()
		type Temperature struct {
			City  string
			Value int
		}

		temps := []Temperature{
			{City: "Moscow", Value: -5},
			{City: "London", Value: 10},
			{City: "Berlin", Value: -15},
		}

		enumerator := FromSlice(temps)
		max, ok := enumerator.MaxInt(func(t Temperature) int { return t.Value })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 10 {
			t.Errorf("Expected max temperature 10, got %d", max)
		}
	})
}

func TestMaxIntEdgeCases(t *testing.T) {
	t.Run("max int with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 5, 5, 5})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 5 {
			t.Errorf("Expected max 5, got %d", max)
		}
	})

	t.Run("max int with large numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1000000, 999999, 1000001})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 1000001 {
			t.Errorf("Expected max 1000001, got %d", max)
		}
	})

	t.Run("max int with max int value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{100, 2147483647, 50})

		max, ok := enumerator.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 2147483647 {
			t.Errorf("Expected max 2147483647, got %d", max)
		}
	})
}

func TestMaxIntWithOperations(t *testing.T) {
	t.Run("max int after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		max, ok := filtered.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 6 {
			t.Errorf("Expected max 6 (from even numbers), got %d", max)
		}
	})

	t.Run("max int after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})
		taken := enumerator.Take(3)

		max, ok := taken.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 8 {
			t.Errorf("Expected max 8 (from first 3 elements), got %d", max)
		}
	})

	t.Run("max int after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 20, 8, 1, 9, 3})
		skipped := enumerator.Skip(2)

		max, ok := skipped.MaxInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9 (after skipping 2 elements), got %d", max)
		}
	})
}

func TestMaxIntCustomKeySelector(t *testing.T) {
	t.Run("max int by absolute value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, 2, -1, 8, -3})

		max, ok := enumerator.MaxInt(func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		})

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 8 {
			t.Errorf("Expected max absolute value 8, got %d", max)
		}
	})

	t.Run("max int by squared value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-3, 2, -4, 1})

		max, ok := enumerator.MaxInt(func(x int) int { return x * x })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 16 {
			t.Errorf("Expected max squared value 16, got %d", max)
		}
	})
}

func TestMaxIntNonComparable(t *testing.T) {
	t.Run("max int from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name []string
			Age  int
		}

		people := []Person{
			{Name: []string{"Alice"}, Age: 30},
			{Name: []string{"Bob"}, Age: 25},
			{Name: []string{"Charlie"}, Age: 35},
		}

		enumerator := FromSliceAny(people)
		max, ok := enumerator.MaxInt(func(p Person) int { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 35 {
			t.Errorf("Expected max age 35, got %d", max)
		}
	})
}

func TestMaxInt64NonComparable(t *testing.T) {
	t.Run("max int64 from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name []string
			Age  int64
		}

		people := []Person{
			{Name: []string{"Alice"}, Age: 30},
			{Name: []string{"Bob"}, Age: 25},
			{Name: []string{"Charlie"}, Age: 35},
		}

		enumerator := FromSliceAny(people)
		max, ok := enumerator.MaxInt64(func(p Person) int64 { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 35 {
			t.Errorf("Expected max age 35, got %d", max)
		}
	})
}

func BenchmarkMaxInt(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{5, 2, 8, 1, 9}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxInt(selector.Int)
			if !ok || result != 9 {
				b.Fatalf("Expected 9, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxInt(selector.Int)
			if !ok || result != 999 {
				b.Fatalf("Expected 999, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxInt(selector.Int)
			if !ok || result != 9999 {
				b.Fatalf("Expected 9999, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]int{})
			result, ok := enumerator.MaxInt(selector.Int)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %d, ok: %v", result, ok)
			}
		}
	})
}
