package main

import (
	"net/http"

	"example.com/go-demo-1/pkg/config"
	"example.com/go-demo-1/pkg/handlers"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {

	// mux := pat.New()

	// mux.Get("/", http.HandleFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandleFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux

}
