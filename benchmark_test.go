package genericmap

import (
	"fmt"
	"testing"
)

// BenchmarkSet measures the performance of Set operations
func BenchmarkSet(b *testing.B) {
	m := NewWithCapacity[int, string](b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Set(i, fmt.Sprintf("value-%d", i%100)) // Create some duplicate values
	}
}

// BenchmarkGet measures the performance of Get operations
func BenchmarkGet(b *testing.B) {
	m := NewWithCapacity[int, string](1000)

	// Setup data
	for i := 0; i < 1000; i++ {
		m.Set(i, fmt.Sprintf("value-%d", i%100))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(i % 1000)
	}
}

// BenchmarkGetKeys measures the performance of reverse lookup operations
func BenchmarkGetKeys(b *testing.B) {
	m := NewWithCapacity[int, string](1000)

	// Setup data with duplicate values
	for i := 0; i < 1000; i++ {
		m.Set(i, fmt.Sprintf("value-%d", i%10)) // 10 different values, 100 keys each
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.GetKeys(fmt.Sprintf("value-%d", i%10))
	}
}

// BenchmarkRemove measures the performance of Remove operations
func BenchmarkRemove(b *testing.B) {
	// Setup a fresh map for each benchmark run
	m := NewWithCapacity[int, string](b.N)
	for i := 0; i < b.N; i++ {
		m.Set(i, fmt.Sprintf("value-%d", i%100))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Remove(i)
	}
}

// BenchmarkConcurrentReadWrite measures concurrent performance
func BenchmarkConcurrentReadWrite(b *testing.B) {
	m := NewWithCapacity[int, string](1000)

	// Setup initial data
	for i := 0; i < 1000; i++ {
		m.Set(i, fmt.Sprintf("value-%d", i%100))
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%3 == 0 {
				// 33% writes
				m.Set(i%1000, fmt.Sprintf("value-%d", i%100))
			} else if i%3 == 1 {
				// 33% reads
				_, _ = m.Get(i % 1000)
			} else {
				// 33% reverse lookups
				_ = m.GetKeys(fmt.Sprintf("value-%d", i%100))
			}
			i++
		}
	})
}
