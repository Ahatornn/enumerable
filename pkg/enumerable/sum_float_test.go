package enumerable

import (
	"math"
	"testing"
)

func TestSumFloat(t *testing.T) {
	t.Run("sum simple float32 values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.1, 2.2, 3.3, 4.4, 5.5})

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		expected := float32(16.5)
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum simple float32 values for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]float32{
			{1.1},
			{2.2},
			{3.3},
			{4.4},
			{5.5},
		})

		sum := enumerator.SumFloat(func(slice []float32) float32 {
			if len(slice) > 0 {
				return slice[0]
			}
			return 0
		})

		expected := float32(16.5)
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum with transformation", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.0, 2.0, 3.0, 4.0})

		sum := enumerator.SumFloat(func(n float32) float32 { return n * n }) // Sum of squares

		expected := float32(30.0) // 1 + 4 + 9 + 16 = 30
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{42.5})

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		expected := float32(42.5)
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{})

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		if sum != 0 {
			t.Errorf("Expected sum 0 for empty slice, got %f", sum)
		}
	})

	t.Run("sum nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[float32] = nil

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		if sum != 0 {
			t.Errorf("Expected sum 0 for nil enumerator, got %f", sum)
		}
	})

	t.Run("sum with negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{-1.5, -2.5, 3.0, 4.0})

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		expected := float32(3.0) // -1.5 + -2.5 + 3.0 + 4.0 = 3.0
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{0.0, 0.0, 5.5, 0.0})

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		expected := float32(5.5)
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})
}

func TestSumFloatStruct(t *testing.T) {
	t.Run("sum struct field", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price float32
		}

		products := []Product{
			{Name: "Apple", Price: 10.5},
			{Name: "Banana", Price: 5.25},
			{Name: "Orange", Price: 8.75},
		}

		enumerator := FromSlice(products)
		sum := enumerator.SumFloat(func(p Product) float32 { return p.Price })

		expected := float32(24.5) // 10.5 + 5.25 + 8.75 = 24.5
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum struct field with zero values", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price float32
		}

		products := []Product{
			{Name: "Free", Price: 0.0},
			{Name: "Cheap", Price: 1.99},
			{Name: "Free2", Price: 0.0},
		}

		enumerator := FromSlice(products)
		sum := enumerator.SumFloat(func(p Product) float32 { return p.Price })

		expected := float32(1.99)
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})
}

func TestSumFloatString(t *testing.T) {
	t.Run("sum parsed floats from strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"1.5", "2.5", "3.0", "4.0"})

		sum := enumerator.SumFloat(func(s string) float32 {
			switch s {
			case "1.5":
				return 1.5
			case "2.5":
				return 2.5
			case "3.0":
				return 3.0
			case "4.0":
				return 4.0
			default:
				return 0
			}
		})

		expected := float32(11.0)
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})
}

func TestSumFloatWithOperations(t *testing.T) {
	t.Run("sum after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6})
		filtered := enumerator.Where(func(n float32) bool { return n > 3.0 })

		sum := filtered.SumFloat(func(n float32) float32 { return n })

		expected := float32(19.8) // 4.4 + 5.5 + 6.6 = 16.5
		if math.Abs(float64(sum-expected)) > 0.001 {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8})
		taken := enumerator.Take(4)

		sum := taken.SumFloat(func(n float32) float32 { return n })

		expected := float32(11.0) // 1.1 + 2.2 + 3.3 + 4.4 = 11.0
		if math.Abs(float64(sum-expected)) > 0.001 {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})
}

func TestSumFloatEdgeCases(t *testing.T) {
	t.Run("sum with small numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{0.1, 0.2, 0.3})

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		expected := float32(0.6)
		if math.Abs(float64(sum-expected)) > 0.001 {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat(float32(2.5), 4)

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		expected := float32(10.0) // 2.5 * 4 = 10.0
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("sum with very small differences", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.000001, 1.000002, 1.000003})

		sum := enumerator.SumFloat(func(n float32) float32 { return n })

		expected := float32(3.000006)
		if math.Abs(float64(sum-expected)) > 0.00001 {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkSumFloat(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.SumFloat(func(n float32) float32 { return n })
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]float32, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = float32(i) * 1.5
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.SumFloat(func(n float32) float32 { return n })
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]float32, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = float32(i) * 0.1
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.SumFloat(func(n float32) float32 { return n })
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]float32{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.SumFloat(func(n float32) float32 { return n })
		}
	})
}
