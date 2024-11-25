package main

import (
	"fmt"
	"net/http"

	"github.com/aghyad-khlefawi/identity/api"
	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	fmt.Printf("HTTP Server listening on port 8080")

	api.RegisterRoutes(router)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err.Error())
	}

}
