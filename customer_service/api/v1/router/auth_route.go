package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/uber-demo/customer/api/v1/controllers"
)

func Auth(r *chi.Mux) {
	r.Post("/login", controllers.LoginCustomer)
}