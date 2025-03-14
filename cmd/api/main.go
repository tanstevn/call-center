package main

import (
	"call-center/api"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	app := chi.NewRouter()

	api.RoutesHandler(app)

	err := http.ListenAndServe("localhost:8000", app)

	if err != nil {
		fmt.Println(err.Error())
	}
}
