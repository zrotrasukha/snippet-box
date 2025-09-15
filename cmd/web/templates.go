package main

import (
	"html/template"
	"path/filepath"

	"github.com/zrotrasukha/snippetbox/internal/modeles"
)

type templateData struct {
	Snippet     *modeles.Snippet
	Snippets    []*modeles.Snippet
	CurrentYear int
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// parsing base file
		ts, err := template.ParseFiles("ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// parsing partial file
		ts, err = ts.ParseGlob("ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}
	return cache, nil
}
