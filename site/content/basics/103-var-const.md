---
title: "103. 变量与常量"
weight: 103
description: "var、:= 短声明、const、零值"
---

## 知识点

Go 中变量声明有两种主要方式：使用 `var` 关键字进行显式声明（如 `var a = "testString"`），或使用 `:=` 短声明进行类型推断（如 `f := float32(0)`）。可以同时声明多个变量，例如 `var b, c = 1, 2`。未显式初始化的变量会被赋予零值（数值类型为 0，bool 为 false，字符串为空字符串）。

常量使用 `const` 关键字定义，支持字符、字符串、布尔值和数值类型。常量可以在编译期进行表达式计算，例如 `const i = 3e20 / h` 会在编译时完成除法运算，这在运行时不会产生额外开销。

## 示例代码

{{< code "../cmd/103-var-const/main.go" >}}

## 运行方式

```bash
go run ./cmd/103-var-const/
```

## 源码位置

[`cmd/103-var-const/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/103-var-const/main.go)
