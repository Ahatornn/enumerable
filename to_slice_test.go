package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestToSlice(t *testing.T) {
	t.Run("convert non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []int{1, 2, 3, 4, 5}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("convert non-empty slice for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := [][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if len(result[i]) != len(v) {
				t.Errorf("Expected slice length %d at index %d, got %d", len(v), i, len(result[i]))
				continue
			}
			for j, val := range v {
				if result[i][j] != val {
					t.Errorf("Expected %d at index [%d][%d], got %d", val, i, j, result[i][j])
				}
			}
		}
	})

	t.Run("convert single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		if len(result) != 1 {
			t.Fatalf("Expected length 1, got %d", len(result))
		}

		if result[0] != 42 {
			t.Errorf("Expected 42, got %d", result[0])
		}
	})

	t.Run("convert empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected empty slice, got nil")
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("convert nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.ToSlice()

		if result == nil {
			t.Error("Expected empty slice for nil enumerator, got nil")
		}
		if len(result) != 0 {
			t.Errorf("Expected empty slice for nil enumerator, got slice with length %d", len(result))
		}
	})

	t.Run("convert string slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []string{"hello", "world", "go"}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
			}
		}
	})
}

func TestToSliceStruct(t *testing.T) {
	t.Run("convert struct slice", func(t *testing.T) {
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
		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		if len(result) != len(people) {
			t.Fatalf("Expected length %d, got %d", len(people), len(result))
		}

		for i, v := range people {
			if result[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, result[i])
			}
		}
	})

	t.Run("convert empty struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}

		enumerator := FromSlice(people)
		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected empty slice, got nil")
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})
}

func TestToSliceBoolean(t *testing.T) {
	t.Run("convert boolean slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []bool{true, false, true, false}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %t at index %d, got %t", v, i, result[i])
			}
		}
	})
}

func TestToSliceWithOperations(t *testing.T) {
	t.Run("to slice after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		result := filtered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []int{2, 4, 6}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("to slice after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		result := distinct.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []int{1, 2, 3, 4}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("to slice after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(4)

		result := taken.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []int{1, 2, 3, 4}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})
}

func TestToSliceEdgeCases(t *testing.T) {
	t.Run("to slice with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []int{0, 0, 0}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("to slice with empty strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "", ""})

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []string{"", "", ""}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected '%s' at index %d, got '%s'", v, i, result[i])
			}
		}
	})

	t.Run("to slice with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 3)

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []string{"test", "test", "test"}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
			}
		}
	})

	t.Run("to slice with range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 5) // 1, 2, 3, 4, 5

		result := enumerator.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []int{1, 2, 3, 4, 5}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})
}

func TestOrderEnumeratorToSlice(t *testing.T) {
	t.Run("order enumerator to slice with integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		result := ordered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []int{1, 2, 3, 5, 8, 9}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("order enumerator any to slice with strings", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{"banana", "apple", "cherry", "date"})

		ordered := enumerator.OrderBy(comparer.ComparerString)
		result := ordered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []string{"apple", "banana", "cherry", "date"}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
			}
		}
	})

	t.Run("order enumerator to slice with empty source", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		result := ordered.ToSlice()

		if result == nil {
			t.Error("Expected empty slice, got nil")
		}
		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("order enumerator any to slice with empty source", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{})

		ordered := enumerator.OrderBy(comparer.ComparerString)
		result := ordered.ToSlice()

		if result == nil {
			t.Error("Expected empty slice, got nil")
		}
		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("order enumerator to slice with single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		result := ordered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("Expected [42], got %v", result)
		}
	})

	t.Run("order enumerator any to slice with single element", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{"single"})

		ordered := enumerator.OrderBy(comparer.ComparerString)
		result := ordered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}
		if len(result) != 1 || result[0] != "single" {
			t.Errorf("Expected [\"single\"], got %v", result)
		}
	})

	t.Run("order enumerator to slice with struct and multiple sorting levels", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Charlie", Age: 30},
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Diana", Age: 25},
		}
		enumerator := FromSlice(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age }).
			ThenBy(func(a, b Person) int { return compareStrings(a.Name, b.Name) })
		result := ordered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Diana", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 30},
		}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, result[i])
			}
		}
	})

	t.Run("order enumerator any to slice with complex struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name     string
			Priority int
			Options  []string
		}

		configs := []Config{
			{Name: "High", Priority: 1, Options: []string{"opt1", "opt2"}},
			{Name: "Low", Priority: 3, Options: []string{"opt3"}},
			{Name: "Medium", Priority: 2, Options: []string{"opt4", "opt5"}},
		}
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })
		result := ordered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		if len(result) != 3 {
			t.Fatalf("Expected length 3, got %d", len(result))
		}

		expectedPriorities := []int{1, 2, 3}
		for i, expected := range expectedPriorities {
			if result[i].Priority != expected {
				t.Errorf("Expected priority %d at index %d, got %d", expected, i, result[i].Priority)
			}
		}
	})

	t.Run("order enumerator to slice preserves stability", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			Index int
		}

		items := []Item{
			{Value: 2, Index: 1},
			{Value: 1, Index: 2},
			{Value: 2, Index: 3},
			{Value: 1, Index: 4},
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })
		result := ordered.ToSlice()

		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		if result[0].Index != 2 || result[1].Index != 4 {
			t.Errorf("Stability not preserved for Value=1: got indices %d, %d", result[0].Index, result[1].Index)
		}

		if result[2].Index != 1 || result[3].Index != 3 {
			t.Errorf("Stability not preserved for Value=2: got indices %d, %d", result[2].Index, result[3].Index)
		}
	})

	t.Run("multiple to slice calls return same sorted result", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 4, 1, 5, 9, 2, 6})

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		result1 := ordered.ToSlice()
		result2 := ordered.ToSlice()

		if len(result1) != len(result2) {
			t.Fatalf("Expected same lengths, got %d and %d", len(result1), len(result2))
		}

		for i := range result1 {
			if result1[i] != result2[i] {
				t.Errorf("Expected same elements at index %d, got %d and %d", i, result1[i], result2[i])
			}
		}

		expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
		for i, v := range expected {
			if result1[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result1[i])
			}
		}
	})

	t.Run("order enumerator to slice with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)
		result := ordered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		expected := []int{9, 8, 5, 3, 2, 1}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("order enumerator any to slice with descending order", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price float64
		}

		products := []Product{
			{Name: "Laptop", Price: 1200.0},
			{Name: "Phone", Price: 800.0},
			{Name: "Tablet", Price: 500.0},
		}
		var enumerator = FromSliceAny(products)

		ordered := enumerator.OrderByDescending(func(a, b Product) int {
			if a.Price < b.Price {
				return -1
			}
			if a.Price > b.Price {
				return 1
			}
			return 0
		})
		result := ordered.ToSlice()

		if result == nil {
			t.Fatal("Expected slice, got nil")
		}

		if len(result) != 3 {
			t.Fatalf("Expected length 3, got %d", len(result))
		}

		expectedPrices := []float64{1200.0, 800.0, 500.0}
		for i, expected := range expectedPrices {
			if result[i].Price != expected {
				t.Errorf("Expected price %f at index %d, got %f", expected, i, result[i].Price)
			}
		}
	})
}

func BenchmarkOrderEnumeratorToSlice(b *testing.B) {
	b.Run("order enumerator to slice small", func(b *testing.B) {
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3, 7, 4, 6})

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b int) int { return a - b })
			_ = ordered.ToSlice()
		}
	})

	b.Run("order enumerator any to slice small", func(b *testing.B) {
		var enumerator = FromSliceAny([]string{"banana", "apple", "cherry", "date", "elderberry"})

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(comparer.ComparerString)
			_ = ordered.ToSlice()
		}
	})

	b.Run("order enumerator to slice medium", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = 1000 - i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b int) int { return a - b })
			_ = ordered.ToSlice()
		}
	})

	b.Run("order enumerator to slice with multiple levels", func(b *testing.B) {
		type Person struct {
			Name string
			Age  int
		}

		people := make([]Person, 100)
		for i := 0; i < 100; i++ {
			people[i] = Person{
				Name: fmt.Sprintf("Person%d", i),
				Age:  i % 10,
			}
		}
		enumerator := FromSlice(people)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age }).
				ThenBy(func(a, b Person) int { return compareStrings(a.Name, b.Name) })
			_ = ordered.ToSlice()
		}
	})
}

func BenchmarkToSlice(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToSlice()
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToSlice()
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToSlice()
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToSlice()
		}
	})
}
