package api

import (
	routesV1 "call-center/api/routes/v1"

	"github.com/go-chi/chi"
	middle "github.com/go-chi/chi/middleware"
)

func MapRoutes(router *chi.Mux) *chi.Mux {
	router.Use(middle.StripSlashes)

	router.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Route("/call", routesV1.MapCallRoutes)
			v1.Route("/user", routesV1.MapUserRoutes)
		})
	})

	return router
}
