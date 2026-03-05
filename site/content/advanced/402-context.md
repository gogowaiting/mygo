---
title: "402. Context"
weight: 402
description: "WithCancel、WithTimeout、WithDeadline、WithValue"
---

## 知识点

`context.Context` 是 Go 并发编程中控制 goroutine 生命周期的核心机制，提供四大用法。

**WithCancel** 用于手动取消：调用 `cancel()` 后，所有监听 `ctx.Done()` 的 goroutine 都会收到信号。**WithTimeout** 用于超时自动取消：设置一个时间限制，到期后自动触发取消，务必 `defer cancel()` 释放资源。**WithDeadline** 类似 WithTimeout，但指定的是绝对时间点。**WithValue** 用于在调用链中传递请求级元数据（如请求 ID），应使用自定义类型作为 key 避免冲突。

实战中 Context 常用于"竞争模式"：启动多个 goroutine 并发请求，第一个完成的取消其余。这在微服务中非常常见，可以有效降低延迟。

## 示例代码

{{< code "../cmd/402-context/main.go" >}}

## 运行方式

```bash
go run ./cmd/402-context/
```

## 源码位置

[`cmd/402-context/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/402-context/main.go)
