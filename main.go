package main

import (
	"log"
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8101"
	}

	log.Printf("starting on port:" + port)

	server := NewServer()
	server.Run(":" + port)

	//log.Fatal(http.ListenAndServe(":"+port, router))
}

func NewServer() *negroni.Negroni {

	n := negroni.Classic()
	router := NewRouter()

	n.UseHandler(router)
	return n
}
