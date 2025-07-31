# GenericMap - Make 命令文档

本文档提供 GenericMap 项目所有可用 Make 命令的详细指南。

## 快速开始

```bash
# 运行所有测试
make test

# 运行基准测试
make bench

# 显示所有可用命令
make help
```

## 测试命令

### `make test`
运行所有测试并显示详细输出。

```bash
make test
```

**示例输出：**
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
仅运行短测试，跳过较长的集成测试。

```bash
make test-short
```

### `make test-race`
使用 Go 的竞态检测器运行测试，识别数据竞争条件。

```bash
make test-race
```

**重要提示：** 这对于并发代码验证至关重要。

### `make test-coverage`
生成 HTML 格式的全面测试覆盖率报告。

```bash
make test-coverage
```

**输出文件：**
- `coverage.out` - 覆盖率数据文件
- `coverage.html` - 可视化 HTML 覆盖率报告

**当前覆盖率：98.5%**

### `make test-coverage-func`
在终端中显示函数级别的测试覆盖率。

```bash
make test-coverage-func
```

**示例输出：**
```
github.com/costa92/genericmap/map.go:28:    New                     100.0%
github.com/costa92/genericmap/map.go:60:    Set                     100.0%
github.com/costa92/genericmap/map.go:85:    Get                     100.0%
...
total:                                      (statements)            98.5%
```

## 基准测试命令

### `make bench`
运行所有基准测试并显示内存分配统计信息。

```bash
make bench
```

**示例输出：**
```
BenchmarkSet-8                   2915026    426.2 ns/op    76 B/op    1 allocs/op
BenchmarkGet-8                  82098064     14.34 ns/op    0 B/op    0 allocs/op
BenchmarkGetKeys-8               1608144    771.9 ns/op   904 B/op    2 allocs/op
BenchmarkRemove-8                5715165    248.5 ns/op     0 B/op    0 allocs/op
BenchmarkConcurrentReadWrite-8   8365099    142.2 ns/op    32 B/op    1 allocs/op
```

### `make bench-cpu`
运行基准测试并进行 CPU 性能分析。

```bash
make bench-cpu
```

**生成文件：** `cpu.prof` 用于性能分析

### `make bench-mem`
运行基准测试并进行内存性能分析。

```bash
make bench-mem
```

**生成文件：** `mem.prof` 用于内存分析

### `make bench-compare`
多次运行基准测试用于性能对比。

```bash
make bench-compare
```

**输出文件：** `bench.txt` 包含详细结果

## 代码质量命令

### `make fmt`
根据 Go 标准格式化所有 Go 代码。

```bash
make fmt
```

**功能：**
- 应用 `gofmt -s -w .`
- 简化代码结构
- 确保格式一致性

### `make vet`
运行 Go 的静态分析工具查找潜在问题。

```bash
make vet
```

**检测内容：**
- 可疑的代码结构
- 潜在的错误
- 不可移植的代码

### `make lint`
运行 golint 进行代码风格检查。

```bash
make lint
```

**注意：** 如果 golint 不存在会自动安装。

### `make check`
综合代码质量检查，结合格式化、静态分析和测试。

```bash
make check
```

**等效于：**
```bash
make fmt
make vet  
make test
```

## 开发工作流命令

### `make dev`
日常编码的标准开发工作流。

```bash
make dev
```

**执行步骤：**
1. 代码格式化 (`fmt`)
2. 静态分析 (`vet`)
3. 运行所有测试 (`test`)

### `make ci`
完整的 CI/CD 流水线验证。

```bash
make ci
```

**执行步骤：**
1. 模块验证 (`mod-verify`)
2. 代码格式化 (`fmt`)
3. 静态分析 (`vet`)
4. 竞态条件测试 (`test-race`)
5. 覆盖率分析 (`test-coverage`)

## 模块管理

### `make mod-tidy`
清理和整理 Go 模块依赖项。

```bash
make mod-tidy
```

### `make mod-verify`
验证依赖项未被修改。

```bash
make mod-verify
```

### `make mod-download`
下载所有必需的模块。

```bash
make mod-download
```

## 构建命令

### `make build`
构建包以验证编译。

```bash
make build
```

## 性能分析

### `make profile-cpu`
使用生成的性能数据分析 CPU 性能。

```bash
make profile-cpu
```

**前提条件：** 首先运行 `make bench-cpu`

### `make profile-mem`
分析内存使用模式。

```bash
make profile-mem
```

**前提条件：** 首先运行 `make bench-mem`

## 清理命令

### `make clean`
移除所有生成的文件和构建产物。

```bash
make clean
```

**移除文件：**
- `coverage.out`
- `coverage.html`
- `cpu.prof`
- `mem.prof`
- `bench.txt`

### `make clean-cache`
清除 Go 的构建缓存。

```bash
make clean-cache
```

## 实用命令

### `make help`
显示所有可用的 Make 目标及其描述。

```bash
make help
```

## 常用工作流程

### 提交代码前
```bash
make dev
```

### 性能测试
```bash
make bench
make bench-compare
```

### 完整质量检查
```bash
make ci
```

### 调试性能问题
```bash
make bench-cpu
make profile-cpu
```

### 覆盖率分析
```bash
make test-coverage
# 在浏览器中打开 coverage.html
```

## 性能指标

在 Apple M1 上的当前基准测试结果：

| 操作 | 时间/操作 | 分配次数/操作 | 字节/操作 |
|------|-----------|---------------|-----------|
| Set | 426.2 ns | 1 | 76 B |
| Get | 14.34 ns | 0 | 0 B |
| GetKeys | 771.9 ns | 2 | 904 B |
| Remove | 248.5 ns | 0 | 0 B |
| 并发读写 | 142.2 ns | 1 | 32 B |

## 测试覆盖率

- **总体覆盖率：** 98.5%
- **所有函数：** 90%+ 覆盖率
- **关键路径：** 100% 覆盖率

## 使用技巧

1. **日常开发：** 使用 `make dev` 进行快速验证
2. **PR 前：** 始终运行 `make ci`
3. **性能调优：** 使用 `make bench-compare` 测量改进效果
4. **调试：** 使用 `make test-race` 捕获并发问题
5. **覆盖率目标：** 保持 >95% 的测试覆盖率

## 常见问题解答

### Q: 如何快速验证代码更改？
A: 运行 `make dev`，它会执行格式化、静态分析和测试。

### Q: 如何检查性能回归？
A: 使用 `make bench-compare` 在代码更改前后运行基准测试。

### Q: 如何生成覆盖率报告？
A: 运行 `make test-coverage`，然后在浏览器中打开生成的 `coverage.html`。

### Q: 如何分析内存泄漏？
A: 运行 `make bench-mem` 然后 `make profile-mem` 进行内存分析。

### Q: CI 流水线应该运行什么命令？
A: 运行 `make ci`，它包含了所有必要的验证步骤。

## 最佳实践

1. **开发前：** 运行 `make clean` 清理环境
2. **开发中：** 频繁运行 `make dev` 验证更改
3. **提交前：** 运行 `make ci` 确保质量
4. **发布前：** 运行完整的基准测试和性能分析
5. **维护：** 定期运行 `make mod-tidy` 清理依赖

## 自定义扩展

如果需要添加新的 Make 目标，请在 `Makefile` 中添加，并更新此文档。遵循现有的命名约定和结构模式。