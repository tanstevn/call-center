package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"call-center/api"
	"call-center/api/dependencies"
	"call-center/internal/shared/constants"

	"github.com/go-chi/chi"
)

func main() {
	if err := dependencies.RegisterMediatorHandlers(); err != nil {
		log.Fatalf("Mediator registration failed: %v", err)
	}

	router := chi.NewRouter()
	app := api.MapRoutes(router)

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(constants.Port),
		Handler:      app,
		IdleTimeout:  constants.IdleTimeout,
		ReadTimeout:  constants.ReadTimeout,
		WriteTimeout: constants.WriteTimeout,
	}

	fmt.Println("Starting GO Call Center API...🚀")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start GO Call Center API server...")
	}
}
