---
title: "504. 响应格式"
weight: 504
description: "JSON、XML、重定向、自定义 404"
---

## 知识点

Gin 支持多种响应格式，满足不同 API 需求。`c.JSON()` 返回紧凑 JSON，`c.IndentedJSON()` 返回格式化 JSON（便于调试）。`c.XML()` 返回 XML 响应，通过 struct tag（如 `xml:"tags>tag"`）控制 XML 结构。

`c.Redirect(status, url)` 实现 HTTP 重定向，状态码使用 `http.StatusMovedPermanently`（301）或 `http.StatusFound`（302）。`r.NoRoute(handler)` 注册全局 404 处理器，未匹配到路由时自动调用。

实际项目中通常封装统一的响应结构：`gin.H{"code": 0, "data": ..., "message": "success"}` 表示成功，`gin.H{"code": 404, "error": "..."}` 表示错误，便于前端统一处理。

## 示例代码

{{< code "../cmd/504-gin-response/main.go" >}}

## 运行方式

```bash
go run ./cmd/504-gin-response/
```

## 源码位置

[`cmd/504-gin-response/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/504-gin-response/main.go)
