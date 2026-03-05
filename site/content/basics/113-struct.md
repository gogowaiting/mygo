---
title: "113. 结构体"
weight: 113
description: "声明、嵌套、空结构体、泛型 Set"
---

## 知识点

结构体（struct）是 Go 中将多个不同类型字段聚合在一起的复合类型。可以通过字面量直接初始化，也可以指明字段名初始化，未指明的字段会使用零值。使用 `&` 可以获取结构体指针，通过指针可以直接修改字段值。

Go 支持结构体嵌套，将一个结构体嵌入另一个结构体中，内嵌结构体的字段和方法会被提升到外层，形成类似"继承"的效果。这种组合优于继承的设计是 Go 的核心理念之一。

空结构体 `struct{}` 不占用任何内存空间，常用于实现 Set（集合）数据结构。通过 `map[K]struct{}` 可以高效地实现集合操作，因为值类型不分配内存，只做占位用途。Go 1.18 之后支持泛型，可以定义类型安全的泛型 Set。

## 示例代码

{{< code "../cmd/113-struct/main.go" >}}

## 运行方式

```bash
go run ./cmd/113-struct/
```

## 源码位置

[`cmd/113-struct/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/113-struct/main.go)
