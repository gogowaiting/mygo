---
title: "503. 请求绑定"
weight: 503
description: "JSON/Form/Query 绑定、验证规则"
---

## 知识点

请求绑定是将 HTTP 请求数据自动映射到 Go 结构体的过程。Gin 支持 JSON、Form、Query 等多种绑定方式，并通过 `binding` struct tag 实现数据校验。

`c.ShouldBindJSON(&u)` 绑定 JSON Body，`c.ShouldBind(&u)` 绑定 Form 表单，`c.ShouldBindQuery(&u)` 绑定查询参数。`ShouldBind` 系列方法出错时返回 error，由开发者自行处理（推荐）。`MustBind` 出错自动返回 400，不推荐生产使用。

`binding` tag 支持丰富的校验规则：`required`（必填）、`min`/`max`（长度限制）、`gte`/`lte`（数值范围）、`email`（邮箱格式）、`omitempty`（可选字段）。校验失败时返回 400 状态码和错误详情。

## 示例代码

{{< code "../cmd/503-gin-bind/main.go" >}}

## 运行方式

```bash
go run ./cmd/503-gin-bind/
```

## 源码位置

[`cmd/503-gin-bind/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/503-gin-bind/main.go)
