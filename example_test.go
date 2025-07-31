package genericmap

import (
	"fmt"
	"sort"
	"sync"
)

// ExampleBasicUsage demonstrates basic usage of the genericmap package.
func ExampleMap_basicUsage() {
	// Create a new map
	m := New[string, int]()

	// Add key-value pairs
	m.Set("apple", 5)
	m.Set("banana", 2)
	m.Set("cherry", 5)

	// Retrieve values
	if value, exists := m.Get("apple"); exists {
		fmt.Printf("apple has %d items\n", value)
	}

	// List all keys (sorted for consistent output)
	keys := m.List()
	sort.Strings(keys)
	fmt.Printf("All keys: %v\n", keys)

	// List all values (sorted for consistent output)
	values := m.Values()
	sort.Ints(values)
	fmt.Printf("All values: %v\n", values)

	// Reverse lookup - find keys for a value (sorted for consistent output)
	keysWithValue5 := m.GetKeys(5)
	sort.Strings(keysWithValue5)
	fmt.Printf("Keys with value 5: %v\n", keysWithValue5)

	// Get map size
	fmt.Printf("Map contains %d items\n", m.Len())

	// Remove an item
	removed := m.Remove("banana")
	fmt.Printf("Removed banana: %t, new size: %d\n", removed, m.Len())

	// Output:
	// apple has 5 items
	// All keys: [apple banana cherry]
	// All values: [2 5 5]
	// Keys with value 5: [apple cherry]
	// Map contains 3 items
	// Removed banana: true, new size: 2
}

// ExampleInitialData demonstrates creating a map with initial data.
func ExampleMap_initialData() {
	// Create with initial data
	initial := map[string]int{
		"apple":  5,
		"banana": 2,
		"cherry": 5,
	}
	m := New[string, int](initial)

	fmt.Printf("Initial data loaded: %d items\n", m.Len())
	keys := m.GetKeys(5)
	fmt.Printf("Has value 5 with %d keys\n", len(keys))

	// Output:
	// Initial data loaded: 3 items
	// Has value 5 with 2 keys
}

// ExampleConcurrent demonstrates thread-safe concurrent usage.
func ExampleMap_concurrent() {
	m := New[int, string]()
	var wg sync.WaitGroup

	// Multiple goroutines writing
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			m.Set(id, fmt.Sprintf("item-%d", id))
		}(i)
	}

	// Multiple goroutines reading
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				_ = m.Len()
				_ = m.List()
				_ = m.GetKeys("test")
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Concurrent operations completed: %d items\n", m.Len())

	// Output:
	// Concurrent operations completed: 5 items
}

// ExampleStringType demonstrates using string keys and values.
func ExampleMap_stringType() {
	m := New[string, string]()
	m.Set("user1", "alice@example.com")
	m.Set("user2", "bob@example.com")
	m.Set("user3", "alice@example.com")

	emails := m.GetKeys("alice@example.com")
	fmt.Printf("Users with email alice@example.com: %v\n", emails)

	// Output:
	// Users with email alice@example.com: [user1 user3]
}

// ExampleUserIDMapping demonstrates a mapping users to groups.
func ExampleMap_userGroupMapping() {
	// Map user IDs to group names
	userGroups := New[int, string]()

	userGroups.Set(1001, "admins")
	userGroups.Set(1002, "users")
	userGroups.Set(1003, "moderators")
	userGroups.Set(1004, "users")
	userGroups.Set(1005, "admins")

	// Find all users in admins group
	adminUsers := userGroups.GetKeys("admins")
	fmt.Printf("Admin user IDs: %v\n", adminUsers)

	// Count users in each group
	fmt.Printf("Total groups: %d\n", len(userGroups.Values()))

	// Count unique groups properly
	uniqueGroups := make(map[string]bool)
	for _, group := range userGroups.Values() {
		uniqueGroups[group] = true
	}
	fmt.Printf("Unique groups: %d\n", len(uniqueGroups))

	// Output:
	// Admin user IDs: [1001 1005]
	// Total groups: 5
	// Unique groups: 3
}
