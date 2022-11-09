package games

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

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

	allGames, err := h.repository.GetGames()
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Title": "About",
		"Games": allGames,
	}
	err = h.templates.GetWrappedTemplate("games/games.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
	}
}

func (h *GameHandler) GameFrame(w http.ResponseWriter, r *http.Request) {
	id := common.GetQueryParam("id", r.URL)
	if id == "" {
		common.HandleError(errors.New("id invalid"), w, r, http.StatusBadRequest)
		return
	}

	game, err := h.repository.GetGame(id)
	if err != nil {
		common.HandleError(err, w, r, http.StatusBadRequest)
		return
	}
	if game == nil {
		common.HandleError(err, w, r, http.StatusBadRequest)
		return
	}
	data := map[string]interface{}{
		"GameFile": game.GameFile,
	}
	err = h.templates.GetTemplate("games/frame.html").Execute(w, data)
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
		err := h.templates.GetWrappedTemplate("games/upload.html").Execute(w, data)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
		}
	case "POST":
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			common.HandleError(err, w, r, http.StatusBadRequest)
			return
		}
		name := r.FormValue("game-name")
		description := r.FormValue("game-description")
		frameWidth, err := strconv.Atoi(r.FormValue("game-frame-width"))
		if err != nil {
			frameWidth = 480
		}
		frameHeight, err := strconv.Atoi(r.FormValue("game-frame-height"))
		if err != nil {
			frameHeight = 360
		}
		formScreenshotFile, _, err := r.FormFile("game-screenshot")
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
			return
		}
		defer formScreenshotFile.Close()
		screenshotFilename := common.RandomCharacters(16) + ".png"

		// This is path which we want to store the file
		f, err := os.OpenFile("static/screenshots/"+screenshotFilename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
			return
		}
		// Copy the file to the destination path
		_, err = io.Copy(f, formScreenshotFile)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
			return
		}

		formGameFile, _, err := r.FormFile("game-file")
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
			return
		}
		defer formGameFile.Close()
		formGameFileFilename := common.RandomCharacters(16) + ".wasm"

		// This is path which we want to store the file
		f2, err := os.OpenFile("static/games/"+formGameFileFilename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
			return
		}
		// Copy the file to the destination path
		_, err = io.Copy(f2, formGameFile)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
			return
		}

		game := Game{
			Id:          common.RandomCharacters(16),
			Name:        name,
			Description: description,
			Screenshot:  screenshotFilename,
			GameFile:    formGameFileFilename,
			FrameWidth:  frameWidth,
			FrameHeight: frameHeight,
		}
		err = h.repository.AddGame(&game)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/games", http.StatusSeeOther)
	default:
		common.HandleError(errors.New("unsupported method"), w, r, http.StatusBadRequest)
	}
}

func (h *GameHandler) EditGame(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		data := map[string]interface{}{
			"Title": "Edit Game",
		}
		err := h.templates.GetWrappedTemplate("games/edit.html").Execute(w, data)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
		}
	case "POST":
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			common.HandleError(err, w, r, http.StatusBadRequest)
			return
		}
		id := common.GetQueryParam("id", r.URL)
		name := r.FormValue("game-name")
		description := r.FormValue("game-description")

		game := Game{
			Name:        name,
			Description: description,
		}
		err = h.repository.EditGame(id, &game)
		if err != nil {
			common.HandleError(err, w, r, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/games", http.StatusSeeOther)
	default:
		common.HandleError(errors.New("unsupported method"), w, r, http.StatusBadRequest)
	}
}

func (h *GameHandler) Play(w http.ResponseWriter, r *http.Request) {
	id := common.GetQueryParam("id", r.URL)

	game, err := h.repository.GetGame(id)
	if err != nil {
		common.HandleError(err, w, r, http.StatusBadRequest)
		return
	}
	if game == nil {
		common.HandleError(err, w, r, http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"Title":       "Play",
		"Id":          id,
		"FrameWidth":  game.FrameWidth,
		"FrameHeight": game.FrameHeight,
	}
	err = h.templates.GetWrappedTemplate("games/play.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
	}
}
