package enumerable

import (
	"testing"
)

func TestSumInt(t *testing.T) {
	t.Run("sum simple integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 15 {
			t.Errorf("Expected sum 15, got %d", sum)
		}
	})

	t.Run("sum simple integers for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1},
			{2},
			{3},
			{4},
			{5},
		})

		sum := enumerator.SumInt(func(slice []int) int {
			if len(slice) > 0 {
				return slice[0]
			}
			return 0
		})

		if sum != 15 {
			t.Errorf("Expected sum 15, got %d", sum)
		}
	})

	t.Run("sum with transformation", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4})

		sum := enumerator.SumInt(func(n int) int { return n * n }) // Sum of squares

		if sum != 30 { // 1 + 4 + 9 + 16 = 30
			t.Errorf("Expected sum 30, got %d", sum)
		}
	})

	t.Run("sum single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 42 {
			t.Errorf("Expected sum 42, got %d", sum)
		}
	})

	t.Run("sum empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 0 {
			t.Errorf("Expected sum 0 for empty slice, got %d", sum)
		}
	})

	t.Run("sum nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 0 {
			t.Errorf("Expected sum 0 for nil enumerator, got %d", sum)
		}
	})

	t.Run("sum with negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-1, -2, 3, 4})

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 4 { // -1 + -2 + 3 + 4 = 4
			t.Errorf("Expected sum 4, got %d", sum)
		}
	})

	t.Run("sum with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 5, 0})

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 5 {
			t.Errorf("Expected sum 5, got %d", sum)
		}
	})
}

func TestSumIntStruct(t *testing.T) {
	t.Run("sum struct field", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price int
		}

		products := []Product{
			{Name: "Apple", Price: 10},
			{Name: "Banana", Price: 5},
			{Name: "Orange", Price: 8},
		}

		enumerator := FromSlice(products)
		sum := enumerator.SumInt(func(p Product) int { return p.Price })

		if sum != 23 { // 10 + 5 + 8 = 23
			t.Errorf("Expected sum 23, got %d", sum)
		}
	})

	t.Run("sum struct field with zero values", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price int
		}

		products := []Product{
			{Name: "Free", Price: 0},
			{Name: "Cheap", Price: 1},
			{Name: "Free2", Price: 0},
		}

		enumerator := FromSlice(products)
		sum := enumerator.SumInt(func(p Product) int { return p.Price })

		if sum != 1 {
			t.Errorf("Expected sum 1, got %d", sum)
		}
	})
}

func TestSumIntString(t *testing.T) {
	t.Run("sum string lengths", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "bb", "ccc"})

		sum := enumerator.SumInt(func(s string) int { return len(s) })

		if sum != 6 { // 1 + 2 + 3 = 6
			t.Errorf("Expected sum 6, got %d", sum)
		}
	})

	t.Run("sum parsed integers from strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"1", "2", "3", "4"})

		sum := enumerator.SumInt(func(s string) int {
			switch s {
			case "1":
				return 1
			case "2":
				return 2
			case "3":
				return 3
			case "4":
				return 4
			default:
				return 0
			}
		})

		if sum != 10 {
			t.Errorf("Expected sum 10, got %d", sum)
		}
	})
}

func TestSumIntWithOperations(t *testing.T) {
	t.Run("sum after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		sum := filtered.SumInt(func(n int) int { return n })

		if sum != 12 { // 2 + 4 + 6 = 12
			t.Errorf("Expected sum 12, got %d", sum)
		}
	})

	t.Run("sum after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(4)

		sum := taken.SumInt(func(n int) int { return n })

		if sum != 10 { // 1 + 2 + 3 + 4 = 10
			t.Errorf("Expected sum 10, got %d", sum)
		}
	})

	t.Run("sum after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		sum := distinct.SumInt(func(n int) int { return n })

		if sum != 10 { // 1 + 2 + 3 + 4 = 10
			t.Errorf("Expected sum 10, got %d", sum)
		}
	})
}

func TestSumIntEdgeCases(t *testing.T) {
	t.Run("sum with large numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1000000, 2000000, 3000000})

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 6000000 {
			t.Errorf("Expected sum 6000000, got %d", sum)
		}
	})

	t.Run("sum with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(5, 4)

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 20 { // 5 + 5 + 5 + 5 = 20
			t.Errorf("Expected sum 20, got %d", sum)
		}
	})

	t.Run("sum with range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 5) // 1, 2, 3, 4, 5

		sum := enumerator.SumInt(func(n int) int { return n })

		if sum != 15 {
			t.Errorf("Expected sum 15, got %d", sum)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkSumInt(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.SumInt(func(n int) int { return n })
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.SumInt(func(n int) int { return n })
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.SumInt(func(n int) int { return n })
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.SumInt(func(n int) int { return n })
		}
	})
}
