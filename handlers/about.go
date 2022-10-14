package handlers

import (
	"net/http"
)

type AboutParams struct {
	Title string
}

func (s *RenderHandlers) About(w http.ResponseWriter, r *http.Request) {
	p := AboutParams{
		Title: "About",
	}
	err := s.Templates.GetTemplate("about/about.html").Execute(w, p)
	if err != nil {
		panic(err)
	}
}
