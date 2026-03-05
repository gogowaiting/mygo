---
title: "116. 切片扩容"
weight: 116
description: "append 行为、容量增长、底层数组共享"
---

## 知识点

切片（slice）是 Go 中最常用的动态数组。`append` 是向切片追加元素的核心函数。当切片的长度等于容量（`len == cap`）时，`append` 会触发扩容，分配一个更大的底层数组，通常按约 2 倍增长（1, 2, 4, 8, 16, 32...）。

一个关键的陷阱是底层数组共享：当容量足够时，`append` 返回的新切片与原切片共享底层数组。修改新切片的元素会影响原切片。如果需要独立副本，应使用 `make` + `copy` 显式复制，避免意外的副作用。

`append` 的另一个常用场景是合并两个切片，使用 `slice...` 语法将第二个切片打散后追加。

## 示例代码

{{< code "../cmd/116-slice-append/main.go" >}}

## 运行方式

```bash
go run ./cmd/116-slice-append/
```

## 源码位置

[`cmd/116-slice-append/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/116-slice-append/main.go)
