---
title: "defer 机制"
weight: 17
description: "执行顺序、参数求值时机"
---

## 问题

defer 的执行顺序和参数求值时机是什么？

## 答案

`defer` 语句用于延迟执行函数调用，遵循两个关键规则。执行顺序：多个 defer 按 LIFO（后进先出）栈顺序执行，最后注册的 defer 最先执行。参数求值：defer 的参数在 defer 语句处立即求值，而不是在执行时求值。

经典示例：`defer fmt.Println(x)` 中的 x 在 defer 语句处就确定了值，即使后续修改 x 也不影响。但如果使用匿名函数闭包 `defer func() { fmt.Println(x) }()`，闭包捕获的是变量引用，执行时读取的是最新值。

defer 的常见用途：1）资源释放，如 `defer file.Close()`、`defer conn.Close()`；2）修改命名返回值（defer 可以读写命名返回值）；3）`defer recover()` 捕获 panic，必须在 defer 中调用 recover 才有效；4）计时和追踪，如 `defer timeTrack(time.Now(), "myFunc")`。

Go 1.21 之后，`defer` 的性能已经大幅优化。编译器会将简单的 defer（如 `defer f.Close()`）内联处理，避免堆分配。在性能敏感的热路径中，defer 的开销可以忽略不计。

## 关键词

LIFO, 参数立即求值, 闭包捕获, recover, 资源释放, 命名返回值
