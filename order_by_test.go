package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestOrderBy(t *testing.T) {
	t.Run("order by ascending integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.ToSlice()
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

	t.Run("order by descending integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		result := ordered.ToSlice()
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

	t.Run("order by strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"banana", "apple", "cherry", "date"})

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.ToSlice()
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

	t.Run("order by descending strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"banana", "apple", "cherry", "date"})

		ordered := enumerator.OrderByDescending(comparer.ComparerString)

		result := ordered.ToSlice()
		expected := []string{"date", "cherry", "banana", "apple"}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
			}
		}
	})

	t.Run("order by with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.ToSlice()
		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("order by with single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.ToSlice()
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("Expected [42], got %v", result)
		}
	})

	t.Run("order by with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.ToSlice()
		if len(result) != 0 {
			t.Errorf("Expected empty slice for nil enumerator, got length %d", len(result))
		}
	})

	t.Run("order by struct fields", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Charlie", Age: 35},
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		}
		enumerator := FromSlice(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })

		result := ordered.ToSlice()
		expected := []Person{
			{Name: "Bob", Age: 25},
			{Name: "Alice", Age: 30},
			{Name: "Charlie", Age: 35},
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

	t.Run("order by descending struct fields", func(t *testing.T) {
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
		enumerator := FromSlice(products)

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
		expected := []Product{
			{Name: "Laptop", Price: 1200.0},
			{Name: "Phone", Price: 800.0},
			{Name: "Tablet", Price: 500.0},
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

	t.Run("order by with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.ToSlice()
		expected := []int{1, 1, 2, 2, 3, 3}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("order by preserves stability", func(t *testing.T) {
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

		if result[0].Value != 1 || result[1].Value != 1 || result[2].Value != 2 || result[3].Value != 2 {
			t.Errorf("Values not in correct order: %v", result)
		}

		if result[0].Index != 2 || result[1].Index != 4 {
			t.Errorf("Stability not preserved for Value=1: got indices %d, %d", result[0].Index, result[1].Index)
		}

		if result[2].Index != 1 || result[3].Index != 3 {
			t.Errorf("Stability not preserved for Value=2: got indices %d, %d", result[2].Index, result[3].Index)
		}
	})

	t.Run("order by with ThenBy chaining", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Bob", Age: 30},
			{Name: "Alice", Age: 30},
			{Name: "Charlie", Age: 25},
		}
		enumerator := FromSlice(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age }).
			ThenBy(func(a, b Person) int { return compareStrings(a.Name, b.Name) })

		result := ordered.ToSlice()
		expected := []Person{
			{Name: "Charlie", Age: 25},
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 30},
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
}

func TestOrderByAny(t *testing.T) {
	t.Run("order by ascending integers with any enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.ToSlice()
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

	t.Run("order by descending integers with any enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		result := ordered.ToSlice()
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

	t.Run("order by strings with any enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{"banana", "apple", "cherry", "date"})

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.ToSlice()
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

	t.Run("order by complex struct with slices and maps", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name     string
			Options  []string
			Meta     map[string]interface{}
			Priority int
		}

		configs := []Config{
			{Name: "High", Options: []string{"opt1", "opt2"}, Meta: map[string]interface{}{"key": "value1"}, Priority: 1},
			{Name: "Low", Options: []string{"opt3"}, Meta: map[string]interface{}{"key": "value2"}, Priority: 3},
			{Name: "Medium", Options: []string{"opt4", "opt5", "opt6"}, Meta: map[string]interface{}{"key": "value3"}, Priority: 2},
		}
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })

		result := ordered.ToSlice()
		if len(result) != 3 {
			t.Fatalf("Expected length 3, got %d", len(result))
		}

		if result[0].Priority != 1 || result[1].Priority != 2 || result[2].Priority != 3 {
			t.Errorf("Expected priority order 1,2,3, got %d,%d,%d",
				result[0].Priority, result[1].Priority, result[2].Priority)
		}

		if result[0].Name != "High" || result[1].Name != "Medium" || result[2].Name != "Low" {
			t.Errorf("Expected names High,Medium,Low, got %s,%s,%s",
				result[0].Name, result[1].Name, result[2].Name)
		}
	})

	t.Run("order by struct with custom comparison logic", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name  string
			Tags  []string
			Score float64
		}

		users := []User{
			{Name: "Alice", Tags: []string{"dev", "go", "python"}, Score: 85.5},
			{Name: "Bob", Tags: []string{"js", "react"}, Score: 92.0},
			{Name: "Charlie", Tags: []string{"java", "spring", "microservices", "docker"}, Score: 78.0},
		}
		var enumerator = FromSliceAny(users)

		// Сортировка по количеству тегов
		ordered := enumerator.OrderBy(func(a, b User) int {
			lenA, lenB := len(a.Tags), len(b.Tags)
			if lenA < lenB {
				return -1
			}
			if lenA > lenB {
				return 1
			}
			return 0
		})

		result := ordered.ToSlice()
		if len(result) != 3 {
			t.Fatalf("Expected length 3, got %d", len(result))
		}

		if len(result[0].Tags) != 2 || len(result[1].Tags) != 3 || len(result[2].Tags) != 4 {
			t.Errorf("Expected tag counts 2,3,4, got %d,%d,%d",
				len(result[0].Tags), len(result[1].Tags), len(result[2].Tags))
		}
	})

	t.Run("order by descending with complex struct", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name       string
			Categories []string
			Price      float64
			Rating     float64
		}

		products := []Product{
			{Name: "Laptop", Categories: []string{"electronics", "computers"}, Price: 1200.0, Rating: 4.5},
			{Name: "Phone", Categories: []string{"electronics", "mobile"}, Price: 800.0, Rating: 4.8},
			{Name: "Tablet", Categories: []string{"electronics", "mobile"}, Price: 500.0, Rating: 4.2},
		}
		var enumerator = FromSliceAny(products)

		ordered := enumerator.OrderByDescending(func(a, b Product) int {
			if a.Rating < b.Rating {
				return -1
			}
			if a.Rating > b.Rating {
				return 1
			}
			return 0
		})

		result := ordered.ToSlice()
		if len(result) != 3 {
			t.Fatalf("Expected length 3, got %d", len(result))
		}

		expectedRatings := []float64{4.8, 4.5, 4.2}
		for i, expected := range expectedRatings {
			if result[i].Rating != expected {
				t.Errorf("Expected rating %f at index %d, got %f", expected, i, result[i].Rating)
			}
		}
	})

	t.Run("order by with empty slice any enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.ToSlice()
		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("order by with nil any enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = nil

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.ToSlice()
		if len(result) != 0 {
			t.Errorf("Expected empty slice for nil enumerator, got length %d", len(result))
		}
	})

	t.Run("order by preserves stability for any enumerator", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			Index int
			Data  []string
		}

		items := []Item{
			{Value: 2, Index: 1, Data: []string{"a"}},
			{Value: 1, Index: 2, Data: []string{"b"}},
			{Value: 2, Index: 3, Data: []string{"c"}},
			{Value: 1, Index: 4, Data: []string{"d"}},
		}
		var enumerator = FromSliceAny(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })

		result := ordered.ToSlice()
		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		// Проверяем порядок Value
		if result[0].Value != 1 || result[1].Value != 1 || result[2].Value != 2 || result[3].Value != 2 {
			t.Errorf("Values not in correct order: %v", result)
		}

		// Проверяем стабильность
		if result[0].Index != 2 || result[1].Index != 4 {
			t.Errorf("Stability not preserved for Value=1: got indices %d, %d", result[0].Index, result[1].Index)
		}

		if result[2].Index != 1 || result[3].Index != 3 {
			t.Errorf("Stability not preserved for Value=2: got indices %d, %d", result[2].Index, result[3].Index)
		}
	})

	t.Run("order by with ThenBy chaining for any enumerator", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Category string
			Name     string
			Tags     []string
			Score    float64
		}

		records := []Record{
			{Category: "A", Name: "Second", Tags: []string{"tag1"}, Score: 85.0},
			{Category: "A", Name: "First", Tags: []string{"tag2"}, Score: 90.0},
			{Category: "B", Name: "Third", Tags: []string{"tag3"}, Score: 75.0},
		}
		var enumerator = FromSliceAny(records)

		ordered := enumerator.OrderBy(func(a, b Record) int { return compareStrings(a.Category, b.Category) }).
			ThenBy(func(a, b Record) int { return compareStrings(a.Name, b.Name) })

		result := ordered.ToSlice()
		expectedCategories := []string{"A", "A", "B"}
		expectedNames := []string{"First", "Second", "Third"}

		if len(result) != 3 {
			t.Fatalf("Expected length 3, got %d", len(result))
		}

		for i, expectedCategory := range expectedCategories {
			if result[i].Category != expectedCategory {
				t.Errorf("Expected category %s at index %d, got %s", expectedCategory, i, result[i].Category)
			}
		}

		for i, expectedName := range expectedNames {
			if result[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, result[i].Name)
			}
		}
	})
}

func compareStrings(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func BenchmarkOrderBy(b *testing.B) {
	b.Run("order by integers small", func(b *testing.B) {
		items := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b int) int { return a - b })
			_ = ordered.ToSlice()
		}
	})

	b.Run("order by integers medium", func(b *testing.B) {
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

	b.Run("order by strings", func(b *testing.B) {
		items := []string{"banana", "apple", "cherry", "date", "elderberry", "fig", "grape"}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b string) int { return compareStrings(a, b) })
			_ = ordered.ToSlice()
		}
	})

	b.Run("order by descending", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderByDescending(func(a, b int) int { return a - b })
			_ = ordered.ToSlice()
		}
	})

	b.Run("order by with nil enumerator", func(b *testing.B) {
		var enumerator Enumerator[int] = nil

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b int) int { return a - b })
			_ = ordered.ToSlice()
		}
	})
}

func BenchmarkOrderByAny(b *testing.B) {
	b.Run("order by integers small any", func(b *testing.B) {
		var enumerator = FromSliceAny([]int{5, 2, 8, 1, 9, 3, 7, 4, 6})

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(comparer.ComparerInt)
			_ = ordered.ToSlice()
		}
	})

	b.Run("order by strings any", func(b *testing.B) {
		var enumerator = FromSliceAny([]string{"banana", "apple", "cherry", "date", "elderberry", "fig", "grape"})

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(comparer.ComparerString)
			_ = ordered.ToSlice()
		}
	})

	b.Run("order by complex struct any", func(b *testing.B) {
		type Config struct {
			Name     string
			Priority int
			Options  []string
		}

		configs := make([]Config, 100)
		for i := 0; i < 100; i++ {
			configs[i] = Config{
				Name:     fmt.Sprintf("Config%d", i),
				Priority: 100 - i,
				Options:  []string{fmt.Sprintf("opt%d_1", i), fmt.Sprintf("opt%d_2", i)},
			}
		}
		var enumerator = FromSliceAny(configs)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })
			_ = ordered.ToSlice()
		}
	})

	b.Run("order by descending any", func(b *testing.B) {
		var enumerator = FromSliceAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderByDescending(comparer.ComparerInt)
			_ = ordered.ToSlice()
		}
	})

	b.Run("order by with nil any enumerator", func(b *testing.B) {
		var enumerator EnumeratorAny[int] = nil

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(comparer.ComparerInt)
			_ = ordered.ToSlice()
		}
	})
}
