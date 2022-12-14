package games

import (
	"errors"
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

func (h *GameHandler) ListGames(w http.ResponseWriter, r *http.Request) {

	allGames := h.repository.GetGames()
	data := map[string]interface{}{
		"Title": "About",
		"Games": allGames,
	}
	err := h.templates.GetWrappedTemplate("games/games.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
	}
}

func (h *GameHandler) Play(w http.ResponseWriter, r *http.Request) {
	id := common.GetQueryParam("id", r.URL)

	game := h.repository.GetGame(id)
	if game == nil {
		common.HandleError(errors.New("game not found"), w, r, http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"Title": "Play",
		"Game":  game,
	}
	err := h.templates.GetWrappedTemplate("games/play.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
	}
}
