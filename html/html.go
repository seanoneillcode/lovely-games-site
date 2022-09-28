package html

import (
	"embed"
	"io"
	"text/template"
)

//go:embed *
var files embed.FS

var (
	about = parse("about/about.html")
	index = parse("index.html")
)

type AboutParams struct {
	Title string
}

func About(w io.Writer, p AboutParams) error {
	return about.Execute(w, p)
}

type IndexParams struct {
	Title string
}

func Index(w io.Writer, p IndexParams) error {
	return index.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(template.New("layout.html").ParseFS(files, "layout.html", file))
}
