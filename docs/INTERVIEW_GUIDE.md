# Go 面试模块

## 使用方式

```bash
# 列出全部题目
go run ./cmd/interview list

# 按主题查看
go run ./cmd/interview topics

# 随机抽题
go run ./cmd/interview random

# 查看具体题目
go run ./cmd/interview show go-gmp
go run ./cmd/interview show go-channel-close
go run ./cmd/interview show go-gc
```

## 题库覆盖

| 主题 | 题目数 | 典型题目 |
|------|--------|----------|
| runtime | 2 | GMP 模型、GC 机制 |
| concurrency | 2 | Channel 关闭语义、Context |
| type-system | 1 | interface nil 比较 |
| memory | 1 | 逃逸分析 |
| language | 1 | defer 执行时机 |

## 复盘模板

每次答完题后，按以下模板记录：

- **题目**：
- **一句话答案**：
- **深入点 1（原理）**：
- **深入点 2（工程实践）**：
- **常见坑**：
- **相关代码位置**：

## 建议节奏

- 每天 3 题：1 道基础 + 1 道并发 + 1 道 runtime
- 每周 1 次模拟：随机 15 题，限定 30 分钟
- 答不出的题目加入复盘清单，隔天重做
