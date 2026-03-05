package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 501 Gin 基础入门
// 知识点：引擎初始化、默认路由、启动服务
// 运行：go run ./cmd/501-gin-basic
// 测试：curl http://localhost:8081/ping

func main() {
	// gin.Default() 创建带 Logger 和 Recovery 中间件的引擎
	r := gin.Default()

	// GET 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// GET 带简单响应
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Gin!")
	})

	// GET 返回 HTML
	r.GET("/hello", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>Hello Gin!</h1>"))
	})

	r.Run(":8081")
}
