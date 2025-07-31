# GenericMap

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-blue.svg)](https://golang.org/)
[![测试覆盖率](https://img.shields.io/badge/coverage-98.5%25-brightgreen.svg)](./MAKEFILE_CN.md)
[![许可证](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/costa92/genericmap)](https://goreportcard.com/report/github.com/costa92/genericmap)

高性能、线程安全的Go泛型双向映射实现，支持O(1)操作和全面的反向查找功能。

## ✨ 核心特性

- 🚀 **高性能**: O(1) 时间复杂度的设置、获取和删除操作
- 🔄 **双向查找**: 高效的正向（键→值）和反向（值→键集合）查找
- 🔒 **线程安全**: 内置RWMutex实现并发访问支持
- 🎯 **泛型支持**: 基于Go泛型的类型安全，支持任何可比较类型
- 📊 **内存优化**: 减少内存分配和GC压力
- 🧪 **充分测试**: 98.5%测试覆盖率，包含全面的基准测试

## 🚀 快速开始

### 安装

```bash
go get github.com/costa92/genericmap
```

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/costa92/genericmap"
)

func main() {
    // 创建新的映射
    m := genericmap.New[string, int]()
    
    // 添加键值对
    m.Set("苹果", 5)
    m.Set("香蕉", 2)
    m.Set("樱桃", 5)
    
    // 通过键获取值
    if value, exists := m.Get("苹果"); exists {
        fmt.Printf("苹果: %d个\n", value) // 苹果: 5个
    }
    
    // 反向查找 - 查找值为5的所有键
    keys := m.GetKeys(5)
    fmt.Printf("数量为5的水果: %v\n", keys) // [苹果 樱桃]
    
    // 删除项目
    removed := m.Remove("香蕉")
    fmt.Printf("已删除香蕉: %t\n", removed) // true
}
```

## 📖 API 文档

### 创建映射

```go
// 创建空映射
m := genericmap.New[string, int]()

// 使用初始数据创建
initial := map[string]int{"张三": 25, "李四": 30}
m := genericmap.New[string, int](initial)

// 指定容量创建（性能优化）
m := genericmap.NewWithCapacity[string, int](1000)
```

### 核心操作

```go
// 设置/更新
m.Set("用户ID", 12345)

// 获取
value, exists := m.Get("用户ID")

// 删除
removed := m.Remove("用户ID")

// 获取大小
size := m.Len()
```

### 双向查找

```go
// 获取特定值对应的所有键
keys := m.GetKeys(12345)

// 列出所有键
allKeys := m.List()

// 列出所有值
allValues := m.Values()
```

### 线程安全

```go
// 并发访问安全
var wg sync.WaitGroup

// 多个goroutine可以安全地访问映射
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        m.Set(fmt.Sprintf("用户-%d", id), id)
        value, _ := m.Get(fmt.Sprintf("用户-%d", id))
        keys := m.GetKeys(id)
        fmt.Printf("协程 %d: 值=%d, 键=%v\n", id, value, keys)
    }(i)
}

wg.Wait()
```

## 🏎️ 性能表现

### 基准测试结果 (Apple M1)

| 操作 | 时间/操作 | 分配次数/操作 | 字节/操作 |
|------|-----------|---------------|-----------|
| Set | 426.2 ns | 1 | 76 B |
| Get | 14.34 ns | 0 | 0 B |
| GetKeys | 771.9 ns | 2 | 904 B |
| Remove | 248.5 ns | 0 | 0 B |
| 并发读写 | 142.2 ns | 1 | 32 B |

### 性能优化

- **O(1) 操作**: 设置、获取和删除操作都是常数时间复杂度
- **优化反向查找**: 使用 `map[V]map[K]struct{}` 而非 `map[V][]K` 实现O(1)删除
- **减少分配**: 操作期间最小化内存分配
- **单次查找**: 消除Set操作中的冗余查找

## 📋 使用场景

### 用户-角色映射

```go
// 用户ID到角色名的映射，支持反向查找
userRoles := genericmap.New[int, string]()

userRoles.Set(1001, "管理员")
userRoles.Set(1002, "普通用户")
userRoles.Set(1003, "管理员")

// 查找所有管理员用户
adminUsers := userRoles.GetKeys("管理员") // [1001, 1003]
```

### 缓存系统

```go
// 带反向索引的缓存系统
cache := genericmap.New[string, []byte]()

cache.Set("用户:123", []byte("用户数据"))
cache.Set("文章:456", []byte("文章数据"))

// 根据缓存值查找键
keys := cache.GetKeys([]byte("用户数据")) // ["用户:123"]
```

### 配置管理

```go
// 配置项管理，支持按环境分组
config := genericmap.New[string, string]()

config.Set("数据库.主机", "生产环境")
config.Set("缓存.主机", "生产环境")
config.Set("队列.主机", "测试环境")

// 查找所有生产环境服务
prodServices := config.GetKeys("生产环境") // ["数据库.主机", "缓存.主机"]
```

### 商品分类系统

```go
// 商品ID到分类的映射
products := genericmap.New[string, string]()

products.Set("P001", "电子产品")
products.Set("P002", "服装")
products.Set("P003", "电子产品")
products.Set("P004", "书籍")

// 查找特定分类下的所有商品
electronics := products.GetKeys("电子产品") // ["P001", "P003"]
```

## 🛠️ 开发指南

### 环境要求

- Go 1.25 或更高版本
- Make（可选，用于使用Makefile命令）

### 构建和测试

```bash
# 运行所有测试
make test

# 运行基准测试
make bench

# 运行测试并生成覆盖率报告
make test-coverage

# 格式化和检查代码
make check

# 显示所有可用命令
make help
```

### 项目结构

```
├── map.go              # 核心实现
├── map_test.go         # 单元测试
├── example_test.go     # 示例测试
├── benchmark_test.go   # 性能基准测试
├── Makefile           # 构建自动化
├── MAKEFILE.md        # 英文Make命令文档
├── MAKEFILE_CN.md     # 中文Make命令文档
├── README.md          # 英文项目说明
└── README_CN.md       # 中文项目说明（本文件）
```

## 🧪 测试

项目通过全面的测试保持高代码质量：

- **单元测试**: 所有功能的完整覆盖
- **示例测试**: 有文档的使用示例
- **基准测试**: 性能回归测试
- **竞态测试**: 并发访问验证
- **覆盖率**: 98.5% 测试覆盖率

运行测试：

```bash
# 基础测试
go test -v

# 带竞态检测
go test -race -v

# 带覆盖率
go test -cover -v
```

## 🔧 高级用法

### 自定义类型

```go
type 用户ID int
type 部门名称 string

userMap := genericmap.New[用户ID, 部门名称]()
userMap.Set(用户ID(1001), 部门名称("技术部"))
```

### 性能调优

```go
// 预分配容量以获得更好性能
expectedSize := 10000
m := genericmap.NewWithCapacity[string, int](expectedSize)

// 批量操作
for i := 0; i < expectedSize; i++ {
    m.Set(fmt.Sprintf("键-%d", i), i)
}
```

### 错误处理

```go
// 操作前检查键是否存在
if value, exists := m.Get("键"); exists {
    // 处理现有值
    fmt.Printf("找到: %v\n", value)
} else {
    // 处理缺失的键
    fmt.Println("键不存在")
}

// 检查删除是否成功
if removed := m.Remove("键"); removed {
    fmt.Println("成功删除")
} else {
    fmt.Println("键不存在")
}
```

## 📊 与其他方案对比

| 特性 | GenericMap | sync.Map | 普通map + mutex |
|------|------------|----------|-----------------|
| 类型安全 | ✅ 泛型 | ❌ interface{} | ✅ 类型化 |
| 反向查找 | ✅ O(1) | ❌ 不支持 | ❌ O(n) 扫描 |
| 性能 | ⚡ 优化 | 🐌 接口开销 | 💾 内存高效 |
| 线程安全 | ✅ 内置 | ✅ 内置 | ⚙️ 手动 |
| API 简洁性 | ✅ 清晰 | ❌ 复杂 | ✅ 简单 |

## 🤝 贡献

欢迎贡献代码！请随时提交Pull Request。对于重大更改，请先开issue讨论您想要更改的内容。

### 开发工作流

1. Fork 仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 运行测试 (`make dev`)
4. 提交更改 (`git commit -am '添加amazing功能'`)
5. 推送到分支 (`git push origin feature/amazing-feature`)
6. 开启Pull Request

### 代码规范

- 遵循Go约定和最佳实践
- 运行 `make fmt` 格式化代码
- 确保 `make ci` 通过后再提交PR
- 保持测试覆盖率在95%以上

## 📄 许可证

本项目采用MIT许可证 - 详情请见 [LICENSE](LICENSE) 文件。

## 🙏 致谢

- Go团队提供的优秀泛型实现
- 社区反馈和贡献
- Go社区的性能优化技术

## 📚 相关文档

- [English README](./README.md)
- [英文 Make 命令文档](./MAKEFILE.md)
- [中文 Make 命令文档](./MAKEFILE_CN.md)
- [Go 包文档](https://pkg.go.dev/github.com/costa92/genericmap)

## 🔗 相关链接

- [GitHub 仓库](https://github.com/costa92/genericmap)
- [Go 包索引](https://pkg.go.dev/github.com/costa92/genericmap)
- [问题追踪](https://github.com/costa92/genericmap/issues)

## 🎯 最佳实践

### 性能优化建议

1. **预分配容量**: 如果知道预期大小，使用 `NewWithCapacity`
2. **批量操作**: 尽量减少锁的争用
3. **合理使用反向查找**: 在需要频繁反向查找的场景下使用

### 内存管理

1. **及时清理**: 删除不需要的键值对
2. **避免内存泄漏**: 注意大对象的引用
3. **监控性能**: 使用基准测试监控性能变化

### 并发使用

1. **读多写少**: 充分利用RWMutex的读写分离
2. **避免长时间持锁**: 快速完成操作后释放锁
3. **合理的并发粒度**: 避免过度并发导致锁争用

## 🔮 未来计划

- [ ] 支持更多的查询方式
- [ ] 添加序列化支持
- [ ] 提供更多的统计信息
- [ ] 支持自定义比较函数

---

<p align="center">
  <strong>GenericMap</strong> - 为Go提供高性能双向映射
</p>

<p align="center">
  如果这个项目对您有帮助，请给我们一个⭐️
</p>