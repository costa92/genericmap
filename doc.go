// Package genericmap provides a thread-safe, generic bidirectional map implementation.
// 
// The genericmap package offers a highly efficient map with both forward (key->value)
// and reverse (value->keys) lookup capabilities, designed specifically for Go 1.18+.
//
// Features:
//
// • Thread-safe concurrent access using sync.RWMutex
// • Generic type support for both keys and values
// • Efficient O(1) reverse lookups via prebuilt reverse index
// • Memory optimized with swap-and-truncate removal
// • Fully compatible with Go idioms
// 
// Basic Usage:
//
//  import "your-project/pkg/genericmap"
//
//  // Create empty map
//  m := genericmap.New[string, int]()
//  m.Set("apple", 5)
//  m.Set("banana", 2)
//
//  // Create with initial data
//  initial := map[string]int{"a": 1, "b": 2}
//  m := genericmap.New[string, int](initial)
//
//  // Forward lookup
//  value, exists := m.Get("apple")
//
//  // Reverse lookup
//  keys := m.GetKeys(5) // Gets all keys with value 5
//
//  // Thread-safe operations
//  m.Len()     // Get map size
//  m.List()    // Get all keys
//  m.Values()  // Get all values
//  m.Remove("apple")
//
// Concurrent Usage:
//
//  The map is safe for concurrent use from multiple goroutines without external locking.
//  Reader goroutines use read locks, writer goroutines use write locks.
//
// Example:
//
//  m := genericmap.New[int, string]()
//  
//  go func() { m.Set(1, "admin") }()
//  go func() { fmt.Println(m.Len()) }()
//  go func() { keys := m.GetKeys("admin") }()
// 
// Performance Characteristics:
//
// • Set:         O(1) average case for new keys
// • Get:         O(1) guaranteed
// • GetKeys:     O(k) where k is number of keys for value
// • Remove:      O(1) average case
// • Len:         O(1)
// 
// Memory Usage:
//
// The map maintains an additional reverse index, which uses O(n) space where n
// is the number of unique values in the map.
package genericmap