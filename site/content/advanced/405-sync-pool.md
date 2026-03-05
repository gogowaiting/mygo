---
title: "405. sync.Pool"
weight: 405
description: "对象复用、减少 GC 压力"
---

## 知识点

`sync.Pool` 是 Go 标准库提供的对象池，用于复用临时对象以减少 GC 压力。在高并发场景下频繁创建和销毁对象（如 `bytes.Buffer`）会产生大量垃圾，Pool 让对象在使用后归还，下次直接复用。

`New` 字段是一个工厂函数，当池中无可用对象时调用创建新实例。`pool.Get()` 获取对象（类型为 `interface{}`），需要类型断言转换为具体类型。使用完毕后 `pool.Put(buf)` 归还到池中。使用前必须 `buf.Reset()` 清除上次残留数据。

注意：Pool 中的对象可能在任意 GC 周期被回收，因此不适合存放需要长期持有的对象。它最适合生命周期短、创建开销大的临时对象。

## 示例代码

{{< code "../cmd/405-sync-pool/main.go" >}}

## 运行方式

```bash
go run ./cmd/405-sync-pool/
```

## 源码位置

[`cmd/405-sync-pool/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/405-sync-pool/main.go)
