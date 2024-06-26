package main

import (
	"log"
	"net/http"
	"net_http_middleware/controllers"
	"net_http_middleware/models"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func firstMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("In Middleware - Before handler")
		f(w, r) // original function call
		log.Print("In Middleware - After handler")
	}
}

func secondMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("In Middleware 2 - Before handler")
		f(w, r) // original function call
		log.Print("In Middleware 2 - After handler")
	}
}

// Middleware handler to log requests
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s: Time taken - %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func main() {
	addr := ":8080"

	models.ConnectDatabase()
	models.DBMigrate()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/blogs", firstMiddleware(secondMiddleware(controllers.BlogsIndex)))

	muxWithLogger := NewLogger(mux)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, muxWithLogger))
}
