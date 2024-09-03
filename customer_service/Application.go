package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	database "github.com/uber-demo/customer/db"
	"github.com/uber-demo/customer/db/migration"
)


type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rating int `json:"rating"`
	Trips []Trip `json:"trips"`
}

type Trip struct {
	Tripid, Source, Destination string 
	Cost float64 
	PaymentMethod string
	Status int
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	fmt.Println(id, len(customersArray))
	if id >= len(customersArray) || id < 0 {
		respondWithError(w, http.StatusBadRequest, "User does not exists")
		return
	}
	respondWithJSON(w, http.StatusOK, customersArray[id])
	// OR
	// render.JSON(w, r, customer) 

	
}

var customersArray = make([]Customer, 0)

func (cust *Customer) Save() string {
	cust.Id = uuid.NewString()
	customersArray = append(customersArray, *cust)
	fmt.Println(customersArray)
	return cust.Id
}


func respondWithError(w http.ResponseWriter, code int, errorMsg string) {
	respondWithJSON(w, code, map[string]string{"message": errorMsg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}


func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, customersArray)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer

	err := render.DecodeJSON(r.Body, &customer)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	var custID = customer.Save()
	respondWithJSON(w, http.StatusOK, map[string]string{"CustID" : custID})
}


func main() {
	router := chi.NewRouter()

	err := database.Init() 
	defer database.CloseDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	migration.Upgrade1()

	router.Use(middleware.Logger)

	router.Get("/v1/customer/{id}", GetCustomer)
	router.Get("/v1/customer", GetAllCustomers)
	router.Post("/v1/customer", CreateCustomer)

	http.ListenAndServe(":3000", router)

}