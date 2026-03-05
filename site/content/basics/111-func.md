---
title: "111. 函数"
weight: 111
description: "函数签名、多返回值、不定参数、闭包"
---

## 知识点

函数是 Go 中实现功能的基本单元，接受参数并返回结果。Go 函数支持多返回值，这在错误处理中非常常见，函数可以同时返回结果和错误信息。命名规则上，函数名首字母大写表示包外可导出，首字母小写则只能在包内部调用。

Go 支持可变参数（variadic parameters），通过 `...type` 语法接收不定数量的同类型参数。调用时使用 `slice...` 可以将切片打散传入。此外 Go 还支持匿名函数和闭包，函数可以作为返回值，捕获外部变量形成闭包状态。

`init()` 是特殊的初始化函数，在模块被导入时自动执行，无需手动调用，常用于包级别的初始化操作。

## 示例代码

{{< code "../cmd/111-func/main.go" >}}

## 运行方式

```bash
go run ./cmd/111-func/
```

## 源码位置

[`cmd/111-func/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/111-func/main.go)
