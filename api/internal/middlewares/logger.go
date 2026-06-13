package middlewares

import (
	"log"
	"net/http"
	"time"
)

type responseWriterWrapper struct {
	ResponseWriter http.ResponseWriter
	statusCode     int
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		start := time.Now()
		next.ServeHTTP(wrapped, r)
		log.Println(r.Method, r.URL.Path, wrapped.statusCode, time.Since(start))
	})

}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterWrapper) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *responseWriterWrapper) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}
