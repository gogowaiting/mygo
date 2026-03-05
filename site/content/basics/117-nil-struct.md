---
title: "117. nil 与空结构体"
weight: 117
description: "nil 判定、零值语义、空结构体 Set"
---

## 知识点

空结构体 `struct{}` 不占用任何内存空间，是 Go 中实现"只关心存在性，不关心值"场景的理想选择。最常见的用途是配合 `map` 实现 Set（集合），因为值类型不分配内存，相比 `map[string]bool` 更加高效。

`struct{}` 的另一个重要用途是作为 channel 的信号类型。当 channel 只用于协调 goroutine 的执行顺序，不需要传递实际数据时，使用 `chan struct{}` 语义最清晰，表示"这是一个信号，不是数据"。

通过 `map` + `struct{}` 可以实现完整的集合操作：`Add` 添加元素、`Has` 判定存在、`Remove` 删除元素。这种模式在 Go 标准库和社区项目中被广泛使用。

## 示例代码

{{< code "../cmd/117-nil-struct/main.go" >}}

## 运行方式

```bash
go run ./cmd/117-nil-struct/
```

## 源码位置

[`cmd/117-nil-struct/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/117-nil-struct/main.go)
