package api

import (
	"call-center/internal/handlers"

	"github.com/go-chi/chi"
	middle "github.com/go-chi/chi/middleware"
)

func RoutesHandler(router *chi.Mux) *chi.Mux {
	router.Use(middle.StripSlashes)

	router.Route("/api", func(api chi.Router) {
		// @Tags calls
		api.Route("/call", handlers.CallRoutesHandler)

		// @Taags users
		api.Route("/user", handlers.UserRoutesHandler)
	})

	return router
}
