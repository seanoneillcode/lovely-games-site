package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"seanoneillcode/lovely-games-site/handlers"
	"seanoneillcode/lovely-games-site/handlers/games"
	"seanoneillcode/lovely-games-site/html"
)

func main() {

	rand.Seed(time.Now().Unix())

	// Development mode flag is used to return embedded resources or read them from file system every request.
	isDevelopmentMode := flag.Bool("dev", false, "set to true for local development")
	flag.Parse()

	render := handlers.NewRenderHandlers(*isDevelopmentMode)
	templates := html.NewTemplates(*isDevelopmentMode)
	gameHandler := games.NewGameHandler(templates, games.NewRepository())

	http.HandleFunc("/", render.Index)
	http.HandleFunc("/games", gameHandler.ListGames)
	http.HandleFunc("/games/play", gameHandler.Play)
	http.HandleFunc("/games/frame", gameHandler.GameFrame)

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
