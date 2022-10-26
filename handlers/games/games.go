package games

import (
	"net/http"

	"seanoneillcode/lovely-games-site/handlers/common"
	"seanoneillcode/lovely-games-site/html"
)

type GameHandler struct {
	templates  *html.Templates
	repository *Repository
}

func NewGameHandler(templates *html.Templates, repository *Repository) *GameHandler {
	return &GameHandler{
		templates:  templates,
		repository: repository,
	}
}

func (h *GameHandler) Games(w http.ResponseWriter, r *http.Request) {

	games, err := h.repository.GetGames()
	if err != nil {
		common.HandleError(err, w, r)
	}

	data := map[string]interface{}{
		"Title": "About",
		"Games": games,
	}
	err = h.templates.GetTemplate("games.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r)
	}
}

func (h *GameHandler) AddGame(w http.ResponseWriter, r *http.Request) {

	// parse request

	//
	game := Game{
		Id:          "3",
		Name:        "Indie Game",
		Description: "New indie game",
		Screenshot:  "indie.png",
	}
	err := h.repository.AddGame(game)
	if err != nil {
		common.HandleError(err, w, r)
	}
}
