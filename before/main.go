package main

import (
	"log"
	"net/http"
	"net_http_middleware/controllers"
	"net_http_middleware/models"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	addr := ":8080"

	models.ConnectDatabase()
	models.DBMigrate()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/blogs", controllers.BlogsIndex)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
