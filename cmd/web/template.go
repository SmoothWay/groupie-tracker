package main

import (
	"html/template"
	"path/filepath"

	groupie "github.com/SmoothWay/groupie-tracker/pkg"
)

type templateData struct {
	CurrentYear int
	Page        *groupie.ArtOutput
	Data        interface{}
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*-page.html"))
	if err != nil {
		return nil, err
	}

	// loop through the pages one by one.

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*-layout.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*-partial.html"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts

	}
	return cache, nil
}
