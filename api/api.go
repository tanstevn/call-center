package api

import (
	"call-center/internal/handlers"

	"github.com/go-chi/chi"
	middle "github.com/go-chi/chi/middleware"
)

func RoutesHandler(router *chi.Mux) {
	router.Use(middle.StripSlashes)

	router.Route("/api", func(api chi.Router) {
		api.Route("/call", handlers.CallRoutesHandler)
		api.Route("/user", handlers.UserRoutesHandler)
	})
}
