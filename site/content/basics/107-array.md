---
title: "107. 数组"
weight: 107
description: "声明、初始化、多维数组、值类型"
---

## 知识点

Go 数组是固定长度的同类型元素序列，长度是类型的一部分。声明方式有多种：`var a [5]int` 声明后索引赋值、`[5]int{1,2,3,4,5}` 字面量初始化、`[...]int{10,20,30}` 编译器自动推导长度、以及按索引初始化如 `[5]int{1: 100, 3: 300}`。

数组是值类型，赋值和传参时会进行拷贝，修改拷贝不影响原数组。这是与切片的重要区别。多维数组通过嵌套声明实现，遍历数组推荐使用 `range`。

## 示例代码

{{< code "../cmd/107-array/main.go" >}}

## 运行方式

```bash
go run ./cmd/107-array/
```

## 源码位置

[`cmd/107-array/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/107-array/main.go)
