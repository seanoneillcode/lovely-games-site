package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	f := NewFileHandler("./static")
	http.HandleFunc("/", f.handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
