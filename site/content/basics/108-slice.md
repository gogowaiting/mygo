---
title: "108. 切片"
weight: 108
description: "初始化、截取、追加、底层数组"
---

## 知识点

切片是对底层数组的引用，其运行时结构包含三个字段：指向底层数组的指针、长度（len）和容量（cap）。切片的默认零值为 nil。使用 `make([]string, 3)` 可以初始化一个指定长度的空切片。

切片通过 `append` 追加元素，使用 `copy` 进行深拷贝，使用 `s[low:high]` 语法截取子切片。需要注意的是，子切片与原切片共享底层数组，修改子切片会影响原切片。当追加元素导致长度超过容量时，Go 会自动扩容底层数组（Go 1.18+ 策略：容量 < 256 时翻倍，>= 256 时增长 1.25 倍加 192）。

## 示例代码

{{< code "../cmd/108-slice/main.go" >}}

## 运行方式

```bash
go run ./cmd/108-slice/
```

## 源码位置

[`cmd/108-slice/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/108-slice/main.go)
