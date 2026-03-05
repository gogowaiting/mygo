---
title: "502. 路由"
weight: 502
description: "路径参数、通配符、查询参数、路由分组"
---

## 知识点

Gin 的路由系统支持路径参数、通配符、查询参数和路由分组，是构建 RESTful API 的基础。

路径参数使用 `:name` 语法定义（如 `/user/:id`），通过 `c.Param("id")` 获取。通配符使用 `*` 语法（如 `/files/*filepath`），匹配剩余所有路径。查询参数通过 `c.DefaultQuery("q", defaultValue)` 获取，可以设置默认值。

路由分组 `r.Group("/api/v1")` 将具有相同前缀的路由组织在一起，分组内可以应用中间件。`r.Any("/health")` 注册所有 HTTP 方法。`r.Routes()` 返回所有已注册路由，可用于调试。

## 示例代码

{{< code "../cmd/502-gin-route/main.go" >}}

## 运行方式

```bash
go run ./cmd/502-gin-route/
```

## 源码位置

[`cmd/502-gin-route/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/502-gin-route/main.go)
