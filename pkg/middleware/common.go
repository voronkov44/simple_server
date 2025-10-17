package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (writer *WrapperWriter) WriteHeader(statusCode int) {
	writer.ResponseWriter.WriteHeader(statusCode)
	writer.StatusCode = statusCode
}
