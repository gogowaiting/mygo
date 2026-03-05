# 项目架构

## 设计目标

- 按知识点拆分，每个 demo 独立可运行
- 数字编号标识类别和顺序
- 学习 CLI 统一入口管理所有 demo
- 工程示例与学习 demo 分离

## 目录分层

```
mygo/
├── cmd/              # 学习 demo（按编号分组）
│   ├── 1xx           # 基础语法
│   ├── 2xx           # 并发编程
│   ├── 3xx           # 算法
│   ├── 4xx           # 进阶
│   ├── learn/        # 学习导航 CLI
│   └── interview/    # 面试训练 CLI
│
├── examples/         # 工程实践示例
├── patterns/         # 设计模式
├── internal/         # 内部元数据
└── docs/             # 文档
```

## 编号规则

| 前缀 | 类别 | 示例 |
|------|------|------|
| 1xx | 基础语法 | 101-pre, 102-hello, ..., 117-nil-struct |
| 2xx | 并发编程 | 201-channel, 202-goroutine, ..., 204-timer |
| 3xx | 算法 | 301-bubble-sort, ..., 304-quick-sort |
| 4xx | 进阶 | 401-guess-number, ..., 407-generic-set |

## CLI 工具

### learn CLI (`cmd/learn/`)

学习目录元数据存储在 `internal/learn/catalog.go`，CLI 读取元数据并：
- `list`：列出全部 demo
- `tree`：按主题分组展示
- `topic <name>`：查看特定主题
- `run <id>`：运行指定 demo

### interview CLI (`cmd/interview/`)

面试题库存储在 `internal/interview/questions.go`，CLI 提供：
- `list`：列出全部题目
- `topics`：按主题分组
- `random`：随机抽题
- `show <id>`：查看具体题目

## 质量门禁

```bash
make ci   # fmt + vet + test + build
```

- `make fmt`：代码格式化
- `make vet`：静态分析
- `make test`：运行所有测试
- `make build`：编译检查
