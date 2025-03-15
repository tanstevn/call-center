package routes

import (
	"call-center/api/controllers"

	"github.com/go-chi/chi"
)

func MapUserRoutes(router chi.Router) {
	controller := controllers.NewUserController()

	router.Get("/list", controller.GetUserList())
}
