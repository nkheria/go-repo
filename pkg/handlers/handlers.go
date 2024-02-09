package handlers

import (
	"net/http"

	"example.com/go-demo-1/pkg/config"
	"example.com/go-demo-1/pkg/models"
	"example.com/go-demo-1/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

const portNumber = ":8081"

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Welcome again!"
	//send the data to the template

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
