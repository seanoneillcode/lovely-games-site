package handlers

import (
	"net/http"

	"seanoneillcode/lovely-games-site/handlers/common"
	"seanoneillcode/lovely-games-site/html"
)

type RenderHandlers struct {
	Templates *html.Templates
}

func NewRenderHandlers(isDevelopmentMode bool) *RenderHandlers {
	return &RenderHandlers{
		Templates: html.NewTemplates(isDevelopmentMode),
	}
}

func (s *RenderHandlers) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "About",
	}
	err := s.Templates.GetWrappedTemplate("index.html").Execute(w, data)
	if err != nil {
		common.HandleError(err, w, r, http.StatusInternalServerError)
	}
}
