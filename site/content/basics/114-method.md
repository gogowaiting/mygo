---
title: "114. 方法"
weight: 114
description: "值接收者、指针接收者、接口多态"
---

## 知识点

Go 的方法是带有接收者（receiver）的函数。接收者可以是值类型或指针类型，决定了方法能否修改原始数据。值接收者 `func (r rect) perim() int` 操作的是数据的副本，不会修改原始值；指针接收者 `func (r *rect) scale(factor int)` 可以直接修改原始结构体的字段。

指针类型的变量可以调用值接收者方法，Go 会自动解引用。但反过来，值类型的变量不一定能调用指针接收者方法（取决于是否可寻址）。方法值（method value）可以将方法绑定到具体实例，保存为变量后调用。

通过定义接口并让不同类型实现接口方法，可以实现多态。将不同类型放入同一个接口切片中，用统一的方式处理，这是 Go 实现面向对象多态的核心模式。

## 示例代码

{{< code "../cmd/114-method/main.go" >}}

## 运行方式

```bash
go run ./cmd/114-method/
```

## 源码位置

[`cmd/114-method/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/114-method/main.go)
