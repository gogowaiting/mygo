package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// 507 综合 CRUD 示例
// 知识点：RESTful 风格、内存存储、结构化路由、错误处理
// 运行：go run ./cmd/507-gin-crud
// 测试：
//   curl http://localhost:8087/articles
//   curl -X POST http://localhost:8087/articles -H "Content-Type: application/json" -d '{"title":"Hello","content":"World"}'
//   curl http://localhost:8087/articles/1
//   curl -X PUT http://localhost:8087/articles/1 -H "Content-Type: application/json" -d '{"title":"Updated","content":"New"}'
//   curl -X DELETE http://localhost:8087/articles/1

// ========== 模型定义 ==========

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"   binding:"required"`
	Content string `json:"content" binding:"required"`
}

// ========== 内存存储 ==========

type Store struct {
	mu     sync.RWMutex
	items  map[int]*Article
	nextID int
}

func NewStore() *Store {
	return &Store{
		items:  make(map[int]*Article),
		nextID: 1,
	}
}

func (s *Store) Create(a *Article) *Article {
	s.mu.Lock()
	defer s.mu.Unlock()
	a.ID = s.nextID
	s.nextID++
	s.items[a.ID] = a
	return a
}

func (s *Store) Get(id int) (*Article, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.items[id]
	return a, ok
}

func (s *Store) List() []*Article {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]*Article, 0, len(s.items))
	for _, a := range s.items {
		result = append(result, a)
	}
	return result
}

func (s *Store) Update(id int, title, content string) (*Article, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	a, ok := s.items[id]
	if !ok {
		return nil, false
	}
	a.Title = title
	a.Content = content
	return a, true
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[id]; !ok {
		return false
	}
	delete(s.items, id)
	return true
}

// ========== 处理器 ==========

type Handler struct {
	store *Store
}

func (h *Handler) List(c *gin.Context) {
	articles := h.store.List()
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": articles})
}

func (h *Handler) Create(c *gin.Context) {
	var a Article
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": err.Error()})
		return
	}
	created := h.store.Create(&a)
	c.JSON(http.StatusCreated, gin.H{"code": 0, "data": created})
}

func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": "invalid id"})
		return
	}
	a, ok := h.store.Get(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "error": "article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": a})
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": "invalid id"})
		return
	}
	var a Article
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": err.Error()})
		return
	}
	updated, ok := h.store.Update(id, a.Title, a.Content)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "error": "article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": updated})
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": "invalid id"})
		return
	}
	if !h.store.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "error": "article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "deleted"})
}

// ========== 主函数 ==========

func main() {
	r := gin.Default()
	h := &Handler{store: NewStore()}

	// RESTful 路由
	articles := r.Group("/articles")
	{
		articles.GET("", h.List)          // 列表
		articles.POST("", h.Create)       // 创建
		articles.GET("/:id", h.Get)       // 查询
		articles.PUT("/:id", h.Update)    // 更新
		articles.DELETE("/:id", h.Delete) // 删除
	}

	r.Run(":8087")
}
