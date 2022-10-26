package handlers

import (
	"net/http"

	"seanoneillcode/lovely-games-site/handlers/common"
)

func (s *RenderHandlers) About(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "About",
	}
	err := s.Templates.GetTemplate("about/about.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r)
	}
}
