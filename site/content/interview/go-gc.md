---
title: "GC 机制"
weight: 11
description: "三色标记、写屏障、STW"
---

## 问题

请简述 Go 的垃圾回收机制。

## 答案

Go 使用并发三色标记清除算法（Concurrent Tri-color Mark and Sweep）。GC 分为几个阶段：首先短暂的 STW（Stop The World）开启写屏障，然后并发标记阶段（与用户 goroutine 并行执行），接着再次短暂 STW 关闭写屏障并做最终标记，最后并发清除阶段回收白色对象。

三色标记中，白色表示未访问（待回收），灰色表示已访问但子对象未扫描，黑色表示已访问且子对象已扫描。标记从根对象（全局变量、栈上变量）开始，逐步扫描引用关系。写屏障（write barrier）在并发阶段确保黑色对象不会引用白色对象，防止误回收。

GC 触发条件有三个：堆内存增长到上次 GC 后的 `GOGC` 倍（默认 100，即堆翻倍时触发）；超过 2 分钟未 GC 则强制触发；手动调用 `runtime.GC()`。`GOGC=off` 可关闭自动 GC。

Go 1.19 引入了 `GOMEMLIMIT`，可以设置内存上限，GC 会在接近上限时更积极地回收，减少 OOM 风险。相比早期版本，现代 Go GC 的 STW 时间已经降到微秒级别。

## 关键词

三色标记, 写屏障, STW, GOGC, GOMEMLIMIT, 并发 GC, 标记清除
