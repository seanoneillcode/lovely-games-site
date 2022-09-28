package main

import (
	"log"
	"net/http"
	"os"

	"seanoneillcode/lovely-games-site/html"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	p := html.AboutParams{
		Title: "About",
	}
	html.About(w, p)
}

func index(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{
		Title: "Index",
	}
	html.Index(w, p)
}
