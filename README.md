# mygo

一个按知识点拆分、可运行、可测试的 Go 学习项目。涵盖基础语法、并发编程、经典算法、进阶模式、工程实践和面试题库。

## 项目结构

```
mygo/
├── cmd/                        # 可运行的独立 demo
│   ├── learn/                  # 学习导航 CLI
│   ├── interview/              # 面试训练 CLI
│   │
│   │   ── 基础语法 (1xx) ──
│   ├── 101-pre/                # 预定义标识符
│   ├── 102-hello/              # Hello World
│   ├── 103-var-const/          # 变量与常量
│   ├── 104-for/                # 循环
│   ├── 105-if/                 # 条件分支
│   ├── 106-switch/             # Switch
│   ├── 107-array/              # 数组
│   ├── 108-slice/              # 切片
│   ├── 109-map/                # Map
│   ├── 110-range/              # Range
│   ├── 111-func/               # 函数
│   ├── 112-pointer/            # 指针
│   ├── 113-struct/             # 结构体
│   ├── 114-method/             # 方法
│   ├── 115-interface/          # 接口
│   ├── 116-slice-append/       # 切片扩容
│   ├── 117-nil-struct/         # nil 与空结构体
│   │
│   │   ── 并发编程 (2xx) ──
│   ├── 201-channel/            # Channel
│   ├── 202-goroutine/          # Goroutine
│   ├── 203-range-chan/         # Range Channel
│   ├── 204-timer/              # 定时与超时
│   │
│   │   ── 算法 (3xx) ──
│   ├── 301-bubble-sort/        # 冒泡排序
│   ├── 302-select-sort/        # 选择排序
│   ├── 303-insert-sort/        # 插入排序
│   ├── 304-quick-sort/         # 快速排序
│   │
│   │   ── 进阶 (4xx) ──
│   ├── 401-guess-number/       # 猜数字（综合练习）
│   ├── 402-context/            # Context 取消与超时
│   ├── 403-traffic-control/    # 流量控制
│   ├── 404-worker-pool/        # Worker Pool
│   ├── 405-sync-pool/          # sync.Pool
│   ├── 406-atomic/             # 原子操作
│   └── 407-generic-set/        # 泛型 Set
│
├── examples/                   # 工程实践示例
│   ├── gin-demo/               # Gin + pprof
│   ├── rpc/                    # HTTP handler 与单测
│   ├── taillog/                # WebSocket 日志 tail
│   ├── k8s/                    # K8s Informer
│   ├── docker/                 # Docker SDK
│   ├── pprof/                  # 性能分析
│   └── benchmark-inline/       # go:noinline 基准测试
│
├── patterns/                   # 设计模式与并发实践
│   ├── singleton/              # 单例模式
│   ├── goroutine/              # Goroutine 模式
│   ├── waitgroup/              # WaitGroup
│   ├── interface/              # 接口实践
│   └── repeat/                 # 重复模式
│
├── internal/                   # 内部元数据（不对外暴露）
│   ├── learn/                  # 学习目录 catalog
│   └── interview/              # 面试题库
│
└── docs/                       # 文档
    ├── ARCHITECTURE.md
    └── INTERVIEW_GUIDE.md
```

## 快速开始

### 学习入口 CLI

```bash
# 列出全部知识点
go run ./cmd/learn list

# 按模块查看学习树
go run ./cmd/learn tree

# 查看某一类知识点
go run ./cmd/learn topic advanced

# 运行指定 demo
go run ./cmd/learn run basic-01-hello
go run ./cmd/learn run algo-quick
go run ./cmd/learn run adv-worker-pool
```

### 面试训练 CLI

```bash
go run ./cmd/interview list      # 列出全部题目
go run ./cmd/interview topics    # 按主题分组
go run ./cmd/interview random    # 随机抽题
go run ./cmd/interview show go-gmp  # 查看具体题目
```

### 直接运行 demo

```bash
go run ./cmd/102-hello
go run ./cmd/201-channel
go run ./cmd/304-quick-sort
```

### 工程示例

```bash
go run ./examples/gin-demo
go run ./examples/rpc
go run ./examples/taillog
go run ./examples/docker
go run ./examples/k8s
```

## 学习路径

| 阶段 | 主题 | 内容 | 建议顺序 |
|------|------|------|----------|
| 1 | 基础语法 | 变量、控制流、数据结构、函数、指针、结构体、接口 | 101-117 |
| 2 | 并发编程 | Goroutine、Channel、Range Channel、定时器 | 201-204 |
| 3 | 算法 | 冒泡、选择、插入、快速排序 | 301-304 |
| 4 | 进阶 | Context、流量控制、Worker Pool、sync.Pool、Atomic、泛型 | 401-407 |
| 5 | 面试 | 题库训练（GMP、GC、channel、interface 等） | interview CLI |
| 6 | 工程实践 | HTTP、WebSocket、K8s、Docker、pprof | examples/ |

## 工程质量

```bash
make fmt      # 格式化
make vet      # 静态检查
make test     # 运行测试
make build    # 编译
make ci       # 全部检查（fmt + vet + test + build）
```

## 技术栈

- Go 1.18+
- Gin（HTTP 框架）
- gorilla/websocket（WebSocket）
- Docker SDK（容器操作）
- client-go（Kubernetes）

## License

[MIT](LICENSE)
