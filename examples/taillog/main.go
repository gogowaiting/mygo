package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	defaultLogFilePath = "./logs/train.log"
	defaultAddr        = ":8080"
)

func main() {
	cfg := loadConfig()
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isOriginAllowed(r.Header.Get("Origin"), cfg.allowedOrigins)
		},
	}

	http.HandleFunc("/ws/logs", func(w http.ResponseWriter, r *http.Request) {
		handleLogWebSocket(w, r, cfg, upgrader)
	})
	fmt.Printf("taillog listening on %s\n", cfg.addr)
	log.Fatal(http.ListenAndServe(cfg.addr, nil))
}

type config struct {
	addr           string
	logFilePath    string
	token          string
	allowedOrigins map[string]struct{}
}

func loadConfig() config {
	addr := getenv("TAILLOG_ADDR", defaultAddr)
	logFilePath := getenv("TAILLOG_FILE", defaultLogFilePath)
	token := os.Getenv("TAILLOG_TOKEN")
	allowedOrigins := parseOrigins(os.Getenv("TAILLOG_ALLOWED_ORIGINS"))
	if len(allowedOrigins) == 0 {
		allowedOrigins["http://localhost"] = struct{}{}
		allowedOrigins["http://localhost:3000"] = struct{}{}
		allowedOrigins["http://127.0.0.1"] = struct{}{}
		allowedOrigins["http://127.0.0.1:3000"] = struct{}{}
	}
	return config{
		addr:           addr,
		logFilePath:    logFilePath,
		token:          token,
		allowedOrigins: allowedOrigins,
	}
}

func getenv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func parseOrigins(v string) map[string]struct{} {
	res := make(map[string]struct{})
	for _, origin := range strings.Split(v, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			res[origin] = struct{}{}
		}
	}
	return res
}

func isOriginAllowed(origin string, allowed map[string]struct{}) bool {
	origin = strings.TrimSpace(origin)
	if origin == "" {
		return false
	}
	if _, ok := allowed["*"]; ok {
		return true
	}
	_, ok := allowed[origin]
	return ok
}

func isAuthorized(r *http.Request, cfg config) bool {
	if cfg.token == "" {
		return true
	}
	authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
	expected := "Bearer " + cfg.token
	if authHeader == expected {
		return true
	}
	return strings.TrimSpace(r.URL.Query().Get("token")) == cfg.token
}

func handleLogWebSocket(w http.ResponseWriter, r *http.Request, cfg config, upgrader websocket.Upgrader) {
	if !isAuthorized(r, cfg) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	file, err := os.Open(cfg.logFilePath)
	if err != nil {
		_ = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("日志文件不存在: %s", cfg.logFilePath)))
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Println("Stat failed:", err)
		return
	}
	if _, err := file.Seek(stat.Size(), io.SeekStart); err != nil {
		log.Println("Seek failed:", err)
		return
	}

	reader := bufio.NewReader(file)
	clientDone := make(chan struct{})
	go func() {
		defer close(clientDone)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-clientDone:
			return
		default:
		}

		line, err := reader.ReadString('\n')
		if errors.Is(err, io.EOF) {
			select {
			case <-clientDone:
				return
			case <-ticker.C:
			}
			continue
		}
		if err != nil {
			log.Println("read log failed:", err)
			return
		}
		if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
			log.Println("WebSocket write error:", err)
			return
		}
	}
}
