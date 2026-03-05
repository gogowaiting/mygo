package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 506 文件上传与静态服务
// 知识点：单文件上传、多文件上传、静态文件服务
// 运行：go run ./cmd/506-gin-upload
// 测试：
//   curl -X POST http://localhost:8086/upload -F "file=@/tmp/test.txt"
//   curl http://localhost:8086/static/README.md
//   curl http://localhost:8086/               (查看首页)

func main() {
	r := gin.Default()

	// 限制上传文件大小：10MB
	r.MaxMultipartMemory = 10 << 20

	// 首页（表单上传页面）
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
			<!DOCTYPE html>
			<html>
			<body>
				<h2>单文件上传</h2>
				<form action="/upload" method="post" enctype="multipart/form-data">
					<input type="file" name="file">
					<button type="submit">上传</button>
				</form>
				<h2>多文件上传</h2>
				<form action="/upload-multi" method="post" enctype="multipart/form-data">
					<input type="file" name="files" multiple>
					<button type="submit">上传</button>
				</form>
			</body>
			</html>
		`))
	})

	// 单文件上传
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filename := filepath.Base(file.Filename)
		dst := filepath.Join("/tmp", filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"filename": filename,
			"size":     file.Size,
			"saved_to": dst,
		})
	})

	// 多文件上传
	r.POST("/upload-multi", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		files := form.File["files"]
		var uploaded []string
		for _, file := range files {
			dst := filepath.Join("/tmp", filepath.Base(file.Filename))
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			uploaded = append(uploaded, file.Filename)
		}

		c.JSON(http.StatusOK, gin.H{
			"count":     len(uploaded),
			"filenames": uploaded,
		})
	})

	// 静态文件服务：将 /static 路径映射到指定目录
	// 需要先创建目录：mkdir -p /tmp/static
	r.Static("/static", "/tmp")
	// 或者单个文件
	r.StaticFile("/favicon", "/tmp/favicon.ico")

	// 生成测试文件
	r.GET("/create-test", func(c *gin.Context) {
		dst := "/tmp/test.txt"
		if err := writeFile(dst, "hello gin upload"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"created": dst})
	})

	fmt.Println("Visit http://localhost:8086/ to upload files")
	r.Run(":8086")
}

func writeFile(path, content string) error {
	return writeFileBytes(path, []byte(content))
}

func writeFileBytes(path string, data []byte) error {
	return writeFileData(path, data)
}

func writeFileData(path string, data []byte) error {
	return nil // 简化，实际用 os.WriteFile
}
