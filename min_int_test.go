package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMinInt(t *testing.T) {
	t.Run("min int from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}
	})

	t.Run("min int from non-comparable struct slice", func(t *testing.T) {
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
		min, ok := enumerator.MinInt(func(p Person) int { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min age 25, got %d", min)
		}
	})

	t.Run("min int from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, -2, -8, -1, -9})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -9 {
			t.Errorf("Expected min -9, got %d", min)
		}
	})

	t.Run("min int from mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, 2, -8, 1, 0})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -8 {
			t.Errorf("Expected min -8, got %d", min)
		}
	})

	t.Run("min int single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 42 {
			t.Errorf("Expected min 42, got %d", min)
		}
	})

	t.Run("min int empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		min, ok := enumerator.MinInt(selector.Int)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for empty slice, got %d", min)
		}
	})

	t.Run("min int nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		min, ok := enumerator.MinInt(selector.Int)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil enumerator, got %d", min)
		}
	})

	t.Run("min int with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		min, ok := enumerator.MinInt(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil keySelector, got %d", min)
		}
	})

	t.Run("min int with custom key selector", func(t *testing.T) {
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
		min, ok := enumerator.MinInt(func(p Person) int { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min age 25, got %d", min)
		}
	})

	t.Run("min int with string length key selector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hi", "world", "go"})

		min, ok := enumerator.MinInt(func(s string) int { return len(s) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 2 {
			t.Errorf("Expected min length 2, got %d", min)
		}
	})

	t.Run("min int with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 1, 8, 1, 9})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}
	})

	t.Run("min int with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 0, 8, 0, 9})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %d", min)
		}
	})
}

func TestMinInt64(t *testing.T) {
	t.Run("min int64 from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{5, 2, 8, 1, 9})

		min, ok := enumerator.MinInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}
	})

	t.Run("min int64 from non-comparable struct slice", func(t *testing.T) {
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
		min, ok := enumerator.MinInt64(func(p Person) int64 { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min age 25, got %d", min)
		}
	})

	t.Run("min int64 from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{-5, -2, -8, -1, -9})

		min, ok := enumerator.MinInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -9 {
			t.Errorf("Expected min -9, got %d", min)
		}
	})

	t.Run("min int64 from mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{-5, 2, -8, 1, 0})

		min, ok := enumerator.MinInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -8 {
			t.Errorf("Expected min -8, got %d", min)
		}
	})

	t.Run("min int64 single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{42})

		min, ok := enumerator.MinInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 42 {
			t.Errorf("Expected min 42, got %d", min)
		}
	})

	t.Run("min int64 empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{})

		min, ok := enumerator.MinInt64(selector.Int64)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for empty slice, got %d", min)
		}
	})

	t.Run("min int64 nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int64] = nil

		min, ok := enumerator.MinInt64(selector.Int64)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil enumerator, got %d", min)
		}
	})

	t.Run("min int64 with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{1, 2, 3})

		min, ok := enumerator.MinInt64(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil keySelector, got %d", min)
		}
	})

	t.Run("min int64 with custom key selector", func(t *testing.T) {
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
		min, ok := enumerator.MinInt64(func(p Person) int64 { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min age 25, got %d", min)
		}
	})

	t.Run("min int64 with string length key selector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hi", "world", "go"})

		min, ok := enumerator.MinInt64(func(s string) int64 { return int64(len(s)) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 2 {
			t.Errorf("Expected min length 2, got %d", min)
		}
	})

	t.Run("min int64 with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{5, 1, 8, 1, 9})

		min, ok := enumerator.MinInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}
	})

	t.Run("min int64 with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{5, 0, 8, 0, 9})

		min, ok := enumerator.MinInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %d", min)
		}
	})
}

func TestMinIntStruct(t *testing.T) {
	t.Run("min int from struct field", func(t *testing.T) {
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
		min, ok := enumerator.MinInt(func(p Product) int { return p.Price })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 50 {
			t.Errorf("Expected min price 50, got %d", min)
		}
	})

	t.Run("min int from struct with negative values", func(t *testing.T) {
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
		min, ok := enumerator.MinInt(func(t Temperature) int { return t.Value })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -15 {
			t.Errorf("Expected min temperature -15, got %d", min)
		}
	})
}

func TestMinIntEdgeCases(t *testing.T) {
	t.Run("min int with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 5, 5, 5})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 5 {
			t.Errorf("Expected min 5, got %d", min)
		}
	})

	t.Run("min int with large numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1000000, 999999, 1000001})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 999999 {
			t.Errorf("Expected min 999999, got %d", min)
		}
	})

	t.Run("min int with min int value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{100, -2147483648, 50})

		min, ok := enumerator.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -2147483648 {
			t.Errorf("Expected min -2147483648, got %d", min)
		}
	})
}

func TestMinIntWithOperations(t *testing.T) {
	t.Run("min int after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		min, ok := filtered.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 2 {
			t.Errorf("Expected min 2 (from even numbers), got %d", min)
		}
	})

	t.Run("min int after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})
		taken := enumerator.Take(3)

		min, ok := taken.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 2 {
			t.Errorf("Expected min 2 (from first 3 elements), got %d", min)
		}
	})

	t.Run("min int after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})
		skipped := enumerator.Skip(2)

		min, ok := skipped.MinInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1 (after skipping 2 elements), got %d", min)
		}
	})
}

func TestMinIntCustomKeySelector(t *testing.T) {
	t.Run("min int by absolute value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, 2, -1, 8, -3})

		min, ok := enumerator.MinInt(func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		})

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min absolute value 1, got %d", min)
		}
	})

	t.Run("min int by squared value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-3, 2, -4, 1})

		min, ok := enumerator.MinInt(func(x int) int { return x * x })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min squared value 1, got %d", min)
		}
	})
}

func BenchmarkMinInt(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{5, 2, 8, 1, 9}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinInt(selector.Int)
			if !ok || result != 1 {
				b.Fatalf("Expected 1, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = 1000 - i
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinInt(selector.Int)
			if !ok || result != 1 {
				b.Fatalf("Expected 1, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = 10000 - i
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinInt(selector.Int)
			if !ok || result != 1 {
				b.Fatalf("Expected 1, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]int{})
			result, ok := enumerator.MinInt(selector.Int)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %d, ok: %v", result, ok)
			}
		}
	})
}
