---
title: "507. CRUD 应用"
weight: 507
description: "RESTful CRUD、并发安全存储、Handler 结构"
---

## 知识点

本示例是一个完整的 RESTful CRUD 应用，展示了生产级 Gin 项目的典型结构。分为三层：模型层（Article 结构体）、存储层（Store 使用 `sync.RWMutex` 保证并发安全）、处理器层（Handler 依赖 Store）。

Store 使用内存 map 存储数据，`sync.RWMutex` 的读写锁分离：`RLock()` 用于读操作（List、Get），`Lock()` 用于写操作（Create、Update、Delete），提高并发性能。自增 ID 通过 `nextID` 字段管理。

RESTful 路由映射：`GET /articles` 列表，`POST /articles` 创建，`GET /articles/:id` 查询，`PUT /articles/:id` 更新，`DELETE /articles/:id` 删除。正确的 HTTP 状态码：201 Created、400 Bad Request、404 Not Found。

## 示例代码

{{< code "../cmd/507-gin-crud/main.go" >}}

## 运行方式

```bash
go run ./cmd/507-gin-crud/
```

## 源码位置

[`cmd/507-gin-crud/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/507-gin-crud/main.go)
