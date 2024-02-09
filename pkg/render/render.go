package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"example.com/go-demo-1/pkg/config"
	"example.com/go-demo-1/pkg/models"
)

var app *config.AppConfig
var functions = template.FuncMap{}

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//get the cache from the app config
	tc := app.TemplateCache

	//Requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get the template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing the template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
