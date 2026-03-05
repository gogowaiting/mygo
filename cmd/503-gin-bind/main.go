package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 503 请求绑定与校验
// 知识点：JSON 绑定、表单绑定、Query 绑定、validator 标签
// 运行：go run ./cmd/503-gin-bind
// 测试：
//   curl -X POST http://localhost:8083/json -H "Content-Type: application/json" -d '{"name":"Alice","age":25}'
//   curl -X POST http://localhost:8083/form -d "name=Bob&age=30"
//   curl "http://localhost:8083/query?name=Charlie&age=35"
//   curl -X POST http://localhost:8083/json -H "Content-Type: application/json" -d '{"name":"","age":0}'

// User 定义请求结构体，用 binding tag 指定校验规则
type User struct {
	Name string `json:"name" form:"name" binding:"required,min=1,max=50"`
	Age  int    `json:"age"  form:"age"  binding:"required,gte=1,lte=150"`
}

// UserUpdate 部分字段可选
type UserUpdate struct {
	Name  string `json:"name"  form:"name"`
	Email string `json:"email" form:"email" binding:"omitempty,email"`
}

func main() {
	r := gin.Default()

	// JSON 绑定：Content-Type: application/json
	r.POST("/json", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": u, "source": "json"})
	})

	// 表单绑定：Content-Type: application/x-www-form-urlencoded
	r.POST("/form", func(c *gin.Context) {
		var u User
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": u, "source": "form"})
	})

	// Query 参数绑定
	r.GET("/query", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindQuery(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": u, "source": "query"})
	})

	// ShouldBind vs MustBind
	// ShouldBind：出错返回 error，开发者自行处理
	// MustBind：出错自动返回 400，不推荐生产使用
	r.POST("/update", func(c *gin.Context) {
		var u UserUpdate
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"updated": u})
	})

	r.Run(":8083")
}
