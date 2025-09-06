package enumerable

import (
	"sort"

	"github.com/ahatornn/enumerable/comparer"
)

// OrderEnumerator represents an ordered sequence of elements that supports lazy sorting
// with multiple sorting levels. This type enables fluent chaining of sorting operations.
//
// OrderEnumerator implements lazy evaluation pattern where:
//   - Sorting rules are accumulated when calling OrderBy/ThenBy methods
//   - Actual sorting computation is deferred until enumeration begins
//   - All sorting levels are applied in a single efficient pass during enumeration
//
// The type supports:
//   - Primary sorting with OrderBy/OrderByDescending
//   - Secondary (and further) sorting with ThenBy/ThenByDescending
//   - Stable sorting that preserves relative order of equal elements
//   - Lazy evaluation with early termination support
//   - Zero memory allocation during rule accumulation
//
// Type Parameters:
//
//	T - the comparable type of elements in the sequence (must satisfy comparable constraint)
//
// ⚠️ Performance characteristics:
//   - Rule accumulation: O(1) time, O(1) memory per operation
//   - Actual sorting: O(n log n) time, O(n) memory during first enumeration
//   - Subsequent enumerations: O(n log n) time, O(n) memory (sorting repeated)
//   - Early termination: supported, but full sort still required initially
//
// ⚠️ Memory usage:
//   - Rules storage: O(k) where k is number of sorting levels
//   - During sorting: O(n) for temporary storage of all elements
//   - After sorting: O(k) for rules, elements released
//
// Notes:
//   - Type T must be comparable for proper sorting comparison
//   - All sorting levels are applied simultaneously during enumeration
//   - Sorting is stable - equal elements maintain relative order
//   - Source enumerator is consumed entirely during first enumeration
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
//   - Works seamlessly with other enumerable operations (Where, Take, Skip, etc.)
type OrderEnumerator[T comparable] struct {

	// source is the original enumerator that provides elements for sorting
	// This enumerator is consumed entirely during the first enumeration
	source Enumerator[T]

	// sortLevels contains accumulated sorting rules in priority order
	// Each level defines a comparison function and sort direction
	// Levels are applied from first (highest priority) to last (lowest priority)
	sortLevels []sortLevel[T]
}

func newOrderEnumerator[T comparable](source Enumerator[T],
	comparer comparer.ComparerFunc[T],
	descending bool) OrderEnumerator[T] {
	return OrderEnumerator[T]{
		source: source,
		sortLevels: []sortLevel[T]{
			{
				comparer:   comparer,
				descending: descending,
			},
		},
	}
}

func (o OrderEnumerator[T]) addSortLevel(comparer comparer.ComparerFunc[T], descending bool) OrderEnumerator[T] {
	newSortLevels := make([]sortLevel[T], len(o.sortLevels)+1)
	copy(newSortLevels, o.sortLevels)
	newSortLevels[len(o.sortLevels)] = sortLevel[T]{
		comparer:   comparer,
		descending: descending,
	}

	return OrderEnumerator[T]{
		source:     o.source,
		sortLevels: newSortLevels,
	}
}

func (o OrderEnumerator[T]) getSortedEnumerator() func(func(T) bool) {
	return func(yield func(T) bool) {
		if o.source == nil {
			return
		}

		var items []T
		o.source(func(item T) bool {
			items = append(items, item)
			return true
		})

		if len(items) == 0 {
			return
		}

		sort.SliceStable(items, func(i, j int) bool {
			for _, level := range o.sortLevels {
				result := level.comparer(items[i], items[j])

				if result != 0 {
					if level.descending {
						return result > 0
					}
					return result < 0
				}
			}
			return false
		})

		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

// OrderEnumeratorAny represents an ordered sequence of elements for any type T,
// including non-comparable types, that supports lazy sorting with multiple sorting levels.
// This type enables fluent chaining of sorting operations for complex types.
//
// OrderEnumeratorAny implements lazy evaluation pattern where:
//   - Sorting rules are accumulated when calling OrderBy/ThenBy methods
//   - Actual sorting computation is deferred until enumeration begins
//   - All sorting levels are applied in a single efficient pass during enumeration
//
// The type supports:
//   - Primary sorting with OrderBy/OrderByDescending using custom comparers
//   - Secondary (and further) sorting with ThenBy/ThenByDescending
//   - Stable sorting that preserves relative order of equal elements
//   - Lazy evaluation with early termination support
//   - Zero memory allocation during rule accumulation
//   - Full compatibility with non-comparable types (structs with slices, maps, functions)
//
// Type Parameters:
//
//	T - any type of elements in the sequence (including non-comparable types)
//
// ⚠️ Performance characteristics:
//   - Rule accumulation: O(1) time, O(1) memory per operation
//   - Actual sorting: O(n log n) time, O(n) memory during first enumeration
//   - Subsequent enumerations: O(n log n) time, O(n) memory (sorting repeated)
//   - Early termination: supported, but full sort still required initially
//
// ⚠️ Memory usage:
//   - Rules storage: O(k) where k is number of sorting levels
//   - During sorting: O(n) for temporary storage of all elements
//   - After sorting: O(k) for rules, elements released
//
// Notes:
//   - Works with any type T, including non-comparable types
//   - Custom comparer functions must implement proper comparison logic
//   - Comparer functions should be deterministic and consistent
//   - All sorting levels are applied simultaneously during enumeration
//   - Sorting is stable - equal elements maintain relative order
//   - Source enumerator is consumed entirely during the first enumeration
//   - Thread-safe for rule accumulation, but enumeration should be single-threaded
//   - Works seamlessly with other enumerable operations (Where, Take, Skip, etc.)
//   - More flexible than OrderEnumerator[T comparable] but potentially slower for simple types
type OrderEnumeratorAny[T any] struct {

	// source is the original enumerator that provides elements for sorting
	// This enumerator is consumed entirely during the first enumeration
	source EnumeratorAny[T]

	// sortLevels contains accumulated sorting rules in priority order
	// Each level defines a comparison function and sort direction
	// Levels are applied from first (highest priority) to last (lowest priority)
	sortLevels []sortLevel[T]
}

func newOrderEnumeratorAny[T any](source EnumeratorAny[T],
	comparer comparer.ComparerFunc[T],
	descending bool) OrderEnumeratorAny[T] {
	return OrderEnumeratorAny[T]{
		source: source,
		sortLevels: []sortLevel[T]{
			{
				comparer:   comparer,
				descending: descending,
			},
		},
	}
}

func (o OrderEnumeratorAny[T]) addSortLevel(comparer comparer.ComparerFunc[T], descending bool) OrderEnumeratorAny[T] {
	newSortLevels := make([]sortLevel[T], len(o.sortLevels)+1)
	copy(newSortLevels, o.sortLevels)
	newSortLevels[len(o.sortLevels)] = sortLevel[T]{
		comparer:   comparer,
		descending: descending,
	}

	return OrderEnumeratorAny[T]{
		source:     o.source,
		sortLevels: newSortLevels,
	}
}

type sortLevel[T any] struct {
	comparer   comparer.ComparerFunc[T]
	descending bool
}

func (o OrderEnumeratorAny[T]) getSortedEnumerator() func(func(T) bool) {
	return func(yield func(T) bool) {
		if o.source == nil {
			return
		}

		var items []T
		o.source(func(item T) bool {
			items = append(items, item)
			return true
		})

		if len(items) == 0 {
			return
		}

		sort.SliceStable(items, func(i, j int) bool {
			for _, level := range o.sortLevels {
				result := level.comparer(items[i], items[j])

				if result != 0 {
					if level.descending {
						return result > 0
					}
					return result < 0
				}
			}
			return false
		})

		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}
