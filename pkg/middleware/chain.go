package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Chain - функция для удобства, простой цикл для оборачивания middleware, принимаем неограниченное число middleware
func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		// длина - 1 - так как индекс начинается с нуля
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}

// chain нужен для того, чтобы в main не оборачивать роутер вот так - middleware.Logging(router), так как если мидлваров будет много, оборачивание будет выглядеть вот так, что не очень удобно - middleware.CORS(middleware.Logging(router)), stack := middleware.Chain(middleware.CORS, middleware.Logging)
