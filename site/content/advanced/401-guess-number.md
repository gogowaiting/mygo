---
title: "401. 猜数字"
weight: 401
description: "bufio 读取输入、strconv 转换、游戏循环"
---

## 知识点

猜数字是一个经典的交互式 CLI 程序，综合运用了 Go 的标准输入读取、字符串处理和随机数生成。程序首先使用 `rand.Seed(time.Now().UnixNano())` 初始化随机数种子，确保每次运行产生不同的目标数字。

核心输入处理链路为：`bufio.NewReader(os.Stdin)` 创建带缓冲的读取器，`ReadString('\n')` 读取一行输入，`strings.Trim` 去除换行符，最后 `strconv.Atoi` 将字符串转换为整数。游戏循环中通过 `for` 无限循环持续读取用户输入，根据猜测结果给出"大了/小了"反馈，猜中后 `break` 退出。

错误处理贯穿整个程序：读取输入失败、字符串转整数失败都会输出错误信息并 `return`。这种模式是 CLI 程序的标准做法。

## 示例代码

{{< code "../cmd/401-guess-number/main.go" >}}

## 运行方式

```bash
go run ./cmd/401-guess-number/
```

## 源码位置

[`cmd/401-guess-number/main.go`](https://github.com/gogowaiting/mygo/blob/main/cmd/401-guess-number/main.go)
