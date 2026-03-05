---
title: "115. 接口"
weight: 115
description: "定义、实现、空接口、多态"
---

## 知识点

接口（interface）是方法声明的集合。任何类型只要实现了接口中定义的所有方法，就隐式地实现了该接口，无需显式声明。这种隐式实现解耦了接口定义和实现，是 Go 接口设计的核心思想。

接口可以作为函数参数类型，接受所有实现了该接口的类型，从而实现多态。一个类型的指针接收者实现了接口，那么它的值类型也会隐式地实现该接口。通过接口，可以编写通用的处理逻辑，而不关心具体的类型。

空接口 `interface{}` 没有任何方法要求，因此所有类型都实现了空接口。空接口变量可以赋任何类型的值，类似于其他语言中的 `any` 或 `Object`。这在需要处理未知类型的场景中非常有用。

## 示例代码

{{< code "../cmd/115-interface/main.go" >}}

## 运行方式

```bash
go run ./cmd/115-interface/
```

## 源码位置

[`cmd/115-interface/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/115-interface/main.go)
