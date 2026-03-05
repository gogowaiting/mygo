---
title: "406. Atomic"
weight: 406
description: "无锁并发计数器、sync/atomic"
---

## 知识点

`sync/atomic` 包提供了原子操作，是实现无锁并发的基础。相比 `sync.Mutex`，原子操作直接利用 CPU 指令（如 CAS），性能更高。

本示例启动 100 个 goroutine，每个递增 1000 次，最终结果始终为 100000。`atomic.AddInt64(&counter, 1)` 是原子递增操作，保证多个 goroutine 同时修改不会出现数据竞争。`atomic.LoadInt64(&counter)` 是原子读取，确保读到最新值。

`sync.WaitGroup` 用于等待所有 goroutine 完成：`wg.Add(1)` 在启动前调用，`defer wg.Done()` 在 goroutine 结束时递减计数器，`wg.Wait()` 阻塞直到计数器归零。这种组合是 Go 并发编程的标准模式。

## 示例代码

{{< code "../cmd/406-atomic/main.go" >}}

## 运行方式

```bash
go run ./cmd/406-atomic/
```

## 源码位置

[`cmd/406-atomic/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/406-atomic/main.go)
