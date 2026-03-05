---
title: "301. 冒泡排序"
weight: 301
description: "O(n²) 稳定排序、相邻交换"
---

## 知识点

冒泡排序是最基础的排序算法之一，时间复杂度为 O(n^2)，是一种稳定的排序算法。其核心思想是反复遍历待排序数组，比较相邻的两个元素，如果顺序错误就交换它们。每一轮遍历都会将当前未排序部分的最大值"冒泡"到末尾。

外层循环控制排序轮次，共需要 `n-1` 轮。内层循环负责相邻元素的比较和交换，范围从 `0` 到 `n-i-1`，因为每轮结束后末尾的元素已经有序，无需再比较。

Go 语言支持元组交换，可以直接用 `arr[j], arr[j+1] = arr[j+1], arr[j]` 交换两个元素，无需临时变量。

## 示例代码

{{< code "../cmd/301-bubble-sort/main.go" >}}

## 运行方式

```bash
go run ./cmd/301-bubble-sort/
```

## 源码位置

[`cmd/301-bubble-sort/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/301-bubble-sort/main.go)
