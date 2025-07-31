# GenericMap

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-blue.svg)](https://golang.org/)
[![Test Coverage](https://img.shields.io/badge/coverage-98.5%25-brightgreen.svg)](./MAKEFILE.md)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/costa92/genericmap)](https://goreportcard.com/report/github.com/costa92/genericmap)

A high-performance, thread-safe, generic bidirectional map implementation in Go with O(1) operations and comprehensive reverse lookup capabilities.

## ✨ Features

- 🚀 **High Performance**: O(1) set, get, and remove operations
- 🔄 **Bidirectional Lookup**: Efficient forward (key→value) and reverse (value→keys) operations
- 🔒 **Thread-Safe**: Built-in concurrent access support with RWMutex
- 🎯 **Generic**: Type-safe with Go generics for any comparable types
- 📊 **Memory Optimized**: Reduced allocations and GC pressure
- 🧪 **Well Tested**: 98.5% test coverage with comprehensive benchmarks

## 🚀 Quick Start

### Installation

```bash
go get github.com/costa92/genericmap
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/costa92/genericmap"
)

func main() {
    // Create a new map
    m := genericmap.New[string, int]()
    
    // Add key-value pairs
    m.Set("apple", 5)
    m.Set("banana", 2)
    m.Set("cherry", 5)
    
    // Get value by key
    if value, exists := m.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value) // apple: 5
    }
    
    // Reverse lookup - find all keys with value 5
    keys := m.GetKeys(5)
    fmt.Printf("Items with value 5: %v\n", keys) // [apple cherry]
    
    // Remove item
    removed := m.Remove("banana")
    fmt.Printf("Removed banana: %t\n", removed) // true
}
```

## 📖 API Documentation

### Creating Maps

```go
// Create empty map
m := genericmap.New[string, int]()

// Create with initial data
initial := map[string]int{"a": 1, "b": 2}
m := genericmap.New[string, int](initial)

// Create with capacity (performance optimization)
m := genericmap.NewWithCapacity[string, int](1000)
```

### Core Operations

```go
// Set/Update
m.Set("key", 100)

// Get
value, exists := m.Get("key")

// Remove
removed := m.Remove("key")

// Size
size := m.Len()
```

### Bidirectional Lookup

```go
// Get all keys for a specific value
keys := m.GetKeys(100)

// List all keys
allKeys := m.List()

// List all values
allValues := m.Values()
```

### Thread Safety

```go
// Safe for concurrent use
var wg sync.WaitGroup

// Multiple goroutines can safely access the map
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        m.Set(fmt.Sprintf("key-%d", id), id)
        value, _ := m.Get(fmt.Sprintf("key-%d", id))
        keys := m.GetKeys(id)
        fmt.Printf("Goroutine %d: value=%d, keys=%v\n", id, value, keys)
    }(i)
}

wg.Wait()
```

## 🏎️ Performance

### Benchmark Results (Apple M1)

| Operation | Time/op | Allocs/op | B/op |
|-----------|---------|-----------|------|
| Set | 426.2 ns | 1 | 76 B |
| Get | 14.34 ns | 0 | 0 B |
| GetKeys | 771.9 ns | 2 | 904 B |
| Remove | 248.5 ns | 0 | 0 B |
| Concurrent R/W | 142.2 ns | 1 | 32 B |

### Performance Optimizations

- **O(1) Operations**: Set, Get, and Remove operations are constant time
- **Optimized Reverse Lookup**: Uses `map[V]map[K]struct{}` instead of `map[V][]K` for O(1) removal
- **Reduced Allocations**: Minimal memory allocations during operations
- **Single Map Lookup**: Eliminates redundant lookups in Set operations

## 📋 Use Cases

### User-Group Mapping

```go
// Map user IDs to group names with reverse lookup
userGroups := genericmap.New[int, string]()

userGroups.Set(1001, "admin")
userGroups.Set(1002, "user")
userGroups.Set(1003, "admin")

// Find all admin users
adminUsers := userGroups.GetKeys("admin") // [1001, 1003]
```

### Cache with Reverse Index

```go
// Cache with ability to find keys by cached values
cache := genericmap.New[string, []byte]()

cache.Set("user:123", []byte("user data"))
cache.Set("post:456", []byte("post data"))

// Find all keys containing specific data
keys := cache.GetKeys([]byte("user data")) // ["user:123"]
```

### Configuration Management

```go
// Manage configuration keys with grouping
config := genericmap.New[string, string]()

config.Set("db.host", "production")
config.Set("cache.host", "production")
config.Set("queue.host", "staging")

// Find all production services
prodServices := config.GetKeys("production") // ["db.host", "cache.host"]
```

## 🛠️ Development

### Prerequisites

- Go 1.25 or higher
- Make (optional, for using Makefile commands)

### Building and Testing

```bash
# Run all tests
make test

# Run benchmarks
make bench

# Run tests with coverage
make test-coverage

# Format and lint code
make check

# Show all available commands
make help
```

### Project Structure

```
├── map.go              # Core implementation
├── map_test.go         # Unit tests
├── example_test.go     # Example tests
├── benchmark_test.go   # Performance benchmarks
├── Makefile           # Build automation
├── MAKEFILE.md        # English Make commands documentation
├── MAKEFILE_CN.md     # Chinese Make commands documentation
└── README.md          # This file
```

## 🧪 Testing

The project maintains high code quality with comprehensive testing:

- **Unit Tests**: Complete coverage of all functionality
- **Example Tests**: Documented usage examples
- **Benchmark Tests**: Performance regression testing
- **Race Tests**: Concurrent access validation
- **Coverage**: 98.5% test coverage

Run tests:

```bash
# Basic tests
go test -v

# With race detection
go test -race -v

# With coverage
go test -cover -v
```

## 🔧 Advanced Usage

### Custom Types

```go
type UserID int
type GroupName string

userMap := genericmap.New[UserID, GroupName]()
userMap.Set(UserID(1001), GroupName("admin"))
```

### Performance Tuning

```go
// Pre-allocate capacity for better performance
expectedSize := 10000
m := genericmap.NewWithCapacity[string, int](expectedSize)

// Batch operations
for i := 0; i < expectedSize; i++ {
    m.Set(fmt.Sprintf("key-%d", i), i)
}
```

### Error Handling

```go
// Check if key exists before operations
if value, exists := m.Get("key"); exists {
    // Process existing value
    fmt.Printf("Found: %v\n", value)
} else {
    // Handle missing key
    fmt.Println("Key not found")
}

// Check removal success
if removed := m.Remove("key"); removed {
    fmt.Println("Successfully removed")
} else {
    fmt.Println("Key was not present")
}
```

## 📊 Comparison with Alternatives

| Feature | GenericMap | sync.Map | Regular map + mutex |
|---------|------------|----------|-------------------|
| Type Safety | ✅ Generics | ❌ interface{} | ✅ Typed |
| Reverse Lookup | ✅ O(1) | ❌ Not supported | ❌ O(n) scan |
| Performance | ⚡ Optimized | 🐌 Interface overhead | 💾 Memory efficient |
| Thread Safety | ✅ Built-in | ✅ Built-in | ⚙️ Manual |
| API Simplicity | ✅ Clean | ❌ Complex | ✅ Simple |

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Workflow

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Run tests (`make dev`)
4. Commit your changes (`git commit -am 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

### Code Style

- Follow Go conventions and best practices
- Run `make fmt` to format code
- Ensure `make ci` passes before submitting PR
- Maintain test coverage above 95%

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Go team for excellent generics implementation
- Community feedback and contributions
- Performance optimization techniques from the Go community

## 📚 Documentation

- [English Make Commands](./MAKEFILE.md)
- [中文 Make 命令文档](./MAKEFILE_CN.md)
- [Go Package Documentation](https://pkg.go.dev/github.com/costa92/genericmap)

## 🔗 Links

- [GitHub Repository](https://github.com/costa92/genericmap)
- [Go Package Index](https://pkg.go.dev/github.com/costa92/genericmap)
- [Issue Tracker](https://github.com/costa92/genericmap/issues)

---

<p align="center">
  <strong>GenericMap</strong> - High-performance bidirectional maps for Go
</p>