// Package genericmap provides a thread-safe, generic bidirectional map implementation.
// It supports efficient forward lookups (key->value) and reverse lookups (value->keys).
package genericmap

import (
	"fmt"
	"sync"
)

// Map is a thread-safe, generic map with bidirectional lookup capabilities.
// It supports both key-to-value and value-to-keys operations efficiently.
type Map[K comparable, V comparable] struct {
	data       map[K]V
	reverseMap map[V]map[K]struct{}
	mu         sync.RWMutex
}

// New creates a new generic map with optional initial data.
//
// Examples:
//
//	// Create empty map
//	m := New[string, int]()
//
//	// Create with initial data
//	initial := map[string]int{"a": 1, "b": 2}
//	m := New[string, int](initial)
func New[K comparable, V comparable](initialData ...map[K]V) *Map[K, V] {
	m := &Map[K, V]{
		data:       make(map[K]V),
		reverseMap: make(map[V]map[K]struct{}),
	}

	// Populate with initial data if provided
	if len(initialData) > 0 {
		for _, dataMap := range initialData {
			for k, v := range dataMap {
				m.data[k] = v
				if m.reverseMap[v] == nil {
					m.reverseMap[v] = make(map[K]struct{})
				}
				m.reverseMap[v][k] = struct{}{}
			}
		}
	}

	return m
}

// NewWithCapacity creates a new generic map with specified initial capacity.
// This can improve performance when the expected size is known in advance.
func NewWithCapacity[K comparable, V comparable](capacity int) *Map[K, V] {
	return &Map[K, V]{
		data:       make(map[K]V, capacity),
		reverseMap: make(map[V]map[K]struct{}, capacity),
	}
}

// Set adds or updates a key-value pair in the map.
func (m *Map[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Single lookup to check existing value
	oldValue, exists := m.data[key]
	if exists && oldValue == value {
		return // No-op if key already has this value
	}

	// Remove key from old value's reverse map if key exists
	if exists {
		m.removeFromReverseMap(key, oldValue)
	}

	// Add to data and reverse maps
	m.data[key] = value
	if m.reverseMap[value] == nil {
		m.reverseMap[value] = make(map[K]struct{})
	}
	m.reverseMap[value][key] = struct{}{}
}

// Get retrieves the value associated with the key.
// Returns the value and a boolean indicating if the key exists.
func (m *Map[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.data[key]
	return val, ok
}

// GetKeys retrieves all keys associated with a given value.
// Returns a slice of keys that map to the specified value.
func (m *Map[K, V]) GetKeys(value V) []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if keyMap, ok := m.reverseMap[value]; ok {
		result := make([]K, 0, len(keyMap))
		for key := range keyMap {
			result = append(result, key)
		}
		return result
	}
	return []K{}
}

// List returns all keys in the map.
func (m *Map[K, V]) List() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := make([]K, len(m.data))
	i := 0
	for k := range m.data {
		keys[i] = k
		i++
	}
	return keys
}

// Values returns all values in the map.
func (m *Map[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	values := make([]V, len(m.data))
	i := 0
	for _, v := range m.data {
		values[i] = v
		i++
	}
	return values
}

// Remove removes a key-value pair from the map.
// Returns true if the key existed and was removed, false otherwise.
func (m *Map[K, V]) Remove(key K) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if value, exists := m.data[key]; exists {
		delete(m.data, key)
		m.removeFromReverseMap(key, value)
		return true
	}
	return false
}

// Len returns the number of key-value pairs in the map.
func (m *Map[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data)
}

// String returns a string representation of the map.
func (m *Map[K, V]) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return fmt.Sprintf("Map[%d]{%v}", len(m.data), m.data)
}

// removeFromReverseMap removes a key from the reverse map for a given value.
// This is an internal method and assumes the caller holds the appropriate lock.
func (m *Map[K, V]) removeFromReverseMap(key K, value V) {
	if keyMap, exists := m.reverseMap[value]; exists {
		delete(keyMap, key)
		if len(keyMap) == 0 {
			delete(m.reverseMap, value)
		}
	}
}
