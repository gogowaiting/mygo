---
title: "GMP 模型"
weight: 10
description: "Goroutine、OS 线程、处理器的关系"
---

## 问题

请解释 Go 的 GMP 调度模型。

## 答案

GMP 是 Go 运行时的调度模型，由三个核心组件组成。G（Goroutine）是用户态协程，初始栈仅 2KB，可动态增长，创建成本极低。M（Machine）是操作系统内核线程，负责实际执行代码。P（Processor）是逻辑处理器，持有本地运行队列，数量由 `GOMAXPROCS` 决定（默认等于 CPU 核数）。

调度流程：M 必须绑定 P 才能执行 G。每个 P 维护一个本地 G 队列（最多 256 个），M 从绑定的 P 的队列中取 G 执行。当 P 的本地队列为空时，会从全局队列或其他 P 的队列中偷取 G（work stealing 机制），保证所有 CPU 核心都被充分利用。

当 G 发生阻塞（如系统调用、CGO）时，M 会释放 P，让其他空闲 M 接管该 P 继续执行队列中的其他 G。阻塞结束后，M 尝试获取一个空闲 P，若没有则将 G 放入全局队列，M 进入休眠。这种设计使得少量线程就能高效调度大量 goroutine。

与操作系统线程相比，goroutine 的栈空间小（2KB vs 1-8MB），切换成本低（用户态切换 vs 内核态切换），因此可以轻松创建数十万个 goroutine。

## 关键词

Goroutine, Machine, Processor, GOMAXPROCS, work stealing, 本地队列, 全局队列
