package api

import (
	"call-center/api/routes"

	"github.com/go-chi/chi"
	middle "github.com/go-chi/chi/middleware"
)

func MapRoutes(router *chi.Mux) *chi.Mux {
	router.Use(middle.StripSlashes)

	router.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Route("/call", routes.MapCallRoutes)
			v1.Route("/user", routes.MapUserRoutes)
		})
	})

	return router
}
