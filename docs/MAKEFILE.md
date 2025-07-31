# GenericMap - Make Commands Documentation

This document provides a comprehensive guide to all available Make commands for the GenericMap project.

## Quick Start

```bash
# Run all tests
make test

# Run benchmarks
make bench

# Show all available commands
make help
```

## Test Commands

### `make test`
Runs all tests with verbose output.

```bash
make test
```

**Example Output:**
```
Running tests...
=== RUN   TestNew
--- PASS: TestNew (0.00s)
=== RUN   TestSetAndGet
--- PASS: TestSetAndGet (0.00s)
...
PASS
ok      github.com/costa92/genericmap   0.272s
```

### `make test-short`
Runs only short tests, skipping longer integration tests.

```bash
make test-short
```

### `make test-race`
Runs tests with Go's race detector to identify data race conditions.

```bash
make test-race
```

**Important:** This is crucial for concurrent code validation.

### `make test-coverage`
Generates a comprehensive test coverage report in HTML format.

```bash
make test-coverage
```

**Output:**
- `coverage.out` - Coverage data file
- `coverage.html` - Visual HTML coverage report

**Current Coverage: 98.5%**

### `make test-coverage-func`
Shows function-level test coverage in the terminal.

```bash
make test-coverage-func
```

**Example Output:**
```
github.com/costa92/genericmap/map.go:28:    New                     100.0%
github.com/costa92/genericmap/map.go:60:    Set                     100.0%
github.com/costa92/genericmap/map.go:85:    Get                     100.0%
...
total:                                      (statements)            98.5%
```

## Benchmark Commands

### `make bench`
Runs all benchmark tests with memory allocation statistics.

```bash
make bench
```

**Example Output:**
```
BenchmarkSet-8                   2915026    426.2 ns/op    76 B/op    1 allocs/op
BenchmarkGet-8                  82098064     14.34 ns/op    0 B/op    0 allocs/op
BenchmarkGetKeys-8               1608144    771.9 ns/op   904 B/op    2 allocs/op
BenchmarkRemove-8                5715165    248.5 ns/op     0 B/op    0 allocs/op
BenchmarkConcurrentReadWrite-8   8365099    142.2 ns/op    32 B/op    1 allocs/op
```

### `make bench-cpu`
Runs benchmarks with CPU profiling for performance analysis.

```bash
make bench-cpu
```

**Generates:** `cpu.prof` file for analysis

### `make bench-mem`
Runs benchmarks with memory profiling.

```bash
make bench-mem
```

**Generates:** `mem.prof` file for analysis

### `make bench-compare`
Runs benchmarks multiple times for performance comparison.

```bash
make bench-compare
```

**Output:** `bench.txt` file with detailed results

## Code Quality Commands

### `make fmt`
Formats all Go code according to Go standards.

```bash
make fmt
```

**What it does:**
- Applies `gofmt -s -w .`
- Simplifies code structure
- Ensures consistent formatting

### `make vet`
Runs Go's static analysis tool to find potential issues.

```bash
make vet
```

**Detects:**
- Suspicious constructs
- Potential bugs
- Non-portable code

### `make lint`
Runs golint for code style checking.

```bash
make lint
```

**Note:** Automatically installs golint if not present.

### `make check`
Comprehensive code quality check combining format, vet, and test.

```bash
make check
```

**Equivalent to:**
```bash
make fmt
make vet  
make test
```

## Development Workflow Commands

### `make dev`
Standard development workflow for daily coding.

```bash
make dev
```

**Runs:**
1. Code formatting (`fmt`)
2. Static analysis (`vet`)
3. All tests (`test`)

### `make ci`
Complete CI/CD pipeline validation.

```bash
make ci
```

**Runs:**
1. Module verification (`mod-verify`)
2. Code formatting (`fmt`)
3. Static analysis (`vet`)
4. Race condition testing (`test-race`)
5. Coverage analysis (`test-coverage`)

## Module Management

### `make mod-tidy`
Cleans up and organizes Go module dependencies.

```bash
make mod-tidy
```

### `make mod-verify`
Verifies that dependencies have not been modified.

```bash
make mod-verify
```

### `make mod-download`
Downloads all required modules.

```bash
make mod-download
```

## Build Commands

### `make build`
Builds the package to verify compilation.

```bash
make build
```

## Performance Analysis

### `make profile-cpu`
Analyzes CPU performance using generated profile data.

```bash
make profile-cpu
```

**Prerequisites:** Run `make bench-cpu` first

### `make profile-mem`
Analyzes memory usage patterns.

```bash
make profile-mem
```

**Prerequisites:** Run `make bench-mem` first

## Cleanup Commands

### `make clean`
Removes all generated files and build artifacts.

```bash
make clean
```

**Removes:**
- `coverage.out`
- `coverage.html`
- `cpu.prof`
- `mem.prof`
- `bench.txt`

### `make clean-cache`
Clears Go's build cache.

```bash
make clean-cache
```

## Utility Commands

### `make help`
Displays all available Make targets with descriptions.

```bash
make help
```

## Common Workflows

### Before Committing Code
```bash
make dev
```

### Performance Testing
```bash
make bench
make bench-compare
```

### Full Quality Check
```bash
make ci
```

### Debugging Performance Issues
```bash
make bench-cpu
make profile-cpu
```

### Coverage Analysis
```bash
make test-coverage
# Open coverage.html in browser
```

## Performance Metrics

Current benchmark results on Apple M1:

| Operation | Time/op | Allocs/op | B/op |
|-----------|---------|-----------|------|
| Set | 426.2 ns | 1 | 76 B |
| Get | 14.34 ns | 0 | 0 B |
| GetKeys | 771.9 ns | 2 | 904 B |
| Remove | 248.5 ns | 0 | 0 B |
| Concurrent R/W | 142.2 ns | 1 | 32 B |

## Test Coverage

- **Overall Coverage:** 98.5%
- **All Functions:** 90%+ coverage
- **Critical Paths:** 100% coverage

## Tips

1. **Daily Development:** Use `make dev` for quick validation
2. **Before PR:** Always run `make ci` 
3. **Performance Tuning:** Use `make bench-compare` to measure improvements
4. **Debugging:** Use `make test-race` to catch concurrency issues
5. **Coverage Goals:** Maintain >95% test coverage