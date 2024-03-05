package main

import (
	"KazakhAliexpress/SE2220/pkg/forms"
	"KazakhAliexpress/SE2220/pkg/models"
	"html/template"
	"net/url"
	"path/filepath"
)

type templateData struct {
	Items       []*models.Item
	Item        *models.Item
	CurrentYear int
	FormData    url.Values
	FormErrors  map[string]string
	Form        *forms.Form
	Flash       string
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		// Use the ParseGlob method to add any 'partial' templates to the
		// template set (in our case, it's just the 'footer' partial at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		// Add the template set to the cache, using the name of the page
		// (like 'home.page.tmpl') as the key.
		cache[name] = ts
	}
	return cache, nil
}
