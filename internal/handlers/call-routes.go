package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func CallRoutesHandler(router chi.Router) {
	// @Router api/call [get]
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

	})

	// @Router api/call [post]
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {

	})
}
