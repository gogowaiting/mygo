package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 505 中间件
// 知识点：自定义中间件、全局/路由/分组中间件、中间件链
// 运行：go run ./cmd/505-gin-middleware
// 测试：
//   curl http://localhost:8085/public
//   curl -H "Authorization: mytoken" http://localhost:8085/admin/dashboard
//   curl http://localhost:8085/admin/dashboard  (401)

// ========== 中间件定义 ==========

// Logger 请求日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next() // 执行后续处理器

		latency := time.Since(start)
		status := c.Writer.Status()
		log.Printf("[GIN] %s %s %d %v", c.Request.Method, path, status, latency)
	}
}

// Auth 简单鉴权中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}
		// 模拟 token 校验
		if token != "mytoken" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid token",
			})
			return
		}
		// 将用户信息存入上下文
		c.Set("user", "admin")
		c.Next()
	}
}

// Recovery 自定义 panic 恢复
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC] %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
			}
		}()
		c.Next()
	}
}

func main() {
	r := gin.New() // 不使用 Default，手动注册中间件

	// 全局中间件
	r.Use(Recovery())
	r.Use(Logger())

	// 公开路由（无鉴权）
	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "public access"})
	})

	// 需要鉴权的路由分组
	admin := r.Group("/admin", Auth())
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(http.StatusOK, gin.H{
				"user":    user,
				"message": "welcome to admin dashboard",
			})
		})
		admin.GET("/settings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"settings": "all settings here"})
		})
	}

	// 路由级别中间件（仅对单个路由生效）
	r.GET("/slow", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{"message": "slow response"})
	})

	r.Run(":8085")
}
