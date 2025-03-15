package main

import (
	"fmt"
	"net/http"
	"strconv"

	"call-center/api"
	"call-center/internal/shared/constants"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	app := api.MapRoutes(router)

	fmt.Println("Starting GO Call Center API...")

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(constants.Port),
		Handler:      app,
		IdleTimeout:  constants.IdleTimeout,
		ReadTimeout:  constants.ReadTimeout,
		WriteTimeout: constants.WriteTimeout,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to start GO Call Center API server...")
	}
}
