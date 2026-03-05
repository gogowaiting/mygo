---
title: "112. 指针"
weight: 112
description: "值传递与指针传递"
---

## 知识点

Go 中的函数参数传递都是值传递（pass by value），即函数接收的是参数的副本。如果传入的是基本类型（如 `int`），函数内部对参数的修改不会影响原始值。要让函数修改外部变量，需要使用指针。

指针存储的是变量的内存地址。通过 `&` 取地址操作符获取变量地址，通过 `*` 解引用操作符访问指针指向的值。在函数参数中使用 `*int` 类型，函数内部通过 `*iptr = 0` 就可以直接修改原始变量。

值传递和指针传递的区别是理解 Go 内存模型的基础。在需要修改外部变量或避免大结构体拷贝开销时，应使用指针传递。

## 示例代码

{{< code "../cmd/112-pointer/main.go" >}}

## 运行方式

```bash
go run ./cmd/112-pointer/
```

## 源码位置

[`cmd/112-pointer/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/112-pointer/main.go)
