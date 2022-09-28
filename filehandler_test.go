package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {

	f := NewIndexHandler("./static")
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	f.handleHello(rr, req)

	got := rr.Body.String()
	want := "Hello bob!\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFileHandler(t *testing.T) {

	f := NewIndexHandler("./static")
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	f.handleIndex(rr, req)

	got := rr.Body.String()
	want := "<title>A static page</title>"

	if !strings.Contains(got, want) {
		t.Errorf("expected GET to return index file containing %q\nactually got:\n%s", want, got)
	}
}
