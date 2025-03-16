package main

import (
	"fmt"
	"net/http"
	"strconv"

	"call-center/api"
	"call-center/internal/application/calls/commands"
	"call-center/internal/infrastructure/data/repositories"
	"call-center/internal/shared/constants"

	"github.com/go-chi/chi"
	mediator "github.com/mehdihadeli/go-mediatr"
)

func main() {
	// Repositories
	callRepository := repositories.NewCallRepository()

	// Request handlers
	createCallCommandHandler := commands.NewCreateCallCommandHandler(callRepository)

	// Services
	mediator.RegisterRequestHandler(createCallCommandHandler)

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

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to start GO Call Center API server...")
	}
}
