package main

import (
	"net/http"

	"github.com/GodKimba/super-duper-golang-web-app/pkg/config"
	"github.com/GodKimba/super-duper-golang-web-app/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux

}
