package api

import (
	"fmt"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
}