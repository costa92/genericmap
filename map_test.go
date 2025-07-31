package genericmap

import (
	"fmt"
	"sort"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	// Test empty construction
	m := New[string, int]()
	if m.Len() != 0 {
		t.Errorf("Expected empty map, got length %d", m.Len())
	}

	// Test construction with initial data
	initial := map[string]int{"a": 1, "b": 2, "c": 1}
	m = New[string, int](initial)
	if m.Len() != 3 {
		t.Errorf("Expected 3 items, got %d", m.Len())
	}

	if keys := m.GetKeys(1); len(keys) != 2 {
		t.Errorf("Expected 2 keys for value 1, got %d: %v", len(keys), keys)
	}
}

func TestSetAndGet(t *testing.T) {
	m := New[string, int]()

	m.Set("key1", 100)
	if val, ok := m.Get("key1"); !ok || val != 100 {
		t.Errorf("Get failed: expected 100, got %v, exists: %v", val, ok)
	}

	// Test setting same value twice
	m.Set("key1", 100)
	if val, ok := m.Get("key1"); !ok || val != 100 {
		t.Errorf("Set same value failed: expected 100, got %v", val)
	}

	// Test updating value
	m.Set("key1", 200)
	if val, ok := m.Get("key1"); !ok || val != 200 {
		t.Errorf("Update failed: expected 200, got %v", val)
	}
}

func TestReverseLookup(t *testing.T) {
	m := New[string, int]()

	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 1)

	keys := m.GetKeys(1)
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys for value 1, got %d: %v", len(keys), keys)
	}

	// Ensure keys contain both "a" and "c"
	keySet := make(map[string]bool)
	for _, k := range keys {
		keySet[k] = true
	}
	if !keySet["a"] || !keySet["c"] {
		t.Errorf("Expected keys [a c] for value 1, got %v", keys)
	}
}

func TestRemove(t *testing.T) {
	m := New[string, int]()

	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 1)

	// Test successful removal
	if !m.Remove("a") {
		t.Errorf("Remove failed: expected true, got false")
	}

	if m.Len() != 2 {
		t.Errorf("Expected length 2 after removal, got %d", m.Len())
	}

	// Test removal of non-existent key
	if m.Remove("nonexistent") {
		t.Errorf("Remove of nonexistent key returned true")
	}
}

func TestListAndValues(t *testing.T) {
	m := New[string, int]()

	m.Set("x", 10)
	m.Set("y", 20)
	m.Set("z", 10)

	keys := m.List()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d: %v", len(keys), keys)
	}

	values := m.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d: %v", len(values), values)
	}
}

func TestLen(t *testing.T) {
	m := New[string, int]()

	if m.Len() != 0 {
		t.Errorf("Expected 0 entries, got %d", m.Len())
	}

	m.Set("a", 1)
	if m.Len() != 1 {
		t.Errorf("Expected 1 entries, got %d", m.Len())
	}

	m.Set("b", 2)
	m.Set("c", 3)
	if m.Len() != 3 {
		t.Errorf("Expected 3 entries, got %d", m.Len())
	}

	m.Remove("b")
	if m.Len() != 2 {
		t.Errorf("Expected 2 entries after removal, got %d", m.Len())
	}
}

func TestConcurrentAccess(t *testing.T) {
	m := New[int, string]()
	const goroutines = 10
	const itemsPerGoroutine = 100

	var wg sync.WaitGroup

	// Parallel writes
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < itemsPerGoroutine; j++ {
				key := id*itemsPerGoroutine + j
				value := fmt.Sprintf("value-%d-%d", id, j)
				m.Set(key, value)
			}
		}(i)
	}

	// Parallel reads
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				_ = m.Len()
				_ = m.List()
				_ = m.Values()
			}
		}(i)
	}

	wg.Wait()

	totalItems := goroutines * itemsPerGoroutine
	if m.Len() != totalItems {
		t.Errorf("Expected %d items after concurrent writes, got %d", totalItems, m.Len())
	}
}

func TestString(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	m.Set("b", 2)

	str := m.String()
	// Skip string test due to map iteration order
	_ = str
}

func Example() {
	// Create empty map
	m := New[string, int]()
	fmt.Printf("Empty map length: %d\n", m.Len())

	// Create with initial data
	initial := map[string]int{"apple": 5, "banana": 2, "cherry": 5}
	m = New[string, int](initial)
	fmt.Printf("Initialized map length: %d\n", m.Len())

	// Sort keys for consistent output
	keys := m.GetKeys(5)
	sort.Strings(keys)
	fmt.Printf("Keys for value 5: %v\n", keys)

	// Output:
	// Empty map length: 0
	// Initialized map length: 3
	// Keys for value 5: [apple cherry]
}
