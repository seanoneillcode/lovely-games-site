package html

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed *
var files embed.FS

type Templates struct {
	templatesCache    map[string]*template.Template
	isDevelopmentMode bool
}

func NewTemplates(isDevelopmentMode bool) *Templates {
	return &Templates{
		templatesCache:    map[string]*template.Template{},
		isDevelopmentMode: isDevelopmentMode,
	}
}

func (r *Templates) GetWrappedTemplate(file string) *template.Template {
	// in development mode, we want to load the file everytime in case it's changed.
	if r.isDevelopmentMode {
		return template.Must(template.ParseFiles("html/wrapper.html", fmt.Sprintf("html/%s", file)))
	}

	// cache the templates to save ms and kb
	tmpl, ok := r.templatesCache[file]
	if !ok {
		tmpl = template.Must(template.New("wrapper.html").ParseFS(files, "wrapper.html", file))
		r.templatesCache[file] = tmpl
	}
	return tmpl
}

func (r *Templates) GetTemplate(file string) *template.Template {
	// in development mode, we want to load the file everytime in case it's changed.
	if r.isDevelopmentMode {
		return template.Must(template.ParseFiles(fmt.Sprintf("html/%s", file)))
	}

	// cache the templates to save ms and kb
	tmpl, ok := r.templatesCache[file]
	if !ok {
		tmpl = template.Must(template.New("tmp").ParseFS(files, file))
		r.templatesCache[file] = tmpl
	}
	return tmpl
}
