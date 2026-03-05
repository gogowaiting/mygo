---
title: "403. 流量控制"
weight: 403
description: "加权流量分配、原子计数、无锁方案"
---

## 知识点

流量控制（Traffic Control）用于按比例分配请求到不同服务，例如 A/B 测试中的 50/50 分流。本示例实现了一个无锁的加权流量分配器。

核心思路是：创建一个长度为 `base` 的数组（如比例 3:7 则 base=10），填入 `0,1,2,...,base-1`，然后用 `rand.Shuffle` 打乱顺序。每次请求通过 `atomic.AddUint32` 原子递增全局计数器，对 base 取模后查表，判断结果落在哪个区间即可决定路由。这种方法避免了加锁，性能极高。

`atomic.AddUint32` 保证了并发安全，无需 mutex。打乱数组使得分配在统计上均匀，不会出现连续命中同一服务的情况。

## 示例代码

{{< code "../cmd/403-traffic-control/main.go" >}}

## 运行方式

```bash
go run ./cmd/403-traffic-control/
```

## 源码位置

[`cmd/403-traffic-control/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/403-traffic-control/main.go)
