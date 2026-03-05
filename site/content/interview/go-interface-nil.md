---
title: "接口与 nil"
weight: 15
description: "非 nil 接口不等于 nil 的陷阱"
---

## 问题

为什么一个持有 nil 指针的接口不等于 nil？

## 答案

Go 的接口值（interface value）在运行时由两部分组成：类型信息（type）和数据指针（value）。接口变量只有在 type 和 value 都为 nil 时，才等于 nil。这是 Go 最著名的陷阱之一。

典型场景：`var p *MyError = nil`，然后 `var err error = p`，此时 `err != nil`。因为 err 的 type 是 `*MyError`（非 nil），value 是 nil。接口的 nil 判断是 `type == nil && value == nil`，不是只看 value。

这个问题在错误处理中特别常见。函数返回 `error` 接口时，如果返回了一个 nil 的具体错误指针，调用方 `if err != nil` 判断会为 true，导致逻辑错误。解决方法：在返回接口前显式检查底层值是否为 nil，若是则直接返回 `nil`（不带类型信息）。

可以用 `reflect` 包查看接口的内部结构：`reflect.ValueOf(err).IsNil()` 可以检查底层值，但更好的做法是在编码时避免这种情况。Go 1.18+ 引入了 `any` 作为 `interface{}` 的别名，但底层行为不变。

## 关键词

接口内部结构, type-value pair, nil trap, 类型断言, reflect
