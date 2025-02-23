package main

import (
	"log"
	"net/http"
	"time"

	"github.com/TD-17/library_management/internal/routes"
)

func main() {
	router := routes.SetupRouter()
	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server is running on port 8888...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
