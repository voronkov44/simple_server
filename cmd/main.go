package main

import (
	"fmt"
	"hello/internal/hello"
	"hello/pkg/middleware"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// Handler
	hello.NewHelloHandler(router)

	// Middlewares
	stack := middleware.Chain(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Listening on port 8080")
	server.ListenAndServe()
}
