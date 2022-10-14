package handlers

import (
	"net/http"
)

type IndexParams struct {
	Title string
}

func (s *RenderHandlers) Index(w http.ResponseWriter, r *http.Request) {
	p := IndexParams{
		Title: "Index",
	}
	err := s.Templates.GetTemplate("index.html").Execute(w, p)
	if err != nil {
		panic(err)
	}
}
