package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/uber-demo/customer/api/v1/router"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
}

func InitServer() {
	baseRouter := chi.NewRouter()
	
	RegisterRouters(baseRouter)

	http.ListenAndServe(":3000", baseRouter)

}

func RegisterRouters(baseRouter *chi.Mux) {

	baseRouter.Use(middleware.Logger)

	api := chi.NewRouter()
	v1 := chi.NewRouter()
	v2 := chi.NewRouter()
	customer := chi.NewRouter()
	health := chi.NewRouter()
	test := chi.NewRouter()
	auth := chi.NewRouter()

	baseRouter.Mount("/api", api)
	api.Mount("/v1", v1)
	api.Mount("/v2", v2)
	v1.Mount("/auth", auth)
	v1.Mount("/customer", customer)
	v1.Mount("/health", health)
	v1.Mount("/test", test)

	router.Auth(auth)
	router.Customer(customer)
	router.Health(health)
	router.Test(test)
}