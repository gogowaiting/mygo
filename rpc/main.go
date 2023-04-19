package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/home", &HomeHandler{})

	http.ListenAndServe(":8080", nil)

}

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("hello, go"))
	fmt.Println("connect success")
}
