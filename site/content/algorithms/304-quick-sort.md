---
title: "304. 快速排序"
weight: 304
description: "分治法、三路分区、递归"
---

## 知识点

快速排序是分治法的经典应用，平均时间复杂度为 O(n log n)，是实际应用中最常用的排序算法之一。其核心思路是选取一个基准值（pivot），将数组分割成三部分：小于 pivot 的元素、等于 pivot 的元素、大于 pivot 的元素，然后对左右两部分递归排序，最后拼接结果。

本实现采用三路分区法，正确处理了包含重复元素的情况。选取第一个元素作为 pivot，遍历剩余元素分别放入 `low`、`mid`、`hight` 三个切片。递归排序 `low` 和 `hight` 后，用 `append` 拼接得到最终结果。

三路分区法的优势在于，当数组中有大量重复元素时，等于 pivot 的元素不会被递归处理，显著提升了效率。这种写法虽然额外使用了空间，但逻辑清晰，易于理解。

## 示例代码

{{< code "../cmd/304-quick-sort/main.go" >}}

## 运行方式

```bash
go run ./cmd/304-quick-sort/
```

## 源码位置

[`cmd/304-quick-sort/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/304-quick-sort/main.go)
