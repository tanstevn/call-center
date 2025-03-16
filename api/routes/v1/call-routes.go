package v1

import (
	controllersV1 "call-center/api/controllers/v1"

	"github.com/go-chi/chi"
)

func MapCallRoutes(router chi.Router) {
	controller := controllersV1.NewCallController()

	router.Post("/", controller.CreateCall())
}
