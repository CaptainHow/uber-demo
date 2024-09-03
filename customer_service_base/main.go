package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/uber-demo/customer/api/v1/middleware"
)




func GetHello(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, map[string]string{"hello":"world"})
}



func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	res, err := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(code)
		w.Write(res)
	}
}

func respondWithError(w http.ResponseWriter, code int, errorMsg string) {
	RespondWithJSON(w, code, map[string]string{"message": errorMsg})
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/hello", GetHello)

	server := http.Server{
		Addr: ":3000",
		Handler: middleware.Logging(router),
	}

	log.Println("Starting server on port :3000")
	err := server.ListenAndServe()
	fmt.Println(err)
	
}