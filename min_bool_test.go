package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMinBool(t *testing.T) {
	t.Run("min bool from mixed values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		min, ok := enumerator.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool from all true values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, true})

		min, ok := enumerator.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != true {
			t.Errorf("Expected min true, got %v", min)
		}
	})

	t.Run("min bool from all false values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, false, false})

		min, ok := enumerator.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool single true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true})

		min, ok := enumerator.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != true {
			t.Errorf("Expected min true, got %v", min)
		}
	})

	t.Run("min bool single false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false})

		min, ok := enumerator.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{})

		min, ok := enumerator.MinBool(selector.Bool)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if min != false {
			t.Errorf("Expected min false for empty slice, got %v", min)
		}
	})

	t.Run("min bool nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[bool] = nil

		min, ok := enumerator.MinBool(selector.Bool)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if min != false {
			t.Errorf("Expected min false for nil enumerator, got %v", min)
		}
	})

	t.Run("min bool with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false})

		min, ok := enumerator.MinBool(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if min != false {
			t.Errorf("Expected min false for nil keySelector, got %v", min)
		}
	})

	t.Run("min bool with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Feature struct {
			Name     string
			Enabled  bool
			Priority int
		}

		features := []Feature{
			{Name: "FeatureA", Enabled: true, Priority: 1},
			{Name: "FeatureB", Enabled: false, Priority: 2},
			{Name: "FeatureC", Enabled: true, Priority: 3},
		}

		enumerator := FromSlice(features)
		min, ok := enumerator.MinBool(func(f Feature) bool { return f.Enabled })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false})

		min, ok := enumerator.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})
}

func TestMinBoolStruct(t *testing.T) {
	t.Run("min bool from struct field", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name     string
			IsActive bool
			Age      int
		}

		users := []User{
			{Name: "Alice", IsActive: true, Age: 30},
			{Name: "Bob", IsActive: false, Age: 25},
			{Name: "Charlie", IsActive: true, Age: 35},
		}

		enumerator := FromSlice(users)
		min, ok := enumerator.MinBool(func(u User) bool { return u.IsActive })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool from struct with slice field", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings []string
			IsValid  bool
		}

		configs := []Config{
			{Settings: []string{"a", "b"}, IsValid: true},
			{Settings: []string{"c"}, IsValid: false},
			{Settings: []string{"d", "e", "f"}, IsValid: true},
		}

		enumerator := FromSliceAny(configs)
		min, ok := enumerator.MinBool(func(c Config) bool { return c.IsValid })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})
}

func TestMinBoolEdgeCases(t *testing.T) {
	t.Run("min bool early termination optimization", func(t *testing.T) {
		t.Parallel()
		bools := make([]bool, 1000)
		bools[0] = false
		for i := 1; i < 1000; i++ {
			bools[i] = true
		}

		enumerator := FromSlice(bools)
		min, ok := enumerator.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool without early termination", func(t *testing.T) {
		t.Parallel()
		bools := make([]bool, 1000)
		for i := 0; i < 999; i++ {
			bools[i] = true
		}
		bools[999] = false

		enumerator := FromSlice(bools)
		min, ok := enumerator.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})
}

func TestMinBoolWithOperations(t *testing.T) {
	t.Run("min bool after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})
		filtered := enumerator.Where(func(b bool) bool { return b })

		min, ok := filtered.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != true {
			t.Errorf("Expected min true (from filtered true values), got %v", min)
		}
	})

	t.Run("min bool after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, false, true, false, true})
		taken := enumerator.Take(3)

		min, ok := taken.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false (from first 3 elements), got %v", min)
		}
	})

	t.Run("min bool after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, false, true, false, true})
		skipped := enumerator.Skip(2)

		min, ok := skipped.MinBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false (after skipping 2 elements), got %v", min)
		}
	})
}

func TestMinBoolCustomKeySelector(t *testing.T) {
	t.Run("min bool by condition", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		min, ok := enumerator.MinBool(func(x int) bool { return x%2 == 0 })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool by string empty check", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "", "world", "", "go"})

		min, ok := enumerator.MinBool(func(s string) bool { return s != "" })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})
}

func TestMinBoolNonComparable(t *testing.T) {
	t.Run("min bool from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name    []string
			IsAdmin bool
		}

		people := []Person{
			{Name: []string{"Alice"}, IsAdmin: true},
			{Name: []string{"Bob"}, IsAdmin: false},
			{Name: []string{"Charlie"}, IsAdmin: true},
		}

		enumerator := FromSliceAny(people)
		min, ok := enumerator.MinBool(func(p Person) bool { return p.IsAdmin })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Enabled  bool
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Enabled: true},
			{Settings: map[string]interface{}{"b": 2}, Enabled: false},
			{Settings: map[string]interface{}{"c": 3}, Enabled: true},
		}

		enumerator := FromSliceAny(configs)
		min, ok := enumerator.MinBool(func(c Config) bool { return c.Enabled })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})

	t.Run("min bool with function field struct", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Callback func()
			Active   bool
		}

		handlers := []Handler{
			{Callback: func() {}, Active: true},
			{Callback: func() {}, Active: false},
			{Callback: func() {}, Active: true},
		}

		enumerator := FromSliceAny(handlers)
		min, ok := enumerator.MinBool(func(h Handler) bool { return h.Active })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != false {
			t.Errorf("Expected min false, got %v", min)
		}
	})
}

func BenchmarkMinBool(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []bool{true, false, true, false, true}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinBool(selector.Bool)
			if !ok || result != false {
				b.Fatalf("Expected false, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		// false в начале - должно быть быстро
		items := make([]bool, 1000)
		items[0] = false
		for i := 1; i < 1000; i++ {
			items[i] = true
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinBool(selector.Bool)
			if !ok || result != false {
				b.Fatalf("Expected false, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		// false в конце - должно пройти все элементы
		items := make([]bool, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = true
		}
		items[9999] = false

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinBool(selector.Bool)
			if !ok || result != false {
				b.Fatalf("Expected false, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]bool{})
			result, ok := enumerator.MinBool(selector.Bool)
			if ok || result != false {
				b.Fatalf("Expected false and false, got %v, ok: %v", result, ok)
			}
		}
	})
}
