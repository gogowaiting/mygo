---
title: "204. 定时与超时"
weight: 204
description: "Timer、Ticker、select 超时控制"
---

## 知识点

Go 的 `time` 包提供了多种定时机制。`time.NewTimer` 创建一个单次定时器，到期后向 `timer.C` channel 发送当前时间。`time.NewTicker` 创建周期性定时器，每隔指定时间向 `ticker.C` 发送一次，适合需要定期执行的场景。

`select` 语句配合 `time.After` 可以实现超时控制。`time.After(d)` 返回一个 channel，在指定时间后会收到一个值。将它放入 `select` 的 case 中，如果其他操作先完成就正常处理，否则超时退出。

`time.AfterFunc` 可以在指定时间后执行一个回调函数，适合需要延迟执行某个操作的场景。使用完毕后应调用 `timer.Stop()` / `ticker.Stop()` 释放资源。

## 示例代码

{{< code "../cmd/204-timer/main.go" >}}

## 运行方式

```bash
go run ./cmd/204-timer/
```

## 源码位置

[`cmd/204-timer/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/204-timer/main.go)
