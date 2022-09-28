package main

import (
	"fmt"
	"net/http"
)

type fileHandler struct {
	handler http.Handler
}

func NewFileHandler(dir string) *fileHandler {
	return &fileHandler{
		handler: http.FileServer(http.Dir(dir)),
	}
}

func (f *fileHandler) handle(w http.ResponseWriter, r *http.Request) {
	f.handler.ServeHTTP(w, r)
}

func (f *fileHandler) handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello bob!\n")
}
