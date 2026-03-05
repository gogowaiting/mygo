---
title: "Context 核心用途"
weight: 14
description: "取消传播、超时控制、请求级数据"
---

## 问题

Go 中 Context 的核心用途是什么？

## 答案

`context.Context` 用于三个核心场景。第一是取消传播：当父操作取消时，所有子操作自动取消。`WithCancel` 返回 cancel 函数，调用后所有监听 `ctx.Done()` 的 goroutine 收到信号。`WithTimeout` 和 `WithDeadline` 在超时或到达截止时间后自动触发取消。务必 `defer cancel()` 释放资源。

第二是超时控制：为数据库查询、RPC 调用等操作设置合理的超时时间，防止 goroutine 泄漏。例如 `ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)` 限制操作最多 5 秒。

第三是请求级数据传递：`WithValue` 在调用链中传递请求 ID、认证信息、trace ID 等。应使用自定义类型作为 key（如 `type contextKey string`），避免与其他包冲突。

Context 是树形结构，子 Context 继承父 Context 的 Done 信号。关键规则：不要把 Context 存在 struct 中，应作为函数第一个参数显式传递（`func DoSomething(ctx context.Context, ...)`）。Context 是并发安全的，可以被多个 goroutine 同时使用。

## 关键词

WithCancel, WithTimeout, WithDeadline, WithValue, 取消传播, goroutine 泄漏
