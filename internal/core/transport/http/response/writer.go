package core_http_response

import "net/http"

var statusCodeUninitialized = -1

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     statusCodeUninitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.ResponseWriter.WriteHeader(code)
	rw.statusCode = code
}

func (rw *ResponseWriter) GetStatusCodeOrPanic() int {
	if rw.statusCode == statusCodeUninitialized {
		panic("status code is not initialized")
	}
	return rw.statusCode
}
