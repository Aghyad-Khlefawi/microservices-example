package api

import (
	"fmt"
	"net/http"

	"github.com/aghyad-khlefawi/identity/pkg/servicecollection"
	"github.com/go-chi/chi/v5"
)

type ApiContext struct{
  sc *servicecollection.ServiceCollection
}


func RegisterRoutes(mux *chi.Mux,sc * servicecollection.ServiceCollection) {

  context:= &ApiContext{
		sc,
	}

  mux.Post("/user",NewHandler(HandleCreateUser, context));

  mux.Get("/hc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Healthy")
	})


}


func NewHandler(handler (func(w http.ResponseWriter, r *http.Request, context *ApiContext)), context *ApiContext) (func(http.ResponseWriter, *http.Request)){

	return func(w http.ResponseWriter, r *http.Request) {
		handler(w,r,context)
	}
}
