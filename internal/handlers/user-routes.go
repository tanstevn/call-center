package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func UserRoutesHandler(router chi.Router) {
	// @Router api/user [get]
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

	})

	// @Router api/user [post]
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {

	})
}
