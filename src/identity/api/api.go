package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func RegisterRoutes(mux *chi.Mux) {

  mux.Post("/user",HandleCreateUser);

  mux.Get("/hc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Healthy")
	})


}
