# Go 学习路径

## 1. 基础语法（1xx）

| 编号 | Demo ID | 知识点 | 级别 |
|------|---------|--------|------|
| 101 | basic-00-pre | 预定义标识符 | L1 |
| 102 | basic-01-hello | Hello World | L1 |
| 103 | basic-02-var-const | 变量与常量 | L1 |
| 104 | basic-03-for | 循环 | L1 |
| 105 | basic-04-if | 条件分支 | L1 |
| 106 | basic-05-switch | Switch | L1 |
| 107 | basic-06-array | 数组 | L1 |
| 108 | basic-07-slice | 切片 | L1 |
| 109 | basic-08-map | Map | L1 |
| 110 | basic-09-range | Range | L1 |
| 111 | basic-10-func | 函数 | L1 |
| 112 | basic-11-point | 指针 | L2 |
| 113 | basic-12-struct | 结构体 | L2 |
| 114 | basic-13-method | 方法 | L2 |
| 115 | basic-14-interface | 接口 | L2 |
| 116 | basic-slice-append | 切片扩容 | L2 |
| 117 | basic-nil-struct | nil 与空结构体 | L2 |

## 2. 并发编程（2xx）

| 编号 | Demo ID | 知识点 | 级别 |
|------|---------|--------|------|
| 201 | basic-15-channel | Channel | L2 |
| 202 | basic-16-goroutine | Goroutine | L2 |
| 203 | basic-range-chan | Range Channel | L2 |
| 204 | basic-test-timer | 定时与超时 | L3 |

## 3. 算法（3xx）

| 编号 | Demo ID | 知识点 | 级别 |
|------|---------|--------|------|
| 301 | algo-bubble | 冒泡排序 | L1 |
| 302 | algo-select | 选择排序 | L1 |
| 303 | algo-insert | 插入排序 | L1 |
| 304 | algo-quick | 快速排序 | L2 |

## 4. 进阶（4xx）

| 编号 | Demo ID | 知识点 | 级别 |
|------|---------|--------|------|
| 401 | adv-guess-number | 猜数字 | L2 |
| 402 | adv-context | Context | L3 |
| 403 | adv-traffic-control | 流量控制 | L3 |
| 404 | adv-worker-pool | Worker Pool | L3 |
| 405 | adv-sync-pool | sync.Pool | L3 |
| 406 | adv-atomic | Atomic | L3 |
| 407 | adv-generic-set | Generic Set | L3 |

## 5. 面试模块

```bash
go run ./cmd/interview list      # 列出全部
go run ./cmd/interview random    # 随机抽题
go run ./cmd/interview show go-gmp  # 按题深入
```

## 6. 工程实践（examples/）

| 模块 | 路径 | 说明 |
|------|------|------|
| Gin + pprof | examples/gin-demo | Web 框架与性能分析 |
| HTTP RPC | examples/rpc | HTTP handler 与单测 |
| WebSocket | examples/taillog | 实时日志 tail |
| Kubernetes | examples/k8s | Informer 与事件监听 |
| Docker | examples/docker | Docker SDK 调用 |
| pprof | examples/pprof | 性能分析入门 |
| Benchmark | examples/benchmark-inline | go:noinline 基准测试 |

## 7. 设计模式（patterns/）

| 模块 | 路径 | 说明 |
|------|------|------|
| 单例 | patterns/singleton | 单例模式与测试 |
| Goroutine | patterns/goroutine | 并发模式 |
| WaitGroup | patterns/waitgroup | 同步等待 |
| 接口 | patterns/interface | 接口实践 |
| 重复 | patterns/repeat | 重复执行模式 |

## 运行命令

```bash
# 学习 CLI
go run ./cmd/learn list
go run ./cmd/learn tree
go run ./cmd/learn topic advanced
go run ./cmd/learn run basic-01-hello

# 面试 CLI
go run ./cmd/interview random

# 直接运行
go run ./cmd/102-hello
go run ./examples/gin-demo

# 工程质量
make ci
```
