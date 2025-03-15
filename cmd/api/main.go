package main

import (
	"fmt"
	"net/http"
	"strconv"

	"call-center/api"
	constants "call-center/internal/shared"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	app := api.RoutesHandler(router)

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

	fmt.Println("GO Call Center API server is now running! 🚀")
}
