package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/uber-demo/customer/api/v1/helper"
)

func AuthTest(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(w.Header().Get("cust_id"))
		helper.RespondWithJSON(w, http.StatusOK, map[string]string{"Message": w.Header().Get("cust_id") + " is authorized!"})
	})
}