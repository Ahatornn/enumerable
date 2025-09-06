package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestThenBy(t *testing.T) {
	t.Run("then by ascending with integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		ordered := enumerator.OrderBy(func(a, b int) int {
			modA, modB := a%2, b%2
			if modA < modB {
				return -1
			}
			if modA > modB {
				return 1
			}
			return 0
		}).ThenBy(comparer.ComparerInt)

		result := ordered.ToSlice()
		expected := []int{2, 4, 1, 3, 5}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("then by descending with integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		ordered := enumerator.OrderBy(func(a, b int) int {
			modA, modB := a%2, b%2
			if modA < modB {
				return -1
			}
			if modA > modB {
				return 1
			}
			return 0
		}).ThenByDescending(comparer.ComparerInt)

		result := ordered.ToSlice()
		expected := []int{4, 2, 5, 3, 1}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("then by with strings", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			FirstName string
			LastName  string
		}

		people := []Person{
			{FirstName: "John", LastName: "Smith"},
			{FirstName: "Jane", LastName: "Smith"},
			{FirstName: "Alice", LastName: "Johnson"},
			{FirstName: "Bob", LastName: "Johnson"},
		}
		enumerator := FromSlice(people)
		ordered := enumerator.OrderBy(func(a, b Person) int {
			return compareStrings(a.LastName, b.LastName)
		}).ThenBy(func(a, b Person) int {
			return compareStrings(a.FirstName, b.FirstName)
		})

		result := ordered.ToSlice()
		expectedFirstNames := []string{"Alice", "Bob", "Jane", "John"}
		expectedLastNames := []string{"Johnson", "Johnson", "Smith", "Smith"}

		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		for i, expectedFirst := range expectedFirstNames {
			if result[i].FirstName != expectedFirst {
				t.Errorf("Expected first name %s at index %d, got %s", expectedFirst, i, result[i].FirstName)
			}
		}

		for i, expectedLast := range expectedLastNames {
			if result[i].LastName != expectedLast {
				t.Errorf("Expected last name %s at index %d, got %s", expectedLast, i, result[i].LastName)
			}
		}
	})

	t.Run("then by descending with strings", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Category string
			Name     string
			Price    int
		}

		products := []Product{
			{Category: "Electronics", Name: "Laptop", Price: 1200},
			{Category: "Electronics", Name: "Phone", Price: 800},
			{Category: "Books", Name: "Novel", Price: 15},
			{Category: "Books", Name: "Textbook", Price: 100},
		}
		enumerator := FromSlice(products)
		ordered := enumerator.OrderBy(func(a, b Product) int {
			return compareStrings(a.Category, b.Category)
		}).ThenByDescending(func(a, b Product) int {
			return a.Price - b.Price
		})

		result := ordered.ToSlice()
		expectedCategories := []string{"Books", "Books", "Electronics", "Electronics"}
		expectedNames := []string{"Textbook", "Novel", "Laptop", "Phone"}
		expectedPrices := []int{100, 15, 1200, 800}

		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
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

		for i, expectedPrice := range expectedPrices {
			if result[i].Price != expectedPrice {
				t.Errorf("Expected price %d at index %d, got %d", expectedPrice, i, result[i].Price)
			}
		}
	})

	t.Run("then by with OrderEnumeratorAny", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name     string
			Priority int
			Options  []string
		}

		configs := []Config{
			{Name: "HighA", Priority: 1, Options: []string{"opt1"}},
			{Name: "HighB", Priority: 1, Options: []string{"opt2"}},
			{Name: "MediumA", Priority: 2, Options: []string{"opt3"}},
			{Name: "MediumB", Priority: 2, Options: []string{"opt4"}},
			{Name: "Low", Priority: 3, Options: []string{"opt5"}},
		}
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int {
			return a.Priority - b.Priority
		}).ThenBy(func(a, b Config) int {
			return compareStrings(a.Name, b.Name)
		})

		result := ordered.ToSlice()
		expectedNames := []string{"HighA", "HighB", "MediumA", "MediumB", "Low"}
		expectedPriorities := []int{1, 1, 2, 2, 3}

		if len(result) != 5 {
			t.Fatalf("Expected length 5, got %d", len(result))
		}

		for i, expectedName := range expectedNames {
			if result[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, result[i].Name)
			}
		}

		for i, expectedPriority := range expectedPriorities {
			if result[i].Priority != expectedPriority {
				t.Errorf("Expected priority %d at index %d, got %d", expectedPriority, i, result[i].Priority)
			}
		}
	})

	t.Run("then by descending with OrderEnumeratorAny", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Department string
			Name       string
			Score      float64
			Tags       []string
		}

		users := []User{
			{Department: "IT", Name: "Alice", Score: 95.5, Tags: []string{"dev"}},
			{Department: "IT", Name: "Bob", Score: 87.0, Tags: []string{"qa"}},
			{Department: "HR", Name: "Charlie", Score: 92.0, Tags: []string{"manager"}},
			{Department: "HR", Name: "Diana", Score: 88.5, Tags: []string{"recruiter"}},
		}
		var enumerator = FromSliceAny(users)

		ordered := enumerator.OrderBy(func(a, b User) int {
			return compareStrings(a.Department, b.Department)
		}).ThenByDescending(func(a, b User) int {
			if a.Score < b.Score {
				return -1
			}
			if a.Score > b.Score {
				return 1
			}
			return 0
		})

		result := ordered.ToSlice()
		expectedDepartments := []string{"HR", "HR", "IT", "IT"}
		expectedNames := []string{"Charlie", "Diana", "Alice", "Bob"}
		expectedScores := []float64{92.0, 88.5, 95.5, 87.0}

		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		for i, expectedDepartment := range expectedDepartments {
			if result[i].Department != expectedDepartment {
				t.Errorf("Expected department %s at index %d, got %s", expectedDepartment, i, result[i].Department)
			}
		}

		for i, expectedName := range expectedNames {
			if result[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, result[i].Name)
			}
		}

		for i, expectedScore := range expectedScores {
			if result[i].Score != expectedScore {
				t.Errorf("Expected score %f at index %d, got %f", expectedScore, i, result[i].Score)
			}
		}
	})

	t.Run("multiple then by chaining", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Category string
			Type     string
			Name     string
			Value    int
		}

		records := []Record{
			{Category: "A", Type: "X", Name: "Second", Value: 10},
			{Category: "A", Type: "X", Name: "First", Value: 20},
			{Category: "A", Type: "Y", Name: "Third", Value: 15},
			{Category: "B", Type: "X", Name: "Fourth", Value: 25},
			{Category: "B", Type: "Y", Name: "Fifth", Value: 5},
		}
		enumerator := FromSlice(records)

		// Три уровня сортировки
		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return compareStrings(a.Type, b.Type)
		}).ThenBy(func(a, b Record) int {
			return compareStrings(a.Name, b.Name)
		})

		result := ordered.ToSlice()
		expectedCategories := []string{"A", "A", "A", "B", "B"}
		expectedTypes := []string{"X", "X", "Y", "X", "Y"}
		expectedNames := []string{"First", "Second", "Third", "Fourth", "Fifth"}

		if len(result) != 5 {
			t.Fatalf("Expected length 5, got %d", len(result))
		}

		for i, expectedCategory := range expectedCategories {
			if result[i].Category != expectedCategory {
				t.Errorf("Expected category %s at index %d, got %s", expectedCategory, i, result[i].Category)
			}
		}

		for i, expectedType := range expectedTypes {
			if result[i].Type != expectedType {
				t.Errorf("Expected type %s at index %d, got %s", expectedType, i, result[i].Type)
			}
		}

		for i, expectedName := range expectedNames {
			if result[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, result[i].Name)
			}
		}
	})

	t.Run("mixed ascending and descending then by", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Group string
			Score int
			Name  string
		}

		items := []Item{
			{Group: "A", Score: 100, Name: "LowScore"},
			{Group: "A", Score: 200, Name: "HighScore"},
			{Group: "B", Score: 150, Name: "MediumScore"},
			{Group: "B", Score: 250, Name: "VeryHighScore"},
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b Item) int {
			return compareStrings(a.Group, b.Group)
		}).ThenByDescending(func(a, b Item) int {
			return a.Score - b.Score
		}).ThenBy(func(a, b Item) int {
			return compareStrings(a.Name, b.Name)
		})

		result := ordered.ToSlice()
		expectedGroups := []string{"A", "A", "B", "B"}
		expectedScores := []int{200, 100, 250, 150}
		expectedNames := []string{"HighScore", "LowScore", "VeryHighScore", "MediumScore"}

		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		for i, expectedGroup := range expectedGroups {
			if result[i].Group != expectedGroup {
				t.Errorf("Expected group %s at index %d, got %s", expectedGroup, i, result[i].Group)
			}
		}

		for i, expectedScore := range expectedScores {
			if result[i].Score != expectedScore {
				t.Errorf("Expected score %d at index %d, got %d", expectedScore, i, result[i].Score)
			}
		}

		for i, expectedName := range expectedNames {
			if result[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, result[i].Name)
			}
		}
	})

	t.Run("then by with empty result", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b }).
			ThenBy(func(a, b int) int { return a - b })

		result := ordered.ToSlice()
		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("then by with single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b }).
			ThenBy(func(a, b int) int { return a - b })

		result := ordered.ToSlice()
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("Expected [42], got %v", result)
		}
	})

	t.Run("then by preserves stability at each level", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Primary   int
			Secondary int
			Index     int
		}

		items := []Item{
			{Primary: 1, Secondary: 2, Index: 1},
			{Primary: 1, Secondary: 1, Index: 2},
			{Primary: 1, Secondary: 2, Index: 3},
			{Primary: 1, Secondary: 1, Index: 4},
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Primary - b.Primary }).
			ThenBy(func(a, b Item) int { return a.Secondary - b.Secondary })

		result := ordered.ToSlice()
		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		// Проверяем порядок Secondary
		if result[0].Secondary != 1 || result[1].Secondary != 1 || result[2].Secondary != 2 || result[3].Secondary != 2 {
			t.Errorf("Secondary values not in correct order: %v", result)
		}

		if result[0].Index != 2 || result[1].Index != 4 {
			t.Errorf("Stability not preserved for Secondary=1: got indices %d, %d", result[0].Index, result[1].Index)
		}

		if result[2].Index != 1 || result[3].Index != 3 {
			t.Errorf("Stability not preserved for Secondary=2: got indices %d, %d", result[2].Index, result[3].Index)
		}
	})
}

func BenchmarkThenBy(b *testing.B) {
	b.Run("then by integers", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b int) int {
				modA, modB := a%2, b%2
				if modA < modB {
					return -1
				}
				if modA > modB {
					return 1
				}
				return 0
			}).ThenBy(func(a, b int) int { return a - b })
			_ = ordered.ToSlice()
		}
	})

	b.Run("then by strings", func(b *testing.B) {
		type Person struct {
			FirstName string
			LastName  string
		}

		people := make([]Person, 100)
		for i := 0; i < 100; i++ {
			people[i] = Person{
				FirstName: fmt.Sprintf("First%d", i),
				LastName:  fmt.Sprintf("Last%d", i%10),
			}
		}
		enumerator := FromSlice(people)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b Person) int {
				return compareStrings(a.LastName, b.LastName)
			}).ThenBy(func(a, b Person) int {
				return compareStrings(a.FirstName, b.FirstName)
			})
			_ = ordered.ToSlice()
		}
	})

	b.Run("multiple then by chaining", func(b *testing.B) {
		type Record struct {
			Category string
			Type     string
			Name     string
		}

		records := make([]Record, 200)
		for i := 0; i < 200; i++ {
			records[i] = Record{
				Category: fmt.Sprintf("Cat%d", i%5),
				Type:     fmt.Sprintf("Type%d", i%10),
				Name:     fmt.Sprintf("Name%d", i),
			}
		}
		enumerator := FromSlice(records)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b Record) int {
				return compareStrings(a.Category, b.Category)
			}).ThenBy(func(a, b Record) int {
				return compareStrings(a.Type, b.Type)
			}).ThenBy(func(a, b Record) int {
				return compareStrings(a.Name, b.Name)
			})
			_ = ordered.ToSlice()
		}
	})

	b.Run("then by descending", func(b *testing.B) {
		type Product struct {
			Category string
			Price    int
		}

		products := make([]Product, 100)
		for i := 0; i < 100; i++ {
			products[i] = Product{
				Category: fmt.Sprintf("Cat%d", i%5),
				Price:    i * 10,
			}
		}
		enumerator := FromSlice(products)

		for i := 0; i < b.N; i++ {
			ordered := enumerator.OrderBy(func(a, b Product) int {
				return compareStrings(a.Category, b.Category)
			}).ThenByDescending(func(a, b Product) int {
				return a.Price - b.Price
			})
			_ = ordered.ToSlice()
		}
	})
}
