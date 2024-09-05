package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/uber-demo/customer/api/v1/controllers"
	"github.com/uber-demo/customer/api/v1/helper"
)

func Customer(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helper.RespondWithJSON(w, 200, map[string]string{"Msg":"All customers"})
	})
	r.Get("/{id}", controllers.GetCustomer)
	r.Post("/", controllers.CreateCustomer)
}
