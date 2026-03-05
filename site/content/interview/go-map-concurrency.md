---
title: "Map 并发安全"
weight: 16
description: "原生 map 不支持并发读写"
---

## 问题

为什么 Go 的原生 map 不是并发安全的？

## 答案

Go 的 map 内部使用哈希表实现，读写操作会修改内部结构（如扩容时重新分配 bucket）。当多个 goroutine 同时读写同一个 map 时，会触发运行时检测，输出 `fatal error: concurrent map writes` 并直接崩溃程序。注意这是 fatal error，不是 panic，无法通过 `recover()` 捕获。

Go 从 1.6 开始在运行时加入了 map 并发写的检测。即使两个 goroutine 写不同的 key，只要 map 正在扩容就会触发崩溃。并发读写（一个读一个写）同样不安全，会导致数据竞争。

解决方案有三种：1）`sync.RWMutex` 加锁保护，适合大多数场景，读多写少时用 `RLock()` 提高并发度；2）`sync.Map`，专为读多写少场景优化，提供 `Load`/`Store`/`LoadOrStore`/`Delete` 方法，无需显式加锁；3）Channel 串行化访问，将 map 操作通过 channel 发送给单一 goroutine 执行，符合"不要通过共享内存来通信"的理念。

选择建议：通用场景用 `sync.RWMutex`；读远多于写且 key 相对稳定用 `sync.Map`；需要复杂操作组合时用 Channel 模式。

## 关键词

concurrent map writes, sync.RWMutex, sync.Map, Channel 串行化, fatal error
