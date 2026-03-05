package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", newMux()); err != nil {
		fmt.Println("server exited:", err)
	}
}

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	_, _ = response.Write([]byte("hello, go"))
	fmt.Println("connect success")
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/home", &HomeHandler{})
	return mux
}
