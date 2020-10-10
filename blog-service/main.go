package main

import (
	router "github.com/go-programming-tour-book/blog-service/internal/routers"
	"net/http"
	"time"
)

func main(){
	// 1.go get -u github.com/gin-gonic/gin@v1.5.0
	router := router.NewRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
