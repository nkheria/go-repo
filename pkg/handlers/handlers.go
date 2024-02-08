package handlers

import (
	"net/http"

	"example.com/go-demo-1/pkg/render"
)

const portNumber = ":8081"

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
