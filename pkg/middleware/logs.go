package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK, // дефолтное значение
		}
		next.ServeHTTP(wrapper, r)
		log.Printf("%d %s %s %s %v", wrapper.StatusCode, r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}
