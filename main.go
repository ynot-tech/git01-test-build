package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 1. Define the Home Route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a simple Go web server running in a Docker container on AWS!")
	})

	// 2. Define a Health Check (Crucial for Docker/AWS)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// 3. Get port from environment variable (Best practice for Cloud)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
