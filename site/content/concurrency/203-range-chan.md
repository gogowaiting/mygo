---
title: "203. Range Channel"
weight: 203
description: "生产者-消费者、WaitGroup、扇出模式"
---

## 知识点

`for-range` 是消费 channel 最常用的方式。`range` 会持续接收 channel 中的元素，直到 channel 被关闭。如果 channel 未被关闭，range 会永久阻塞导致死锁。因此生产者在发送完所有数据后必须调用 `close(ch)`。

`sync.WaitGroup` 用于协调多个 goroutine 的完成。调用 `wg.Add(n)` 设置计数器，每个 goroutine 完成时调用 `wg.Done()` 减一，主 goroutine 通过 `wg.Wait()` 阻塞等待所有任务完成。

扇出（Fan-out）模式是一种常见的并发模式：一个生产者通过缓冲 channel 分发任务，多个消费者（worker）从同一 channel 读取并处理任务。channel 关闭后所有 worker 的 range 循环都会退出。

## 示例代码

{{< code "../cmd/203-range-chan/main.go" >}}

## 运行方式

```bash
go run ./cmd/203-range-chan/
```

## 源码位置

[`cmd/203-range-chan/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/203-range-chan/main.go)
