---
title: "106. Switch"
weight: 106
description: "值 switch、类型 switch、无表达式 switch"
---

## 知识点

Go 的 `switch` 语句不需要在每个 case 后加 `break`，默认不会贯穿（fall through）到下一个 case。同一个 case 可以用逗号分隔多个表达式。

无表达式的 switch 是一种 if/else 的替代写法，case 中可以使用任意的布尔条件表达式，而不只是常量值。类型 switch（type switch）使用 `i.(type)` 语法来判断接口变量的实际类型，这在处理多态数据时非常有用。

## 示例代码

{{< code "../cmd/106-switch/main.go" >}}

## 运行方式

```bash
go run ./cmd/106-switch/
```

## 源码位置

[`cmd/106-switch/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/106-switch/main.go)
