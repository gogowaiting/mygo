---
title: "Channel 关闭行为"
weight: 13
description: "关闭后读写的行为"
---

## 问题

Channel 关闭后读写分别会怎样？

## 答案

Channel 关闭后的行为需要从读和写两个方向分别理解。读取已关闭的 channel 不会阻塞也不会 panic，而是立即返回零值。使用 `v, ok := <-ch` 形式时，ok 为 false 表示 channel 已关闭且无数据。`for range ch` 循环在 channel 关闭后自动退出。

向已关闭的 channel 发送数据会触发 panic（`send on closed channel`）。重复关闭一个已经关闭的 channel 也会 panic（`close of closed channel`）。这两个 panic 无法通过 `recover()` 捕获，属于编程错误。

最佳实践是"发送方负责关闭"。因为发送方知道何时不再发送数据，而接收方无法确定是否还有数据到来。常见的关闭模式：使用 `sync.WaitGroup` 等待所有发送方完成后，在主 goroutine 中关闭 channel；或使用 `context.Done()` 信号配合 select 优雅退出。

实际代码中，如果多个 goroutine 都可能向同一个 channel 发送数据，需要额外的同步机制（如计数器或 WaitGroup）来确定何时关闭，否则容易出现向已关闭 channel 发送的 panic。

## 关键词

close, panic, 零值, ok 模式, 发送方关闭, for range, recover
