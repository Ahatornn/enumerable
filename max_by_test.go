package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestMaxBy(t *testing.T) {
	t.Run("max by with int slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})

	t.Run("max by with negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, -2, -8, -1, -9})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != -1 {
			t.Errorf("Expected max -1, got %d", max)
		}
	})

	t.Run("max by with mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, 2, -8, 1, 0})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 2 {
			t.Errorf("Expected max 2, got %d", max)
		}
	})

	t.Run("max by single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 42 {
			t.Errorf("Expected max 42, got %d", max)
		}
	})

	t.Run("max by empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		var expected int
		if max != expected {
			t.Errorf("Expected max %v for empty slice, got %v", expected, max)
		}
	})

	t.Run("max by nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		var expected int
		if max != expected {
			t.Errorf("Expected max %v for nil enumerator, got %v", expected, max)
		}
	})

	t.Run("max by with nil comparer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		max, ok := enumerator.MaxBy(nil)

		if ok {
			t.Error("Expected ok to be false for nil comparer")
		}
		var expected int
		if max != expected {
			t.Errorf("Expected max %v for nil comparer, got %v", expected, max)
		}
	})

	t.Run("max by with string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "banana"})

		max, ok := enumerator.MaxBy(comparer.ComparerString)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zebra" {
			t.Errorf("Expected max 'zebra', got '%s'", max)
		}
	})

	t.Run("max by with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 9, 8, 1, 9})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})

	t.Run("max by with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 0, 8, 0, 9})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})
}

func TestMaxByStruct(t *testing.T) {
	t.Run("max by struct with custom comparer", func(t *testing.T) {
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
		max, ok := enumerator.MaxBy(func(a, b Person) int {
			return comparer.ComparerInt(a.Age, b.Age)
		})

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := Person{Name: "Charlie", Age: 35}
		if max != expected {
			t.Errorf("Expected max person %+v, got %+v", expected, max)
		}
	})

	t.Run("max by struct by name", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price int
		}

		products := []Product{
			{Name: "Laptop", Price: 1000},
			{Name: "Mouse", Price: 50},
			{Name: "Keyboard", Price: 80},
			{Name: "Zebra", Price: 100},
		}

		enumerator := FromSlice(products)
		max, ok := enumerator.MaxBy(func(a, b Product) int {
			return comparer.ComparerString(a.Name, b.Name)
		})

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := Product{Name: "Zebra", Price: 100}
		if max != expected {
			t.Errorf("Expected max product %+v, got %+v", expected, max)
		}
	})
}

func TestMaxByFloat(t *testing.T) {
	t.Run("max by with float64 slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{5.5, 2.2, 8.8, 1.1, 9.9})

		max, ok := enumerator.MaxBy(comparer.ComparerFloat64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9.9 {
			t.Errorf("Expected max 9.9, got %f", max)
		}
	})

	t.Run("max by with negative floats", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{-5.5, -2.2, -8.8, -1.1, -9.9})

		max, ok := enumerator.MaxBy(comparer.ComparerFloat64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != -1.1 {
			t.Errorf("Expected max -1.1, got %f", max)
		}
	})
}

func TestMaxByEdgeCases(t *testing.T) {
	t.Run("max by with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 5, 5, 5})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 5 {
			t.Errorf("Expected max 5, got %d", max)
		}
	})

	t.Run("max by with large numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1000000, 999999, 1000001})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 1000001 {
			t.Errorf("Expected max 1000001, got %d", max)
		}
	})

	t.Run("max by with max int value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{100, 2147483647, 50})

		max, ok := enumerator.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 2147483647 {
			t.Errorf("Expected max 2147483647, got %d", max)
		}
	})
}

func TestMaxByWithOperations(t *testing.T) {
	t.Run("max by after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		max, ok := filtered.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 6 {
			t.Errorf("Expected max 6 (from even numbers), got %d", max)
		}
	})

	t.Run("max by after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})
		taken := enumerator.Take(3)

		max, ok := taken.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 8 {
			t.Errorf("Expected max 8 (from first 3 elements), got %d", max)
		}
	})

	t.Run("max by after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 20, 8, 1, 9, 3})
		skipped := enumerator.Skip(2)

		max, ok := skipped.MaxBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9 (after skipping 2 elements), got %d", max)
		}
	})
}

func TestMaxByCustomComparer(t *testing.T) {
	t.Run("max by with reverse comparer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9})

		reverseComparer := func(a, b int) int {
			return comparer.ComparerInt(b, a)
		}

		min, ok := enumerator.MaxBy(reverseComparer)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}
	})

	t.Run("max by with custom modulus comparer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{15, 7, 23, 4, 11})

		modulusComparer := func(a, b int) int {
			return comparer.ComparerInt(a%10, b%10)
		}

		max, ok := enumerator.MaxBy(modulusComparer)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 7 { // 7 % 10 = 7, which is maximum among 5,7,3,4,1
			t.Errorf("Expected max 7 (modulus 7), got %d", max)
		}
	})
}

func TestMaxByNonComparable(t *testing.T) {
	t.Run("max by with non-comparable struct slice", func(t *testing.T) {
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
		max, ok := enumerator.MaxBy(func(a, b Person) int {
			return comparer.ComparerInt(a.Age, b.Age)
		})

		if !ok {
			t.Error("Expected ok to be true")
		}

		if max.Age != 35 {
			t.Errorf("Expected max age 35, got %d", max.Age)
		}

		if len(max.Name) != 1 || max.Name[0] != "Charlie" {
			t.Errorf("Expected max person name 'Charlie', got %+v", max.Name)
		}
	})

	t.Run("max by with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Score    int
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Score: 100},
			{Settings: map[string]interface{}{"b": 2}, Score: 200},
			{Settings: map[string]interface{}{"c": 3}, Score: 150},
		}

		enumerator := FromSliceAny(configs)
		max, ok := enumerator.MaxBy(func(a, b Config) int {
			return comparer.ComparerInt(a.Score, b.Score)
		})

		if !ok {
			t.Error("Expected ok to be true")
		}

		if max.Score != 200 {
			t.Errorf("Expected max score 200, got %d", max.Score)
		}
	})

	t.Run("max by with function field struct", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Callback func()
			Priority int
		}

		handlers := []Handler{
			{Callback: func() {}, Priority: 1},
			{Callback: func() {}, Priority: 3},
			{Callback: func() {}, Priority: 2},
		}

		enumerator := FromSliceAny(handlers)
		max, ok := enumerator.MaxBy(func(a, b Handler) int {
			return comparer.ComparerInt(a.Priority, b.Priority)
		})

		if !ok {
			t.Error("Expected ok to be true")
		}

		if max.Priority != 3 {
			t.Errorf("Expected max priority 3, got %d", max.Priority)
		}
	})
}

func BenchmarkMaxBy(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{5, 2, 8, 1, 9}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxBy(comparer.ComparerInt)
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
			result, ok := enumerator.MaxBy(comparer.ComparerInt)
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
			result, ok := enumerator.MaxBy(comparer.ComparerInt)
			if !ok || result != 9999 {
				b.Fatalf("Expected 9999, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]int{})
			result, ok := enumerator.MaxBy(comparer.ComparerInt)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %v, ok: %v", result, ok)
			}
		}
	})
}
