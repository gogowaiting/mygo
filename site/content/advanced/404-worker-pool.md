---
title: "404. Worker Pool"
weight: 404
description: "Worker Pool 模式、Context 取消、方向性 Channel"
---

## 知识点

Worker Pool 是 Go 并发编程中最常用的模式之一。N 个 Worker 从共享的 jobs channel 消费任务，计算结果后发送到 results channel。通过 `context.WithTimeout` 控制整体截止时间，超时后所有 worker 优雅停止。

方向性 Channel 是类型安全的关键：worker 函数签名中 `jobs <-chan int` 表示只接收，`results chan<- int` 表示只发送，编译器会阻止误用。`select` 同时监听 `ctx.Done()`（超时信号）和 channel 接收，确保 worker 能及时响应取消。

发送端在所有 job 发送完毕后 `close(jobs)`，worker 检测到 `ok == false` 时退出。主 goroutine 通过 `sync.WaitGroup` 等待所有 worker 完成后再 `close(results)`，最后 range results 收集结果。

## 示例代码

{{< code "../cmd/404-worker-pool/main.go" >}}

## 运行方式

```bash
go run ./cmd/404-worker-pool/
```

## 源码位置

[`cmd/404-worker-pool/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/404-worker-pool/main.go)
