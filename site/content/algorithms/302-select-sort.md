---
title: "302. 选择排序"
weight: 302
description: "O(n²) 原地排序、最小值选择"
---

## 知识点

选择排序的时间复杂度为 O(n^2)，是一种不稳定的原地排序算法。其核心思想是每一轮从未排序部分中找到最小元素的索引，然后将其与未排序部分的第一个元素交换位置。

外层循环 `i` 遍历每个位置，内层循环从 `i` 开始向后扫描，用变量 `min` 记录当前最小值的索引。内层循环结束后，`arr[i]` 和 `arr[min]` 交换。与冒泡排序不同，选择排序每轮只做一次交换，因此在数据移动成本较高时更有优势。

由于选择排序的交换次数少（最多 `n-1` 次），在某些对写操作敏感的场景（如闪存存储）中比冒泡排序更实用。

## 示例代码

{{< code "../cmd/302-select-sort/main.go" >}}

## 运行方式

```bash
go run ./cmd/302-select-sort/
```

## 源码位置

[`cmd/302-select-sort/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/302-select-sort/main.go)
