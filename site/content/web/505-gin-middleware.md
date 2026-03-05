---
title: "505. 中间件"
weight: 505
description: "自定义中间件、全局/分组/路由级"
---

## 知识点

中间件是 Gin 处理请求链的核心机制。每个中间件是一个 `gin.HandlerFunc`，通过 `c.Next()` 调用后续处理器，通过 `c.AbortWithStatusJSON()` 中断链并返回响应。

`gin.New()` 创建空白引擎（不带任何中间件），`gin.Default()` 自带 Logger 和 Recovery。自定义中间件示例：Logger 记录请求耗时，Auth 校验 Authorization header，Recovery 捕获 panic。中间件通过 `c.Set(key, value)` 向后续处理器传递数据，`c.Get(key)` 读取。

中间件可以挂载在三个级别：全局 `r.Use()` 对所有路由生效，分组级 `r.Group("/admin", Auth())` 对分组内路由生效，路由级通过包装函数实现。执行顺序按照注册顺序，请求从外向内穿过中间件链。

## 示例代码

{{< code "../cmd/505-gin-middleware/main.go" >}}

## 运行方式

```bash
go run ./cmd/505-gin-middleware/
```

## 源码位置

[`cmd/505-gin-middleware/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/505-gin-middleware/main.go)
