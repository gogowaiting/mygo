---
title: "407. 泛型 Set"
weight: 407
description: "Go 1.18 泛型、comparable 约束"
---

## 知识点

Go 1.18 引入泛型后，可以实现类型安全的通用数据结构。本示例实现了一个泛型 `Set[T]` 容器，底层使用 `map[T]struct{}` 存储。

`comparable` 约束限制了类型参数必须支持 `==` 和 `!=` 操作，这正是 map key 所要求的。`struct{}{}` 是零大小的值类型，用作 map 的 value 是 Go 中实现 Set 的惯用方式，不占用额外内存。

泛型构造函数 `NewSet[T comparable](items ...T)` 使用可变参数，支持 `NewSet(1, 2, 3)` 的简洁调用。方法定义在类型上 `(s Set[T])`，自动继承类型参数。同一个 Set 实现可以用于 `int`、`string` 等任意 comparable 类型，编译器会在编译期检查类型安全。

## 示例代码

{{< code "../cmd/407-generic-set/main.go" >}}

## 运行方式

```bash
go run ./cmd/407-generic-set/
```

## 源码位置

[`cmd/407-generic-set/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/407-generic-set/main.go)
