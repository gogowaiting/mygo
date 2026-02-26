package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

const (
	logFilePath = "./logs/train.log" // 训练日志路径
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有跨域请求，生产环境需加认证
	},
}

func main() {
	http.HandleFunc("/ws/logs", handleLogWebSocket)
	fmt.Println("🚀 Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLogWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// 打开日志文件
	file, err := os.Open(logFilePath)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("日志文件不存在: %s", logFilePath)))
		return
	}
	defer file.Close()

	// 定位到文件末尾，跳过历史日志
	stat, err := file.Stat()
	if err != nil {
		log.Println("Stat failed:", err)
		return
	}
	file.Seek(stat.Size(), 0)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
			log.Println("WebSocket write error:", err)
			break
		}
	}
}
