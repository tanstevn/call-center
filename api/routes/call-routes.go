package routes

import (
	"call-center/api/controllers"

	"github.com/go-chi/chi"
)

func MapCallRoutes(router chi.Router) {
	controller := controllers.NewCallController()

	router.Post("", controller.CreateCall())
}
