package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/go-demo-1/pkg/config"
	"example.com/go-demo-1/pkg/handlers"
	"example.com/go-demo-1/pkg/render"
)

const portNumber = ":8081"

func main() {

	var app config.AppConfig

	tc, err := render.createTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Application running on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
