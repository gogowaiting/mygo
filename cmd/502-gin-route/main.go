package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 502 路由进阶
// 知识点：路由分组、路径参数、查询参数
// 运行：go run ./cmd/502-gin-route
// 测试：
//   curl http://localhost:8082/user/42
//   curl http://localhost:8082/search?q=gin&page=2
//   curl http://localhost:8082/api/v1/items
//   curl http://localhost:8082/api/v2/items

func main() {
	r := gin.Default()

	// 路径参数 :name
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"user_id": id})
	})

	// 路径参数 + 通配符
	r.GET("/files/*filepath", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"filepath": c.Param("filepath")})
	})

	// 查询参数
	r.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "")
		page := c.DefaultQuery("page", "1")
		c.JSON(http.StatusOK, gin.H{
			"query": query,
			"page":  page,
		})
	})

	// 路由分组 v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/items", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "v1", "items": []string{"a", "b"}})
		})
		v1.GET("/items/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "v1", "id": c.Param("id")})
		})
	}

	// 路由分组 v2
	v2 := r.Group("/api/v2")
	{
		v2.GET("/items", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "v2", "items": []string{"x", "y", "z"}})
		})
	}

	// 支持多个 HTTP 方法
	r.POST("/echo", func(c *gin.Context) {
		body, _ := c.GetRawData()
		c.Data(http.StatusOK, "text/plain", body)
	})

	// Handle 多方法
	r.Any("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 打印所有注册路由
	for _, route := range r.Routes() {
		fmt.Printf("%-6s %s\n", route.Method, route.Path)
	}

	r.Run(":8082")
}
