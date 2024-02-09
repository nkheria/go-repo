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

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Application running on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
