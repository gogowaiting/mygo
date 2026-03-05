---
title: "201. Channel"
weight: 201
description: "无缓冲、有缓冲、for-range 消费"
---

## 知识点

Channel 是 Go 中 goroutine 之间通信的核心机制。通过 `make(chan T)` 创建无缓冲 channel，发送和接收会同步阻塞，直到另一方准备好。通过 `make(chan T, n)` 创建有缓冲 channel，缓冲未满时发送不会阻塞，缓冲为空时接收才会阻塞。

除了直接发送接收，Go 还支持 `for-range` 遍历 channel。`range` 会持续从 channel 中接收值，直到 channel 被 `close()` 关闭。使用 `close(ch)` 可以通知接收方不会有更多数据发送，`for-range` 循环会自动退出。

channel 的零值是 `nil`，对 `nil` channel 的发送和接收都会永久阻塞，因此必须配合 `make` 使用。

## 示例代码

{{< code "../cmd/201-channel/main.go" >}}

## 运行方式

```bash
go run ./cmd/201-channel/
```

## 源码位置

[`cmd/201-channel/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/201-channel/main.go)
