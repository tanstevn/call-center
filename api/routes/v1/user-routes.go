package v1

import (
	controllersV1 "call-center/api/controllers/v1"

	"github.com/go-chi/chi"
)

func MapUserRoutes(router chi.Router) {
	controller := controllersV1.NewUserController()

	router.Get("/list", controller.GetUserList())
}
