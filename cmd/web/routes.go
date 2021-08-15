package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swagathushetty/bookings/pkg/config"
	"github.com/swagathushetty/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//set up routes here

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(Nosurf)
	mux.Use(SessionLoad)
	// mux.Use(WriteToConsole)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//file server to fetch static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

//using pat
// func routes(app *config.AppConfig) http.Handler {
// 	mux := pat.New()

// 	//set up routes here

// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
// 	return mux
// }
