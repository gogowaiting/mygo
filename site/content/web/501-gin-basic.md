---
title: "501. Gin 基础"
weight: 501
description: "gin.Default、路由注册、JSON 响应"
---

## 知识点

Gin 是 Go 最流行的 Web 框架之一，以高性能和简洁 API 著称。`gin.Default()` 创建一个引擎实例，自动附带 Logger（请求日志）和 Recovery（panic 恢复）两个内置中间件，适合快速开发。

路由注册采用 `r.GET(path, handler)` 模式，handler 是 `func(c *gin.Context)` 类型的闭包。`gin.Context` 封装了请求和响应的所有操作。响应方式包括：`c.JSON()` 返回 JSON（最常用），`c.String()` 返回纯文本，`c.Data()` 返回原始字节（可用于 HTML）。

`gin.H` 是 `map[string]interface{}` 的快捷类型，用于快速构造 JSON 响应体。`r.Run(":8081")` 启动 HTTP 服务器，监听指定端口。

## 示例代码

{{< code "../cmd/501-gin-basic/main.go" >}}

## 运行方式

```bash
go run ./cmd/501-gin-basic/
```

## 源码位置

[`cmd/501-gin-basic/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/501-gin-basic/main.go)
