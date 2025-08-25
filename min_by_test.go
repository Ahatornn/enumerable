package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestMinBy(t *testing.T) {
	t.Run("min by with int slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}
	})

	t.Run("min by with non-comparable struct slice", func(t *testing.T) {
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
		min, ok := enumerator.MinBy(func(a, b Person) int {
			return comparer.ComparerInt(a.Age, b.Age)
		})

		if !ok {
			t.Error("Expected ok to be true")
		}

		if min.Age != 25 {
			t.Errorf("Expected min age 25, got %d", min.Age)
		}

		if len(min.Name) != 1 || min.Name[0] != "Bob" {
			t.Errorf("Expected min person name 'Bob', got %+v", min.Name)
		}
	})

	t.Run("min by with negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, -2, -8, -1, -9})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -9 {
			t.Errorf("Expected min -9, got %d", min)
		}
	})

	t.Run("min by with mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, 2, -8, 1, 0})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -8 {
			t.Errorf("Expected min -8, got %d", min)
		}
	})

	t.Run("min by single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 42 {
			t.Errorf("Expected min 42, got %d", min)
		}
	})

	t.Run("min by empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		var expected int
		if min != expected {
			t.Errorf("Expected min %v for empty slice, got %v", expected, min)
		}
	})

	t.Run("min by nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		var expected int
		if min != expected {
			t.Errorf("Expected min %v for nil enumerator, got %v", expected, min)
		}
	})

	t.Run("min by with nil comparer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		min, ok := enumerator.MinBy(nil)

		if ok {
			t.Error("Expected ok to be false for nil comparer")
		}
		var expected int
		if min != expected {
			t.Errorf("Expected min %v for nil comparer, got %v", expected, min)
		}
	})

	t.Run("min by with string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "banana"})

		min, ok := enumerator.MinBy(comparer.ComparerString)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "apple" {
			t.Errorf("Expected min 'apple', got '%s'", min)
		}
	})

	t.Run("min by with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 1, 8, 1, 9})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}
	})

	t.Run("min by with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 0, 8, 0, 9})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %d", min)
		}
	})
}

func TestMinByStruct(t *testing.T) {
	t.Run("min by struct with custom comparer", func(t *testing.T) {
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
		min, ok := enumerator.MinBy(func(a, b Person) int {
			return comparer.ComparerInt(a.Age, b.Age)
		})

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := Person{Name: "Bob", Age: 25}
		if min != expected {
			t.Errorf("Expected min person %+v, got %+v", expected, min)
		}
	})

	t.Run("min by struct by name", func(t *testing.T) {
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
		min, ok := enumerator.MinBy(func(a, b Product) int {
			return comparer.ComparerString(a.Name, b.Name)
		})

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := Product{Name: "Keyboard", Price: 80}
		if min != expected {
			t.Errorf("Expected min product %+v, got %+v", expected, min)
		}
	})
}

func TestMinByFloat(t *testing.T) {
	t.Run("min by with float64 slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{5.5, 2.2, 8.8, 1.1, 9.9})

		min, ok := enumerator.MinBy(comparer.ComparerFloat64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1, got %f", min)
		}
	})

	t.Run("min by with negative floats", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{-5.5, -2.2, -8.8, -1.1, -9.9})

		min, ok := enumerator.MinBy(comparer.ComparerFloat64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -9.9 {
			t.Errorf("Expected min -9.9, got %f", min)
		}
	})
}

func TestMinByEdgeCases(t *testing.T) {
	t.Run("min by with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 5, 5, 5})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 5 {
			t.Errorf("Expected min 5, got %d", min)
		}
	})

	t.Run("min by with large numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1000000, 999999, 1000001})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 999999 {
			t.Errorf("Expected min 999999, got %d", min)
		}
	})

	t.Run("min by with min int value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{100, -2147483648, 50})

		min, ok := enumerator.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -2147483648 {
			t.Errorf("Expected min -2147483648, got %d", min)
		}
	})
}

func TestMinByWithOperations(t *testing.T) {
	t.Run("min by after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		min, ok := filtered.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 2 {
			t.Errorf("Expected min 2 (from even numbers), got %d", min)
		}
	})

	t.Run("min by after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})
		taken := enumerator.Take(3)

		min, ok := taken.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 2 {
			t.Errorf("Expected min 2 (from first 3 elements), got %d", min)
		}
	})

	t.Run("min by after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, 2, 8, 1, 9, 3})
		skipped := enumerator.Skip(2)

		min, ok := skipped.MinBy(comparer.ComparerInt)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1 (after skipping 2 elements), got %d", min)
		}
	})
}

func TestMinByCustomComparer(t *testing.T) {
	t.Run("min by with reverse comparer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9})

		// Reverse comparer - finds maximum instead of minimum
		reverseComparer := func(a, b int) int {
			return comparer.ComparerInt(b, a)
		}

		max, ok := enumerator.MinBy(reverseComparer)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 9 {
			t.Errorf("Expected max 9, got %d", max)
		}
	})

	t.Run("min by with custom modulus comparer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{15, 7, 23, 4, 11})

		// Compare by modulus 10
		modulusComparer := func(a, b int) int {
			return comparer.ComparerInt(a%10, b%10)
		}

		min, ok := enumerator.MinBy(modulusComparer)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 11 { // 11 % 10 = 1, which is minimum
			t.Errorf("Expected min 11 (modulus 1), got %d", min)
		}
	})
}

func BenchmarkMinBy(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{5, 2, 8, 1, 9}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinBy(comparer.ComparerInt)
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
			result, ok := enumerator.MinBy(comparer.ComparerInt)
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
			result, ok := enumerator.MinBy(comparer.ComparerInt)
			if !ok || result != 1 {
				b.Fatalf("Expected 1, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]int{})
			result, ok := enumerator.MinBy(comparer.ComparerInt)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %v, ok: %v", result, ok)
			}
		}
	})
}
