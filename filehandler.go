package main

import (
	"fmt"
	"net/http"
)

type indexHandler struct {
	handler http.Handler
}

func NewIndexHandler(dir string) *indexHandler {
	return &indexHandler{
		handler: http.FileServer(http.Dir(dir)),
	}
}

func (f *indexHandler) handleIndex(w http.ResponseWriter, r *http.Request) {
	f.handler.ServeHTTP(w, r)
}

func (f *indexHandler) handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello bob!\n")
}
