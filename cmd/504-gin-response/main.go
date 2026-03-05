package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 504 响应处理
// 知识点：JSON、XML、YAML、重定向、文件下载
// 运行：go run ./cmd/504-gin-response
// 测试：
//   curl http://localhost:8084/json
//   curl http://localhost:8084/xml
//   curl http://localhost:8084/redirect
//   curl http://localhost:8084/notfound

type Article struct {
	ID    int      `json:"id"   xml:"id"`
	Title string   `json:"title" xml:"title"`
	Tags  []string `json:"tags"  xml:"tags>tag"`
}

func main() {
	r := gin.Default()

	// JSON 响应
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data":    Article{ID: 1, Title: "Gin 入门", Tags: []string{"go", "web"}},
		})
	})

	// 结构体直接序列化
	r.GET("/json-struct", func(c *gin.Context) {
		article := Article{ID: 2, Title: "Gin 进阶", Tags: []string{"gin", "middleware"}}
		c.IndentedJSON(http.StatusOK, article) // 格式化输出
	})

	// XML 响应
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, Article{ID: 3, Title: "Gin XML", Tags: []string{"xml"}})
	})

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8084/json")
	})

	// 404 自定义
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "page not found",
		})
	})

	// 模拟 404
	r.GET("/notfound", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
	})

	// 成功响应封装
	r.GET("/articles/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "999" {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "article not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": Article{ID: 1, Title: "文章 " + id},
		})
	})

	r.Run(":8084")
}
