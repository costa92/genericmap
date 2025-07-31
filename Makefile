# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt

# Package name
PACKAGE=github.com/costa92/genericmap

# Default target
.PHONY: all
all: test

# Test commands
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

.PHONY: test-short
test-short:
	@echo "Running short tests..."
	$(GOTEST) -short -v ./...

.PHONY: test-race
test-race:
	@echo "Running tests with race detection..."
	$(GOTEST) -race -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: test-coverage-func
test-coverage-func:
	@echo "Running tests with function coverage..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -func=coverage.out

# Benchmark commands
.PHONY: bench
bench:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

.PHONY: bench-cpu
bench-cpu:
	@echo "Running CPU benchmarks..."
	$(GOTEST) -bench=. -benchmem -cpuprofile=cpu.prof ./...

.PHONY: bench-mem
bench-mem:
	@echo "Running memory benchmarks..."
	$(GOTEST) -bench=. -benchmem -memprofile=mem.prof ./...

.PHONY: bench-compare
bench-compare:
	@echo "Running benchmarks for comparison..."
	$(GOTEST) -bench=. -benchmem -count=5 ./... | tee bench.txt

# Code quality
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

.PHONY: vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

.PHONY: lint
lint:
	@echo "Running golint..."
	@which golint > /dev/null || (echo "Installing golint..." && $(GOGET) golang.org/x/lint/golint)
	golint ./...

.PHONY: check
check: fmt vet test
	@echo "All checks passed!"

# Module management
.PHONY: mod-tidy
mod-tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy

.PHONY: mod-verify
mod-verify:
	@echo "Verifying modules..."
	$(GOMOD) verify

.PHONY: mod-download
mod-download:
	@echo "Downloading modules..."
	$(GOMOD) download

# Build commands
.PHONY: build
build:
	@echo "Building..."
	$(GOBUILD) -v ./...

# Clean commands
.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f coverage.out coverage.html
	rm -f cpu.prof mem.prof
	rm -f bench.txt

.PHONY: clean-cache
clean-cache:
	@echo "Cleaning build cache..."
	$(GOCMD) clean -cache

# Development workflow
.PHONY: dev
dev: fmt vet test
	@echo "Development checks complete!"

.PHONY: ci
ci: mod-verify fmt vet test-race test-coverage
	@echo "CI pipeline complete!"

# Performance analysis
.PHONY: profile-cpu
profile-cpu: bench-cpu
	@echo "Analyzing CPU profile..."
	$(GOCMD) tool pprof cpu.prof

.PHONY: profile-mem
profile-mem: bench-mem
	@echo "Analyzing memory profile..."
	$(GOCMD) tool pprof mem.prof

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  test           - Run all tests"
	@echo "  test-short     - Run short tests"
	@echo "  test-race      - Run tests with race detection"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  bench          - Run benchmarks"
	@echo "  bench-compare  - Run benchmarks for comparison"
	@echo "  fmt            - Format code"
	@echo "  vet            - Run go vet"
	@echo "  lint           - Run golint"
	@echo "  check          - Run fmt, vet, and test"
	@echo "  dev            - Development workflow (fmt, vet, test)"
	@echo "  ci             - CI pipeline (verify, fmt, vet, race, coverage)"
	@echo "  build          - Build the package"
	@echo "  clean          - Clean build artifacts"
	@echo "  mod-tidy       - Tidy modules"
	@echo "  profile-cpu    - Run CPU profiling"
	@echo "  profile-mem    - Run memory profiling"
	@echo "  help           - Show this help message"
