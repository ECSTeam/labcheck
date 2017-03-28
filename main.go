package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8101"
	}

	router := NewRouter()
	log.Printf("starting on port:" + port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
