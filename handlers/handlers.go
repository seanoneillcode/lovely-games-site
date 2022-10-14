package handlers

import "seanoneillcode/lovely-games-site/html"

type RenderHandlers struct {
	Templates *html.Templates
}

func NewRenderHandlers(isDevelopmentMode bool) *RenderHandlers {
	return &RenderHandlers{
		Templates: html.NewTemplates(isDevelopmentMode),
	}
}
