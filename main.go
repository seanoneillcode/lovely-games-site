package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"seanoneillcode/lovely-games-site/handlers"
)

func main() {

	isDevelopmentMode := flag.Bool("dev", false, "set to true for local development")
	flag.Parse()

	render := handlers.NewRenderHandlers(*isDevelopmentMode)

	http.HandleFunc("/", render.Index)
	http.HandleFunc("/about", render.About)

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
