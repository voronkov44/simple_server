package hello

import (
	"fmt"
	"hello/pkg/res"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}
	router.HandleFunc("GET /ping", handler.Ping())
	router.HandleFunc("GET /hello", handler.Hello())
}

func (handler *HelloHandler) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := Response{
			Answer: "pong",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Anonymous"
		}

		data := Response{
			Answer: fmt.Sprintf("Hello, %s!", name),
		}
		res.Json(w, data, http.StatusOK)
	}
}
