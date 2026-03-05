---
title: "102. Hello World"
weight: 102
description: "fmt 包、格式化输出"
---

## 知识点

最简单的 Go 程序由 `package main` 和 `func main()` 组成。`fmt` 包提供了基本的输入输出功能：`fmt.Println` 用于基本输出，`fmt.Printf` 用于格式化输出，`fmt.Sprintf` 用于构建字符串而不直接打印。

`fmt.Printf` 使用占位符来格式化输出，常用的占位符包括 `%s`（字符串）、`%d`（整数）、`%f`（浮点数）、`%.2f`（保留两位小数的浮点数）等。`fmt.Sprintf` 的格式与 `Printf` 相同，但返回字符串而非打印。

## 示例代码

{{< code "../cmd/102-hello/main.go" >}}

## 运行方式

```bash
go run ./cmd/102-hello/
```

## 源码位置

[`cmd/102-hello/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/102-hello/main.go)
