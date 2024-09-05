package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/uber-demo/customer/api/v1/helper"
)

func Test(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helper.RespondWithJSON(w, 200, map[string]string{"msg":"test"})
	})
}