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

	games, err := h.repository.GetGames()
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Title": "About",
		"Games": games,
	}
	err = h.templates.GetTemplate("games/games.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
	}
}

func (h *GameHandler) UploadGame(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		data := map[string]interface{}{
			"Title": "Upload Game",
		}
		err := h.templates.GetTemplate("games/upload.html").Execute(w, data)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
		}
	case "POST":
		if err := r.ParseForm(); err != nil {
			common.HandleError(err, w, r, http.StatusBadRequest)
		}
		name := r.FormValue("game-name")
		description := r.FormValue("game-description")

		game := Game{
			Id:          common.RandomCharacters(16),
			Name:        name,
			Description: description,
		}
		err := h.repository.AddGame(game)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/games", http.StatusSeeOther)
	default:
		common.HandleError(errors.New("unsupported method"), w, r, http.StatusBadRequest)
	}

}
