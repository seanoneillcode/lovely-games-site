package handlers

import (
	"net/http"

	"seanoneillcode/lovely-games-site/handlers/common"
)

func (s *RenderHandlers) Play(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Play",
	}
	err := s.Templates.GetTemplate("play.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
	}
}
