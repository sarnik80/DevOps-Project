package main

import (
	"net/http"
	"p1-sender/handlers"

	"github.com/go-chi/chi/v5"
)

func GetRoutes(handlers *handlers.Repository) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", handlers.Home)
	mux.Post("/send" , handlers.Send)
	return mux
}
