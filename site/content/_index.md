---
title: "技术学习之旅"
layout: "home"
---

欢迎来到 Amos 的技术笔记。这里是一个按知识点拆分、可运行、可测试的学习平台。

## 学习路径

| 阶段 | 主题 | 内容 | 难度 |
|------|------|------|------|
| 1 | 基础语法 | 变量、控制流、数据结构、函数、指针、结构体、接口 | 入门 |
| 2 | 并发编程 | Goroutine、Channel、定时器 | 中级 |
| 3 | 算法 | 冒泡、选择、插入、快速排序 | 入门-中级 |
| 4 | 进阶模式 | Context、Worker Pool、sync.Pool、Atomic、泛型 | 高级 |
| 5 | Web 开发 | Gin 框架：路由、绑定、中间件、CRUD | 中级 |
| 6 | 面试题库 | GMP、GC、逃逸分析、Channel、接口等 | 中级-高级 |
| 7 | 技术博客 | Python、Git、GC 原理等 | 混合 |

## 快速开始

```bash
# 克隆项目
git clone https://github.com/gogowaiting/mygo.git
cd mygo

# 查看所有知识点
go run ./cmd/learn list

# 运行某个 demo
go run ./cmd/learn run basic-01-hello
```
