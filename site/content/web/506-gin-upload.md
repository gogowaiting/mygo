---
title: "506. 文件上传"
weight: 506
description: "单文件/多文件上传、静态文件服务"
---

## 知识点

文件上传是 Web 应用的常见需求。Gin 通过 `r.MaxMultipartMemory` 设置上传文件大小限制（如 `10 << 20` 表示 10MB），超出限制的文件会使用临时文件而非内存。

单文件上传使用 `c.FormFile("file")` 获取上传的文件头，`c.SaveUploadedFile(file, dst)` 保存到指定路径。多文件上传使用 `c.MultipartForm()` 获取表单，通过 `form.File["files"]` 获取文件列表，循环保存。

`r.Static("/static", "/tmp")` 将 URL 路径映射到本地目录，提供静态文件服务。`r.StaticFile("/favicon", path)` 映射单个文件。这在开发阶段非常方便，生产环境通常使用 Nginx 等反向代理处理静态文件。

## 示例代码

{{< code "../cmd/506-gin-upload/main.go" >}}

## 运行方式

```bash
go run ./cmd/506-gin-upload/
```

## 源码位置

[`cmd/506-gin-upload/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/506-gin-upload/main.go)
