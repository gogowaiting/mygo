---
title: "202. Goroutine"
weight: 202
description: "go 关键字、轻量级线程、生命周期"
---

## 知识点

Goroutine 是 Go 中的轻量级线程，由 `go` 关键字启动。与操作系统线程相比，goroutine 的创建和切换开销极小（初始栈仅几 KB），因此可以轻松创建成千上万个并发任务。

所有 goroutine 共享同一地址空间，但各自拥有独立的栈。需要注意的是，所有 goroutine 会随着 `main` 函数的结束而终止，不论它们是否已经执行完毕。为了观察并发输出，可以在 main 函数末尾使用 `fmt.Scanln` 等方式阻塞主 goroutine。

下面的示例中，`running` 函数在一个无限循环中每隔一秒打印一次计数。通过 `go running()` 启动后，主 goroutine 通过 `Scanln` 等待用户输入，从而让子 goroutine 持续运行。

## 示例代码

{{< code "../cmd/202-goroutine/main.go" >}}

## 运行方式

```bash
go run ./cmd/202-goroutine/
```

## 源码位置

[`cmd/202-goroutine/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/202-goroutine/main.go)
