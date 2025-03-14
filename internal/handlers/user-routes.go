package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func UserRoutesHandler(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {

	})
}
