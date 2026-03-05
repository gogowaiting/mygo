---
title: "110. Range"
weight: 110
description: "range 遍历 slice、map、string"
---

## 知识点

`range` 可以遍历各种数据结构中的元素。遍历切片时返回索引和值，遍历 map 时返回键和值，遍历字符串时返回字节索引和 rune 值。可以使用 `_` 忽略不需要的返回值。

一个重要的陷阱是：`range` 返回的值是原始数据的拷贝，直接修改 range 返回的值不会影响原始数据。要修改原始数据，应该通过索引操作，如 `accounts[i].price += 10`。

## 示例代码

{{< code "../cmd/110-range/main.go" >}}

## 运行方式

```bash
go run ./cmd/110-range/
```

## 源码位置

[`cmd/110-range/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/110-range/main.go)
